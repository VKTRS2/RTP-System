package money

import (
	"fmt"
	"strings"
)

type CurrencyAmount struct {
	Currency Currency
	Amount   Amount
}

func NewCurrencyAmount(currency Currency, amount Amount) *CurrencyAmount {
	return &CurrencyAmount{Currency: currency, Amount: amount}
}

// ParseCurrencyAmount parses a currency and an amount from str with acceptedDecimals.
// If acceptedDecimals is empty, then any decimal number is accepted.
func ParseCurrencyAmount(str string, acceptedDecimals ...int) (result CurrencyAmount, err error) {
	str = strings.TrimSpace(str)

	// Find first separator between currency and amount
	if pos := strings.IndexAny(str, " .,'-+0123456789"); pos != -1 {
		// Try parsing string until separator as currency
		result.Currency, err = NormalizeCurrency(str[:pos])
		if err == nil {
			// Set str to remaining amount part
			str = strings.TrimLeft(str[pos:], " ")
		} else {
			// If currency was not at string start, try from end
			pos = strings.LastIndexAny(str, " .,'-+0123456789")
			if pos != -1 && len(str)-pos > 0 {
				result.Currency, err = NormalizeCurrency(str[pos+1:])
				if err == nil {
					// Set str to remaining amount part
					str = strings.TrimRight(str[:pos+1], " ")
				}
			}
		}
	}

	result.Amount, err = ParseAmount(str, acceptedDecimals...)
	if err != nil {
		return CurrencyAmount{}, err
	}

	return result, nil
}

// String implements the fmt.Stringer interface.
func (ca *CurrencyAmount) String() string {
	return fmt.Sprintf("%s %.2f", ca.Currency, ca.Amount)
}

func (ca *CurrencyAmount) Format(currencyFirst bool, thousandsSep, decimalSep byte, precision int) string {
	amountStr := ca.Amount.Format(thousandsSep, decimalSep, precision)
	if ca.Currency == "" {
		return amountStr
	}
	if currencyFirst {
		return string(ca.Currency) + " " + amountStr
	}
	return amountStr + " " + string(ca.Currency)
}

func (ca *CurrencyAmount) GoString() string {
	return fmt.Sprintf("{Currency: %#v, Amount: %#v}", ca.Currency, ca.Amount)
}

// ScanString tries to parse and assign the passed
// source string as value of the implementing type.
// It returns an error if source could not be parsed.
// If the source string could be parsed, but was not
// in the expected normalized format, then false is
// returned for sourceWasNormalized and nil for err.
// ScanString implements the strfmt.Scannable interface.
func (ca *CurrencyAmount) ScanString(source string) (sourceWasNormalized bool, err error) {
	parsed, err := ParseCurrencyAmount(source, 0, 2)
	if err != nil {
		return false, err
	}
	*ca = parsed
	return ca.String() == source, nil
}
