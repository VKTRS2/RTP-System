package money

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var amountTable = map[string]Amount{
	"22.00":  22.00,
	"123.45": 123.45,
	"123,45": 123.45,
	"0,99":   0.99,

	"1,234.56":       1234.56,
	"1,234,567.89":   1234567.89,
	"10,234,567.89":  10234567.89,
	"100,234,567.89": 100234567.89,

	"1.234,56":       1234.56,
	"1.234.567,89":   1234567.89,
	"10.234.567,89":  10234567.89,
	"100.234.567,89": 100234567.89,

	"-22.00":  -22.00,
	"-123.45": -123.45,
	"-123,45": -123.45,
	"-0,99":   -0.99,

	"-1,234.56":       -1234.56,
	"-1,234,567.89":   -1234567.89,
	"-10,234,567.89":  -10234567.89,
	"-100,234,567.89": -100234567.89,

	"-1.234,56":       -1234.56,
	"-1.234.567,89":   -1234567.89,
	"-10.234.567,89":  -10234567.89,
	"-100.234.567,89": -100234567.89,
}

var invalidAmounts = []string{
	"EUR 123",
	"1234.123.12,0",
	"1,234,56",
	"10,2340,560",
	// "1000,234,560",
}

var nonStandardFormatted = map[string]Amount{
	"3":            3,
	"666":          666,
	"123,4":        123.4,
	"123.4":        123.4,
	"1.2345":       1.2345,
	"1,2345":       1.2345,
	"1000000.8989": 1000000.8989,
	"1000000,8989": 1000000.8989,
	"1,234,567":    1234567,
	"1.234.567":    1234567,
	"1845.1800":    1845.1800,
	"-153.8900":    -153.8900,
	"-91.1000":     -91.1000,
}

func Test_ParseAmount(t *testing.T) {
	for str, refAmount := range amountTable {
		amount, err := ParseAmount(str, 2)
		if err != nil {
			t.Errorf("Could not parse amount %s because of error: '%s'", str, err)
		}
		if amount != refAmount {
			t.Errorf("Parsed '%s' amount %f != %f", str, amount, refAmount)
		}
	}
	for _, str := range invalidAmounts {
		amount, err := ParseAmount(str, 2)
		if err == nil {
			t.Errorf("Parsed invalid amount '%s' as %f", str, amount)
		}
	}
	for str, refAmount := range nonStandardFormatted {
		amount, err := ParseAmount(str)
		if err != nil {
			t.Errorf("Could not parse amount %s because of error: '%s'", str, err)
		}
		if amount != refAmount {
			t.Errorf("Parsed '%s' amount %f != %f", str, amount, refAmount)
		}
	}
}

func Test_StringIsAmount(t *testing.T) {
	for str := range amountTable {
		if !StringIsAmount(str, false) {
			t.Errorf("String not detected as amount: '%s'", str)
		}
	}
	for _, str := range invalidAmounts {
		if StringIsAmount(str, false) {
			t.Errorf("Invalid string detected as amount: '%s'", str)
		}
	}
}

var stringTable = map[Amount]string{
	0:           "0.00",
	0.99:        "0.99",
	-0.99:       "-0.99",
	1:           "1.00",
	-1:          "-1.00",
	20:          "20.00",
	166:         "166.00",
	1.01:        "1.01",
	1.05:        "1.05",
	123456:      "123456.00",
	123456.789:  "123456.79",
	-123456.789: "-123456.79",
	0.055123:    "0.06",
	0.054123:    "0.05",
}

func Test_Amount_String(t *testing.T) {
	for amount, refstr := range stringTable {
		str := amount.String()
		if str != refstr {
			t.Errorf("%v to string is '%s' but should be '%s'", float64(amount), str, refstr)
		}
	}
}

func Test_Amount_RoundToCents(t *testing.T) {
	withoutRoudingError := Amount(137.89)
	withRoundingError := Amount(137.89000000000001)
	assert.NotEqual(t, withoutRoudingError, withRoundingError)

	// Example from production code:
	var amount Amount = 137.89
	var discountPercent int // 0
	var fee Amount          // 0
	total := (amount - (amount * (Amount(discountPercent) / 100)) + fee)
	assert.Equal(t, total, total.RoundToCents())

	// Create always the same pseude random list of integers
	r := rand.New(rand.NewSource(9371))
	testIntegers := make([]int, 1000)
	for i := range testIntegers {
		testIntegers[i] = r.Intn(100000)
		if i&1 == 1 {
			testIntegers[i] = -testIntegers[i]
		}
	}
	// Create all possible cent amounts
	allPossibleCents := make([]int, 100)
	for i := range allPossibleCents {
		allPossibleCents[i] = i
	}

	for _, integer := range testIntegers {
		for _, cents := range allPossibleCents {
			refAmount := Amount(integer) + Amount(cents)*Amount(0.01)
			assert.Equal(t, refAmount, refAmount.RoundToCents())
		}
	}

	roundToCentsTable := map[Amount]Amount{
		137.89000000000001: 137.89,
		0.001:              0,
		0.004:              0,
		0.005:              0.01,
		0.009:              0.01,
		0.099:              0.1,
		9999999.001:        9999999,
		9999999.004:        9999999,
		9999999.005:        9999999.01,
		9999999.009:        9999999.01,
		19999999.55:        19999999.55,
		99.999:             100,
		89.99999:           90,
		189.99999:          190,
		123456789.9999:     123456790,
	}
	for testAmount, refAmount := range roundToCentsTable {
		assert.Equal(t, refAmount, testAmount.RoundToCents())
	}
}

