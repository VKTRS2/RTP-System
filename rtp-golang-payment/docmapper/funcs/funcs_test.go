package funcs_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/docmapper/funcs"
	"github.com/stretchr/testify/require"
	"testing"
)

type InputWanted struct {
	funcName string
	input    string
	wanted   string
}

func TestFuncs(t *testing.T) {

	sarr := []InputWanted{
		{funcName: "Decimal", input: "100", wanted: "1.00"},
	}

	var news string
	for _, s := range sarr {
		switch s.funcName {
		case "Decimal":
			news = funcs.Decimal(s.input)
		default:
			t.Fatalf("func %s not present", s.funcName)
		}

		require.Equal(t, s.wanted, news)
	}
}
