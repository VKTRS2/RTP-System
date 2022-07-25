package docmapper

import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"

type MappingRule struct {
	Name        string `mapstructure:"name,omitempty" yaml:"name,omitempty" json:"name,omitempty"`
	SourceValue string `mapstructure:"source-value,omitempty yaml:"source-value,omitempty" json:"source-value,omitempty"`
	IsExpr      bool   `mapstructure:"is-expr,omitempty" yaml:"is-expr,omitempty" json:"is-expr,omitempty"`
	Guard       string `mapstructure:"guard,omitempty" yaml:"guard,omitempty" json:"guard,omitempty"`
	MapFunc     string `mapstructure:"map-func,omitempty" yaml:"map-func,omitempty" json:"map-func,omitempty"`
	TargetPath  string `mapstructure:"target-path,omitempty" yaml:"target-path,omitempty" json:"target-path,omitempty"`
}

type MappingFunction func(s string) string

type Option func(mc *MappingClass) error

type MappingClass struct {
	Name    string        `mapstructure:"name" yaml:"name"  json:"name"`
	MsgName string        `mapstructure:"msg-name" yaml:"msg-name"  json:"msg-name"`
	Rules   []MappingRule `mapstructure:"rules" yaml:"rules" json:"rules"`
	FuncMap FuncMap       `mapstructure:"-" yaml:"-" json:"-"`
}

type MappableDocument interface {
	Set(docPath string, src interface{}) error
	Get(docPath string) (interface{}, error)
}

type MappableDocMap map[string]interface{}

func (m MappableDocMap) Set(docPath string, src interface{}) error {
	return util.SetProperty(m, docPath, src)
}

func (m MappableDocMap) Get(docPath string) (interface{}, error) {
	return util.GetProperty(m, docPath), nil
}
