package parser

import (
	"aqwari.net/xml/xsd"
	"aqwari.net/xml/xsdgen"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/cmds/generate/registry"
	"github.com/rs/zerolog/log"
	"math"
	"strconv"
	"strings"
)

type Config struct {
	Registry registry.Config `yaml:"registry" mapstructure:"registry" json:"registry"`
}

type Parser struct {
	Cfg          Config
	TypeRegistry registry.TypeRegistry
}

func NewParser(cfg Config) Parser {
	p := Parser{Cfg: cfg, TypeRegistry: registry.NewTypeRegistry(&cfg.Registry)}
	return p
}

func (p *Parser) Parse(msg registry.ISO20022Message, xsdData []byte) error {

	p.TypeRegistry.Msgs = append(p.TypeRegistry.Msgs, msg)

	var cfg xsdgen.Config
	cfg.Option(
		xsdgen.PackageName(msg.PackageName()),
	)

	msgCode, err := cfg.GenCode(xsdData)
	if err != nil {
		return err
	}

	doc, ok := msgCode.DocType(string(msg.Namespace()))
	if !ok {
		return fmt.Errorf("not a root document")
	}

	p.visit(doc)
	return nil
}

func (p *Parser) visit(cplx *xsd.ComplexType) {

	for _, attr := range cplx.Attributes {

		switch tattr := attr.Type.(type) {
		case *xsd.ComplexType:
		case *xsd.SimpleType:
			te := p.newSimpleTypeRegistryEntry(tattr)
			p.addEntry(te)
		case xsd.Builtin:
			te := p.newBuiltinRegistryEntry(cplx.Name.Local, tattr)
			p.addEntry(te)
		default:
			log.Error().Interface("t", attr.Type).Msg("attr type not matched")
		}
	}

	for _, el := range cplx.Elements {

		switch tel := el.Type.(type) {
		case *xsd.ComplexType:
			te := p.newComplexTypeRegistryEntry(tel)
			te.MinOccurs, te.MaxOccurs = MinMaxElementOccurs(el)
			te.ChoiceSmell = XsChoiceSmell(tel)
			p.addEntry(te)
			p.visit(tel)
		case *xsd.SimpleType:
			te := p.newSimpleTypeRegistryEntry(tel)
			te.MinOccurs, te.MaxOccurs = MinMaxElementOccurs(el)
			p.addEntry(te)
		case xsd.Builtin:
			te := p.newBuiltinRegistryEntry(el.Name.Space, tel)
			te.MinOccurs, te.MaxOccurs = MinMaxElementOccurs(el)
			p.addEntry(te)
		default:
			log.Error().Interface("t", el.Type).Msg("element type not matched")
		}
	}
}

func (p *Parser) newComplexTypeRegistryEntry(t *xsd.ComplexType) registry.TypeEntry {

	ns := registry.ISO20022Namespace(t.Name.Space)
	// log.Trace().Str("space", t.Name.Space).Str("local", t.Name.Local).Msg("found complex type")

	te := registry.TypeEntry{
		ElementType: registry.ElementTypeComplex,
		Spaces:      []registry.ISO20022Namespace{ns},
		Local:       t.Name.Local,
		Complex:     t,
	}

	return te
}

func (p *Parser) newSimpleTypeRegistryEntry(t *xsd.SimpleType) registry.TypeEntry {

	ns := registry.ISO20022Namespace(t.Name.Space)
	// log.Trace().Str("space", t.Name.Space).Str("local", t.Name.Local).Msg("found simple type")

	/*
		skipSimpleArray := []string{"Max10MbBinary", "xsdDate", "xsdDateTime"} // "Max10MbBinary", "xsdDate", "xsdDateTime", "ISODateTime", "ISODate"
		for _, s := range skipSimpleArray {
			if s == t.Name.Local {
				return
			}
		}
	*/

	te := registry.TypeEntry{
		ElementType:  registry.ElementTypeSimple,
		Spaces:       []registry.ISO20022Namespace{ns},
		Local:        t.Name.Local,
		Simple:       t,
		Restrictions: t.Restriction,
	}

	return te
}

func (p *Parser) newBuiltinRegistryEntry(s string, t xsd.Builtin) registry.TypeEntry {

	// instead of taking the type namespace (i.e. http://www.w3.org/2001/XMLSchema) i take the element namespace that is I now in what space it has been referred
	ns := registry.ISO20022Namespace(s)
	// log.Trace().Str("space", t.Name().Space).Str("local", t.Name().Local).Msg("found builtin type")

	te := registry.TypeEntry{
		ElementType: registry.ElementTypeBuiltin,
		Spaces:      []registry.ISO20022Namespace{ns},
		Local:       t.Name().Local,
		Builtin:     t,
	}

	return te
}

func (p *Parser) addEntry(te registry.TypeEntry) {

	p.TypeRegistry.AddEntry(te)

	/*
		n := te.Local
		if n == "Document" {
			n = strings.Join([]string{n, te.Spaces[0].Message()}, "#")
		}

		// TODO
		if te.Local == "Max35Text--" {
			fmt.Println()
		}

		found, ok := p.TypeRegistry.Types[te.Local]
		if ok {
			if !found.ContainsSpace(te.Spaces[0]) {
				found.Spaces = append(found.Spaces, te.Spaces[0])
			}
			found.RefCount = 1 + found.RefCount
			p.TypeRegistry.Types[found.Local] = found
		} else {
			te.RefCount = 1
			p.TypeRegistry.Types[te.Local] = te
		}

	*/

}

func MinMaxElementOccurs(el xsd.Element) (int, int) {
	min := 1
	max := 1
	for _, a := range el.Attr {
		if a.Name.Local == "minOccurs" {
			if minOccurs, err := strconv.Atoi(a.Value); err == nil {
				min = minOccurs
			} else {
				log.Error().Err(err).Str("name", el.Name.Local).Str("value", a.Value).Msg("invalid minOccurs attribute")
			}
		}

		if a.Name.Local == "maxOccurs" {
			if a.Value == "unbounded" {
				max = math.MaxInt
			} else {
				if v, err := strconv.Atoi(a.Value); err == nil {
					max = v
				} else {
					log.Error().Err(err).Str("name", el.Name.Local).Str("value", a.Value).Msg("invalid maxOccurs attribute")
				}
			}
		}
	}

	return min, max
}

// XsChoiceSmell Difficult to tell if a complex is a choice or not. A hint could be the name that in iso-20022 is pretty consistent. Anothe roption could be
// the fact that all the children are optional... As the matter of fact, this doen't hold true since Contact4 is a sequence with all optional fields.... but no choice...
// Apparently all choices in pain-13 and 14 have two elements...
func XsChoiceSmell(tel *xsd.ComplexType) bool {

	if !strings.HasSuffix(tel.Name.Local, "Choice") {
		return false
	}

	for _, child := range tel.Elements {
		if !child.Optional {
			return false
		}
	}

	// log.Trace().Str("name", tel.Name.Local).Int("num-children", len(tel.Elements)).Msg("is choice...")
	return true
}
