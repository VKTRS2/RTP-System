package model

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

// GetTypes Method
func (gm *GoModel) GetTypes(pkgName string, structDefs bool) []GoTypeDefinition {

	var types []GoTypeDefinition
	for _, td := range gm.TypeDefs {
		if td.PackageName == pkgName && td.StructDef == structDefs {
			types = append(types, td)
		}
	}

	return types
}

// CountTypes Method
func (gm *GoModel) CountTypes(pkgName string, structDefs bool) int {

	count := 0
	for _, td := range gm.TypeDefs {
		if td.PackageName == pkgName && td.StructDef == structDefs {
			count++
		}
	}

	return count
}

// GetImports Method
func (gm *GoModel) GetImports(pkgName string, structDef bool) []string {

	importsMap := make(map[string]struct{})
	for _, td := range gm.TypeDefs {

		if td.StructDef != structDef {
			continue
		}

		if td.PackageName == pkgName {
			if td.Type.Import != "" {
				if _, ok := importsMap[td.Type.Import]; !ok {
					importsMap[td.Type.Import] = struct{}{}
				}
			}

			for _, a := range td.Attrs {
				if a.Type.Import != "" {
					if _, ok := importsMap[a.Type.Import]; !ok {
						importsMap[a.Type.Import] = struct{}{}
					}
				}
			}
		}

	}

	var imports []string
	for n, _ := range importsMap {
		imports = append(imports, n)
	}
	return imports
}

// ShowInfo Method
func (gm *GoModel) ShowInfo() {

	// It's kind of three loops
	for _, msg := range gm.Msgs {
		log.Trace().Str("msg", msg.Name).Str("pkg", msg.PackageName()).Msg("package information")
		log.Trace().Msg("composite defs")
		types := gm.GetTypes(msg.PackageName(), true)
		showTypeDefinitions(types)

		log.Trace().Msg("simple information")
		types = gm.GetTypes(msg.PackageName(), false)
		showTypeDefinitions(types)
	}

	log.Trace().Msg("common composite defs")
	types := gm.GetTypes("common", true)
	showTypeDefinitions(types)

	log.Trace().Msg("common simple information")
	types = gm.GetTypes("common", false)
	showTypeDefinitions(types)
}

func showTypeDefinitions(types []GoTypeDefinition) {
	for _, td := range types {
		log.Trace().Str("name", td.Name).Str("pkg", td.PackageName).Str("type", td.Type.Name).Str("import", td.Type.Import).Msg("### type def")
		for _, a := range td.Attrs {
			log.Trace().Str("name", a.Name).Bool("built-in", a.Type.IsBuiltin).Str("type", a.Type.Name).Interface("struct,array,opt,xml-attr", []bool{a.StructType, a.Array, a.Optional, a.XMLAttr}).Str("import", a.Type.Import).Msg("--- attr def")
		}
	}
}

func (gm *GoModel) VisitDocument(pkgName string, visitor GoModelVisitor) error {

	var rootEntry GoTypeDefinition
	mapsOfTypes := make(map[string]GoTypeDefinition)
	for _, td := range gm.TypeDefs {

		// The names in the map are package scoped.
		if td.PackageName == pkgName {
			mapsOfTypes[strings.Join([]string{td.PackageName, td.Name}, ".")] = td
		}

		if td.PackageName == "common" {
			mapsOfTypes[strings.Join([]string{td.PackageName, td.Name}, ".")] = td
		}

		if td.PackageName == pkgName && td.Name == "Document" {
			rootEntry = td
		}
	}

	if rootEntry.Name == "" {
		return fmt.Errorf("cannot find the root element %s for package %s", "Document", pkgName)
	}

	_, _ = visitor.Visit("_self", strings.Join([]string{pkgName, "Document"}, "."), true, false, false)

	for _, a := range rootEntry.Attrs {
		// The container package is passed since the attributes have a type that is already scoped or not. If it is not
		// scoped it means that the type is in the same package as the container.... so there is the need to propagate.
		if err := gm.VisitAttr(visitor, mapsOfTypes, rootEntry.PackageName, a); err != nil {
			return err
		}
	}

	return nil
}

func (gm *GoModel) VisitAttr(visitor GoModelVisitor, types map[string]GoTypeDefinition, parentPackageName string, el GoAttr) error {

	lookupKey := el.Type.Name
	if strings.Index(el.Type.Name, ".") < 0 {
		lookupKey = strings.Join([]string{parentPackageName, el.Type.Name}, ".")
	}

	_, err := visitor.Visit(el.Name, lookupKey, el.StructType, el.IsPtr(), el.Array)
	if err != nil {
		return err
	}

	// log.Trace().Int("depth", depth).Str("name", el.Name).Str("type", el.Type.Name).Bool("isPtr", el.IsPtr()).Bool("isArray", el.Array).Msg("attribute")
	if el.StructType {
		// strings.Join([]string{el.Type.Name}, ".")
		elTypeDef, ok := types[lookupKey]
		if !ok {
			return fmt.Errorf("cannot find the type %s of element %s", el.Type.Name, el.Name)
		}

		for _, attr := range elTypeDef.Attrs {
			if err := gm.VisitAttr(visitor, types, elTypeDef.PackageName, attr); err != nil {
				return err
			}
		}
	}

	_, err = visitor.Pop()
	return err
}
