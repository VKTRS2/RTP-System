package docmapper

import (
	"embed"
	"gopkg.in/yaml.v3"
	"path"
)

var Registry map[string]MappingClass

const YamlExtension = ".yml"

func InitEmbeddedRegistry(docClassFS embed.FS, docClassFSRootPath string) (int, error) {

	Registry = make(map[string]MappingClass, 0)
	dirEntries, err := docClassFS.ReadDir(docClassFSRootPath)
	if err != nil {
		return len(Registry), err
	}

	for _, e := range dirEntries {

		fn := e.Name()

		if path.Ext(fn) != YamlExtension {
			continue
		}

		dc, err := readDocClassYMLDefinition(docClassFS, docClassFSRootPath, fn)
		if err != nil {
			return len(Registry), err
		}

		Registry[dc.Name] = dc
	}

	return len(Registry), nil
}

func readDocClassYMLDefinition(docClassFS embed.FS, docClassFSRootPath string, fn string) (MappingClass, error) {

	fileContent, err := docClassFS.ReadFile(path.Join(docClassFSRootPath, fn))
	if err != nil {
		return MappingClass{}, err
	}

	dc, _ := NewMapperClass("", "")
	err = yaml.Unmarshal(fileContent, &dc)
	if err != nil {
		return dc, err
	}

	for i, r := range dc.Rules {
		if isExpression(r.SourceValue) {
			dc.Rules[i].IsExpr = true
		}
	}

	return dc, nil
}
