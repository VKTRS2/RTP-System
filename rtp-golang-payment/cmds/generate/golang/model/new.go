package model

import (
	"aqwari.net/xml/xsd"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/cmds/generate/registry"
	"strings"
)

func NewModel(cfg *ModelConfig, msgs []registry.ISO20022Message, tr registry.TypeRegistry) (GoModel, error) {
	gm := GoModel{Cfg: cfg, Msgs: msgs, TypeRegistry: tr}

	for _, msg := range msgs {
		goMsgTypeDefs, err := gm.newTypeDefinitions(msg.PackageName(), tr)
		if err != nil {
			return gm, err
		}

		gm.TypeDefs = append(gm.TypeDefs, goMsgTypeDefs...)
	}

	goCommonTypedefs, err := gm.newTypeDefinitions("common", tr)
	if err != nil {
		return gm, err
	}
	gm.TypeDefs = append(gm.TypeDefs, goCommonTypedefs...)
	return gm, nil
}

func (gm *GoModel) newTypeDefinitions(pkgName string, tr registry.TypeRegistry) ([]GoTypeDefinition, error) {

	goTypeDefs, err := gm.newComplexTypeDefinitions(pkgName, tr)
	if err != nil {
		return nil, err
	}

	goSimpleTypeDefs, err := gm.newSimpleTypeDefinitions(pkgName, tr)
	if err != nil {
		return nil, err
	}

	return append(goTypeDefs, goSimpleTypeDefs...), nil
}

func (gm *GoModel) newComplexTypeDefinitions(pkgName string, tr registry.TypeRegistry) ([]GoTypeDefinition, error) {

	var goStructs []GoTypeDefinition
	entries := tr.GetEntriesForPackage(registry.ElementTypeComplex, pkgName)
	for _, e := range entries {
		tpn := e.Local
		if strings.HasPrefix(e.Local, "Document#") {
			tpn = "Document"
		}

		gs := GoTypeDefinition{PackageName: pkgName, Name: tpn, StructDef: true, Type: GoType{Name: tpn}, Doc: e.Complex.Doc}
		for _, attr := range e.Complex.Attributes {
			var err error
			var gt GoType
			ga := GoAttr{Name: attr.Name.Local, XMLAttr: true, Optional: attr.Optional}
			switch tattr := attr.Type.(type) {
			case *xsd.SimpleType:
				gt, err = gm.goTypeOfElement(tr, pkgName, tattr.Name.Local)
				if err != nil {
					return nil, err
				}

				ga.Type = gt
			case xsd.Builtin:
				ga.Type, err = MapBuiltinToGoType(tattr.String(), gm.Cfg.XsDtPackageImport)
				if err != nil {
					return nil, fmt.Errorf("%s - %s", attr.Name.Local, err.Error())
				}
			}

			gs.Attrs = append(gs.Attrs, ga)
		}

		// Handle cases such as: ActiveCurrencyAndAmount
		if len(e.Complex.Elements) == 0 {
			if tbase, ok := e.Complex.Base.(xsd.Builtin); ok {
				ga := GoAttr{Name: "Value", Chardata: true}
				var err error
				ga.Type, err = MapBuiltinToGoType(tbase.String(), gm.Cfg.XsDtPackageImport)
				if err != nil {
					return nil, fmt.Errorf("%s - %s", e.Local, err.Error())
				}
				gs.Attrs = append(gs.Attrs, ga)
			}
		}

		for _, el := range e.Complex.Elements {

			var err error
			var gt GoType
			ga := GoAttr{Name: el.Name.Local, Array: el.Plural, Optional: el.Optional}

			switch tel := el.Type.(type) {
			case *xsd.ComplexType:
				gt, err = gm.goTypeOfElement(tr, pkgName, tel.Name.Local)
				if err != nil {
					return nil, err
				}

				ga.StructType = true
				ga.Type = gt
			case *xsd.SimpleType:
				gt, err = gm.goTypeOfElement(tr, pkgName, tel.Name.Local)
				if err != nil {
					return nil, err
				}

				ga.Type = gt
			case xsd.Builtin:
				ga.Type, err = MapBuiltinToGoType(tel.String(), gm.Cfg.XsDtPackageImport)
				if err != nil {
					return nil, fmt.Errorf("(%s) - %s", el.Name.Local, err.Error())
				}

				if ga.Name == "" {
					ga.Name = "Item"
				}
			}

			gs.Attrs = append(gs.Attrs, ga)
		}

		goStructs = append(goStructs, gs)
	}
	return goStructs, nil
}

func (gm *GoModel) newSimpleTypeDefinitions(pkgName string, tr registry.TypeRegistry) ([]GoTypeDefinition, error) {

	var goStructs []GoTypeDefinition
	entries := tr.GetEntriesForPackage(registry.ElementTypeSimple, pkgName)
	for _, e := range entries {
		if bt, ok := e.Simple.Base.(xsd.Builtin); ok {
			gt, err := MapBuiltinToGoType(bt.String(), gm.Cfg.XsDtPackageImport)
			if err != nil {
				return nil, err
			}
			r := e.Restrictions
			gs := GoTypeDefinition{PackageName: pkgName, Name: e.Local, StructDef: false, Type: gt, Doc: e.Simple.Doc, Restrictions: r}
			goStructs = append(goStructs, gs)
		} else {
			return nil, fmt.Errorf("the Base type is not builtin for %s", e.Local)
		}
	}

	return goStructs, nil
}

func (gm *GoModel) goTypeOfElement(tr registry.TypeRegistry, currentPkg string, localTypeName string) (GoType, error) {
	te, ok := tr.GetEntryForType(localTypeName)
	if !ok {
		return GoType{}, fmt.Errorf("this is weird.... not found type %s in registry", localTypeName)
	}

	gtn := localTypeName
	reqImport := ""
	pkgName := te.ComputePackageName(tr.Cfg.SimpleTypesInCommonPackage)
	if pkgName != currentPkg {
		gtn = strings.Join([]string{pkgName, gtn}, ".")
		reqImport = gm.resolveGoImport(pkgName)
	}

	return GoType{Name: gtn, Import: reqImport}, nil
}

func (gm *GoModel) resolveGoImport(importName string) string {
	if importName == "common" {
		return gm.Cfg.PackageImport("common")
	}

	return importName
}
