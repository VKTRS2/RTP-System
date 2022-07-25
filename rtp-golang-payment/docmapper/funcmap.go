package docmapper

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/docmapper/funcs"
	"strings"
)

type FuncMap map[string]interface{}

func GetFuncMap() map[string]interface{} {
	builtins := map[string]interface{}{
		"now":       funcs.Now,
		"rtpId":     funcs.RtpId,
		"uuid":      funcs.NewUUID,
		"objId":     funcs.NewObjectId,
		"decimal":   funcs.Decimal,
		"nop":       func(s string) string { return s },
		"trimSpace": func(s string) string { return strings.TrimSpace(s) },
	}

	return builtins
}

var expressionSmell = []string{
	"now",
	"rtpId",
	"uuid",
	"decimal",
	">",
	"<",
	"(",
	")",
	"=",
	"\"",
}

// isExpression In order not to clutter the process vars assignments in simple cases.... try to detect if this is an expression or not.
// didn't parse the thing but try to find if there is any 'reserved' word in there.
// example: 'hello' is not an expression, '"hello"' is an expression which evaluates to 'hello'. This trick is to avoid something like
// value: '"{$.operazione.commissione}"' in the yamls. Someday I'll get to there.... sure...
func isExpression(e string) bool {
	if e == "" {
		return false
	}

	for _, s := range expressionSmell {
		if strings.Contains(e, s) {
			return true
		}
	}

	return false
}
