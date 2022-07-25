package docmapper

import (
	"encoding/json"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/vars"
	"github.com/PaesslerAG/gval"
	"github.com/rs/zerolog/log"
)

func NewMapperClass(msgName, docClass string, opts ...Option) (MappingClass, error) {

	mc := MappingClass{MsgName: msgName, Name: docClass, FuncMap: GetFuncMap()}
	for _, o := range opts {
		_ = o(&mc)
	}

	return mc, nil
}

func NewMapperClassFromYAML(data []byte, opts ...Option) (MappingClass, error) {

	mc := MappingClass{FuncMap: GetFuncMap()}
	for _, o := range opts {
		_ = o(&mc)
	}

	err := json.Unmarshal(data, &mc)
	if err != nil {
		return mc, err
	}

	return mc, nil
}

func WithFuncMap(funcMap FuncMap) Option {
	return func(mc *MappingClass) error {
		if funcMap != nil {
			m := make(FuncMap)
			for n, v := range funcMap {
				m[n] = v
			}

			for n, v := range GetFuncMap() {
				m[n] = v
			}

			mc.FuncMap = m
		} else {
			mc.FuncMap = GetFuncMap()
		}

		return nil
	}
}

func (mc *MappingClass) GetResolver(opts ...ResolverOption) (*Resolver, error) {

	/*	opts := []ResolverOption{
		WithMappableDoc(sourceMap),
	}*/

	r := &Resolver{}
	for _, o := range opts {
		err := o(r)
		if err != nil {
			return r, err
		}
	}

	return r, nil
}

func (mc *MappingClass) Map(sourceDoc MappableDocument, targetDoc MappableDocument, verbose bool) error {

	resolver, err := mc.GetResolver(WithMappableDoc(sourceDoc))
	if err != nil {
		return err
	}

	for _, r := range mc.Rules {
		err := r.apply(targetDoc, mc.FuncMap, resolver, verbose)
		if err != nil {
			return err
		}
	}

	return nil
}

func (mr *MappingRule) apply(targetDoc MappableDocument, funcs FuncMap, resolver *Resolver, verbose bool) error {

	if verbose {
		log.Trace().Str("name", mr.Name).Str("guard", mr.Guard).Str("map-func", mr.MapFunc).Str("source-value", mr.SourceValue).Str("target-path", mr.TargetPath).Bool("is-expr", mr.IsExpr).Str("", mr.Name).Msg("mapping rule")
	}
	if mr.Guard != "" {

		guardExpr := false
		if isExpression(mr.Guard) {
			guardExpr = true
		}

		s, err := varResolver.ResolveVariables(mr.Guard, varResolver.SimpleVariableReference, resolver.ResolveVar, true)
		if err != nil {
			return err
		}

		if guardExpr && s != "" {
			exprVal, err := gval.Evaluate(s, funcs)
			if err != nil {
				log.Error().Err(err).Str("expr", s).Msg("error in expression evaluation")
			}

			b, ok := exprVal.(bool)
			if !ok || !b {
				return nil
			}
		} else {
			if s == "" {
				return nil
			}
		}
	}

	s, err := varResolver.ResolveVariables(mr.SourceValue, varResolver.SimpleVariableReference, resolver.ResolveVar, true)
	if err != nil {
		return err
	}

	if mr.IsExpr && s != "" {
		exprVal, err := gval.Evaluate(s, funcs)
		if err != nil {
			log.Error().Err(err).Str("expr", s).Msg("error in expression evaluation")
		}

		s = fmt.Sprintf("%v", exprVal)
	}

	if mr.MapFunc != "" && s != "" {
		f, ok := funcs[mr.MapFunc]
		if ok {
			if fun, ok := f.(func(string) string); ok {
				s = fun(s)
			} else {
				log.Warn().Str("func-name", mr.MapFunc).Msg("mapping function name found but not applicable")
			}
		} else {
			log.Warn().Str("func-name", mr.MapFunc).Msg("cannot find function map")
		}
	}

	if verbose {
		log.Trace().Str("name", mr.Name).Str("output", s).Str("target-path", mr.TargetPath).Msg("mapping rule evaluated")
	}

	if s != "" {
		return targetDoc.Set(mr.TargetPath, s)
	}

	return nil
}
