package docmapper

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

type ResolverOption func(r *Resolver) error

type Resolver struct {
	doc MappableDocument
}

func WithMappableDoc(doc MappableDocument) ResolverOption {
	return func(r *Resolver) error {
		r.doc = doc
		return nil
	}
}

var resolverTypePrefix = []string{"$."}

func (pvr *Resolver) ResolveVar(s string) string {

	pfix, err := pvr.getPrefix(s)
	if err != nil {
		return ""
	}

	switch pfix {
	case "$.":
		v, err := pvr.doc.Get(strings.TrimPrefix(s, "$."))
		if err == nil {
			if v == nil {
				return ""
			}

			return fmt.Sprintf("%v", v)
		}

		log.Error().Err(err).Str("path-name", s).Msg("json-path error")
	default:
		v, ok := os.LookupEnv(s)
		if ok {
			return v
		}
	}

	log.Error().Str("var-name", s).Msg("could not resolve variable")
	return ""
}

func (pvr *Resolver) getPrefix(s string) (string, error) {

	matchedPrefix := "env"
	for _, pfix := range resolverTypePrefix {
		if strings.HasPrefix(s, pfix) {
			matchedPrefix = pfix
			break
		}
	}

	isValid := false
	switch matchedPrefix {
	case "$.":
		if pvr.doc != nil {
			isValid = true
		}

	case "env":
		isValid = true
	}

	if !isValid {
		return matchedPrefix, fmt.Errorf("found prefix but resover doesn't have data for resolving")
	}

	return matchedPrefix, nil
}