func Test_Amount_SplitEquallyRoundToCents(t *testing.T) {
	type input struct {
		amount     Amount
		numAmounts int
	}
	data := map[input][]Amount{
		{amount: 100, numAmounts: 0}:     nil,
		{amount: 100.005, numAmounts: 1}: {100.01},
		{amount: 100.005, numAmounts: 2}: {50, 50.01},
		{amount: 100, numAmounts: 3}:     {33.33, 33.33, 33.34},
		{amount: 0.01, numAmounts: 5}:    {0, 0, 0, 0, 0.01},
		{amount: 0.05, numAmounts: 5}:    {0.01, 0.01, 0.01, 0.01, 0.01},
		{amount: 1, numAmounts: 17}:      {0.06, 0.06, 0.06, 0.06, 0.06, 0.06, 0.06, 0.06, 0.06, 0.06, 0.06, 0.06, 0.06, 0.06, 0.06, 0.06, 0.04},

		{amount: -100, numAmounts: 0}:     nil,
		{amount: -100.005, numAmounts: 1}: {-100.01},
		{amount: -100.005, numAmounts: 2}: {-50, -50.01},
		{amount: -100, numAmounts: 3}:     {-33.33, -33.33, -33.34},
		{amount: -0.01, numAmounts: 5}:    {-0, -0, -0, -0, -0.01},
		{amount: -0.05, numAmounts: 5}:    {-0.01, -0.01, -0.01, -0.01, -0.01},
		{amount: -1, numAmounts: 17}:      {-0.06, -0.06, -0.06, -0.06, -0.06, -0.06, -0.06, -0.06, -0.06, -0.06, -0.06, -0.06, -0.06, -0.06, -0.06, -0.06, -0.04},
	}
	for input, expected := range data {
		result := input.amount.SplitEquallyRoundToCents(input.numAmounts)
		assert.Equal(t, expected, result)
	}
}

func Test_Amount_ScaleAmountsToSumRoundToCents(t *testing.T) {
	data := []struct {
		sum      Amount
		amounts  []Amount
		expected []Amount
	}{
		{0, nil, nil},
		{0, []Amount{1}, []Amount{0}},
		{100, []Amount{1}, []Amount{100}},
		{100, []Amount{50.4444345}, []Amount{100}},
		{100, []Amount{100}, []Amount{100}},
		{100, []Amount{1000}, []Amount{100}},
		{100, []Amount{1, 2}, []Amount{33.33, 66.67}},
		{100, []Amount{1, 2, 3}, []Amount{16.67, 33.33, 50}},
		{100, []Amount{1, 2, 3, 4}, []Amount{10, 20, 30, 40}},
		{100, []Amount{-1, +2, -3, +4}, []Amount{10, 20, 30, 40}},
		{100, []Amount{11, 17, 37}, []Amount{16.92, 26.15, 56.93}},
		{1, []Amount{11, 17, 37}, []Amount{0.17, 0.26, 0.57}},

		{-0, nil, nil},
		{-0, []Amount{-1}, []Amount{-0}},
		{-100, []Amount{-1}, []Amount{-100}},
		{-100, []Amount{-50.4444345}, []Amount{-100}},
		{-100, []Amount{-100}, []Amount{-100}},
		{-100, []Amount{-1000}, []Amount{-100}},
		{-100, []Amount{-1, 2}, []Amount{-33.33, -66.67}},
		{-100, []Amount{-1, 2, 3}, []Amount{-16.67, -33.33, -50}},
		{-100, []Amount{-1, 2, 3, 4}, []Amount{-10, -20, -30, -40}},
		{-100, []Amount{-1, +2, -3, +4}, []Amount{-10, -20, -30, -40}},
		{-100, []Amount{-11, 17, 37}, []Amount{-16.92, -26.15, -56.93}},
		{-1, []Amount{-11, 17, 37}, []Amount{-0.17, -0.26, -0.57}},
	}
	for _, test := range data {
		result := ScaleAmountsToSumRoundToCents(test.amounts, test.sum)
		assert.Equal(t, test.expected, result)
	}
}
