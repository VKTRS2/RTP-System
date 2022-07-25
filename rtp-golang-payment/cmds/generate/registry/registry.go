package registry

import (
	"aqwari.net/xml/xsd"
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

type ISO20022Namespace string

func ISO20022NamespaceForMessage(n string) ISO20022Namespace {
	return ISO20022Namespace(fmt.Sprintf("urn:iso:std:iso:20022:tech:xsd:%s", n))
}

func (n ISO20022Namespace) Message() string {
	return strings.TrimPrefix(string(n), "urn:iso:std:iso:20022:tech:xsd:")
}

func packageNameForMessageName(n string) string {
	return strings.ReplaceAll(n, ".", "_")
}

type ISO20022Message struct {
	Name string
}

func (i ISO20022Message) PackageName() string {
	return packageNameForMessageName(i.Name)
}

func (i ISO20022Message) Namespace() ISO20022Namespace {
	return ISO20022NamespaceForMessage(i.Name)
}

type ElementType string

const (
	ElementTypeComplex ElementType = "complex"
	ElementTypeSimple              = "simple"
	ElementTypeBuiltin             = "builtin"
)

type TypeEntry struct {
	RefCount     int
	ElementType  ElementType
	Spaces       []ISO20022Namespace
	Local        string
	Complex      *xsd.ComplexType
	Simple       *xsd.SimpleType
	Builtin      xsd.Builtin
	Restrictions xsd.Restriction
	MinOccurs    int
	MaxOccurs    int
	ChoiceSmell  bool
}

func (te *TypeEntry) ContainsSpace(ns ISO20022Namespace) bool {
	for _, s := range te.Spaces {
		if s == ns {
			return true
		}
	}

	return false
}

func (te *TypeEntry) ComputePackageName(simpleTypesInCommon bool) string {
	pkg := strings.ReplaceAll(te.Spaces[0].Message(), ".", "_")
	if len(te.Spaces) > 1 || te.ElementType == ElementTypeSimple {
		pkg = "common"
	}

	return pkg
}

type Config struct {
	SimpleTypesInCommonPackage bool `yaml:"simple-types-in-common-package" mapstructure:"simple-types-in-common-package" json:"simple-types-in-common-package"`
}

type TypeRegistry struct {
	Cfg   *Config
	Msgs  []ISO20022Message
	Types map[string]TypeEntry
}

func NewTypeRegistry(cfg *Config) TypeRegistry {
	return TypeRegistry{Cfg: cfg, Types: make(map[string]TypeEntry)}
}

func (tr *TypeRegistry) AddEntry(te TypeEntry) {

	if te.Local == "Document" {
		te.Local = strings.Join([]string{te.Local, te.Spaces[0].Message()}, "#")
	}

	// TODO
	if te.Local == "Max35Text--" {
		fmt.Println()
	}

	found, ok := tr.Types[te.Local]
	if ok {
		if !found.ContainsSpace(te.Spaces[0]) {
			found.Spaces = append(found.Spaces, te.Spaces[0])
		}
		found.RefCount = 1 + found.RefCount
		tr.Types[found.Local] = found
	} else {
		te.RefCount = 1
		tr.Types[te.Local] = te
	}

}

func (tr *TypeRegistry) GetEntriesForPackage(elType ElementType, pkgName string) []TypeEntry {

	var tes []TypeEntry
	for _, te := range tr.Types {

		// TODO
		if te.Local == "Max35Text--" {
			fmt.Println()
		}
		if te.ElementType == elType {
			if pkgName == te.ComputePackageName(tr.Cfg.SimpleTypesInCommonPackage) {
				tes = append(tes, te)
			}
		}
	}

	return tes
}

func (tr *TypeRegistry) GetEntryForType(tn string) (TypeEntry, bool) {
	te, ok := tr.Types[tn]
	return te, ok
}

func (tr *TypeRegistry) ShowInfo(inclusionList []ISO20022Message) {

	var pkgs []string
	var tmpList []ISO20022Message
	if inclusionList != nil {
		tmpList = inclusionList
	} else {
		tmpList = tr.Msgs
	}
	for _, m := range tmpList {
		pkgs = append(pkgs, m.PackageName())
	}

	pkgs = append(pkgs, "common")

	for _, pkg := range pkgs {
		log.Trace().Str("pkg", pkg).Msg("complex elements in package")
		showRegistryEntries(tr, ElementTypeComplex, pkg)

		log.Trace().Str("pkg", pkg).Msg("simple elements in package")
		showRegistryEntries(tr, ElementTypeSimple, pkg)

		log.Trace().Str("pkg", pkg).Msg("builtin elements in package")
		showRegistryEntries(tr, ElementTypeBuiltin, pkg)
	}
}

func showRegistryEntries(tr *TypeRegistry, elType ElementType, pkgName string) {
	entries := tr.GetEntriesForPackage(elType, pkgName)
	for i, e := range entries {
		log.Trace().Int("i", i).Str("local", e.Local).Int("ref-count", e.RefCount).Int("num-spaces", len(e.Spaces)).Send()
	}
}
