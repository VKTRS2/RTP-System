package model

import (
	"aqwari.net/xml/xsd"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/cmds/generate/registry"
	"github.com/rs/zerolog/log"
	"strings"
)

type BuiltinTypeMapping struct {
	Builtin        xsd.Builtin
	GolangType     string
	ImportRequired string
}

// types from github.com/metaleap/go-xsd/types
var builtinTypeRegistry = map[string]BuiltinTypeMapping{
	"Decimal":      {Builtin: xsd.Decimal, GolangType: "xsdt.Decimal" /*, ImportRequired: "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt" */},
	"String":       {Builtin: xsd.String, GolangType: "xsdt.String" /*, ImportRequired: "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt" */},
	"Date":         {Builtin: xsd.Date, GolangType: "xsdt.Date" /*, ImportRequired: "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt" */},
	"DateTime":     {Builtin: xsd.DateTime, GolangType: "xsdt.DateTime" /*, ImportRequired: "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt" */},
	"Boolean":      {Builtin: xsd.Boolean, GolangType: "xsdt.Boolean" /*, ImportRequired: "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt" */},
	"Base64Binary": {Builtin: xsd.Base64Binary, GolangType: "xsdt.Base64Binary" /*, ImportRequired: "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt" */},
	"AnyType":      {Builtin: xsd.AnyType, GolangType: "xsdt.AnyType" /*, ImportRequired: "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt" */},
}

func MapBuiltinToGoType(t string, xsdtDefaultImport string) (GoType, error) {
	if m, ok := builtinTypeRegistry[t]; ok {

		// If the import has not been specified but is sort of packaged type (i.e. has a dot in the name.... then use the default from config
		imp := m.ImportRequired
		if imp == "" && strings.Index(m.GolangType, ".") > 0 {
			imp = xsdtDefaultImport
		}
		return GoType{Name: m.GolangType, Import: imp, IsBuiltin: true, Builtin: m.Builtin}, nil
	}

	return GoType{}, fmt.Errorf("could not resolve %s builtin type", t)
}

type GoModel struct {
	Cfg          *ModelConfig
	Msgs         []registry.ISO20022Message
	TypeDefs     []GoTypeDefinition
	TypeRegistry registry.TypeRegistry
}

type GoModelVisitor interface {
	Visit(elName, elType string, isStruct, isPrt, isArray bool) (string, error)
	Pop() (string, error)
}

type GoTypeDefinition struct {
	PackageName  string
	Name         string
	Doc          string
	StructDef    bool
	Type         GoType
	Attrs        []GoAttr
	Restrictions xsd.Restriction
}

func (gs GoTypeDefinition) String() string {
	return gs.Name
}

func (gs GoTypeDefinition) Comment() string {
	if gs.Doc == "" {
		return "type definition"
	}

	return gs.Doc
}

func (gs GoTypeDefinition) HasEnumRestriction() bool {
	if len(gs.Restrictions.Enum) > 0 {
		return true
	}

	return false
}

func (gs GoTypeDefinition) HasLengthRestriction() bool {
	if gs.Restrictions.Length != 0 || gs.Restrictions.MinLength != 0 || gs.Restrictions.MaxLength != 0 {
		return true
	}

	return false
}

func (gs GoTypeDefinition) HasPatternRestriction() bool {
	if gs.Restrictions.Pattern != nil {
		return true
	}

	return false
}

func (gs GoTypeDefinition) HasDecimalRestriction() bool {
	if gs.Restrictions.TotalDigits != 0 || gs.Restrictions.Precision != 0 {
		return true
	}

	return false
}

func (gs GoTypeDefinition) HasAnyRestriction() bool {
	if gs.HasDecimalRestriction() || gs.HasLengthRestriction() || gs.HasEnumRestriction() || gs.HasPatternRestriction() {
		return true
	}

	return false
}

func (gs GoTypeDefinition) GenerateSampleValue() string {

	s := "unrestricted-string"
	if gs.HasDecimalRestriction() {
		log.Warn().Interface("restriction", gs.Restrictions).Msg("unsupported decimal restriction")
	}

	if gs.HasLengthRestriction() {
		l := gs.Restrictions.Length
		if l == 0 {
			l = gs.Restrictions.MinLength + (gs.Restrictions.MaxLength-gs.Restrictions.MinLength)/2
		}

		s = util.GenerateStringOfLength(l)
	}

	if gs.HasEnumRestriction() {
		s = gs.Restrictions.Enum[len(gs.Restrictions.Enum)/2]
	}

	if gs.HasPatternRestriction() {
		var err error
		s, err = util.GenerateStringOfPattern(gs.Restrictions.Pattern.String(), -1)
		if err != nil {
			log.Warn().Err(err).Str("pattern", gs.Restrictions.Pattern.String()).Msg("error of pattern")
		}
	}

	return s
}

type GoType struct {
	Name      string
	Import    string
	IsBuiltin bool
	Builtin   xsd.Builtin
}

type GoAttr struct {
	Name       string
	Type       GoType
	XMLAttr    bool
	Chardata   bool
	Array      bool
	StructType bool
	Optional   bool
}

func (ga GoAttr) IsPtr() bool {
	if ga.StructType {
		if ga.Optional && !ga.Array {
			return true
		}
	}

	return false
}

func (ga GoAttr) XMLName() string {
	if ga.Chardata || (ga.Type.IsBuiltin && ga.Type.Builtin == xsd.AnyType) {
		return ""
	}
	return ga.Name
}

func (ga GoAttr) XMLTags() string {

	var sb strings.Builder

	if ga.XMLAttr {
		sb.WriteString(",attr")
	}

	if ga.Optional {
		sb.WriteString(",omitempty")
	}

	if ga.Chardata {
		sb.WriteString(",chardata")
	}

	if ga.Type.IsBuiltin && ga.Type.Builtin == xsd.AnyType {
		sb.WriteString(",any")
	}

	return sb.String()
}

func (ga GoAttr) String() string {

	modifier := ""
	omitEmpty := ",omitempty"
	if !ga.Optional {
		omitEmpty = ""
	}

	if ga.Array {
		modifier = "[]"
	}

	if ga.StructType {
		if ga.Optional && !ga.Array {
			modifier = "*"
		}
	}

	return fmt.Sprintf("%s\t%s%s `xml:\"%s%s\"`", ga.Name, modifier, ga.Type.Name, ga.Name, omitEmpty)
}
