// Package pain001 can generate Swiss ISO 20022 Pain.001 payment orders.
// This library does _not_ implement the whole specification. Use this at your own risk.
package pain001

import (
	"fmt"
	"strconv"
)

// Debitor represents the payee (aka originator) of a transaction.
type Debitor struct {
	// Name of the debitor. Full name or company name.
	Name string
	// Street name of the debitor without the number.
	Street string
	// StreetNr is the street number of the debitor.
	StreetNr int
	// Postecode of the debitor.
	Postcode int
	// Place name of the debitor.
	Place string
	// Country code of the debitor (ex.: "CH").
	Country string
	// IBAN code of the debiting bank account.
	IBAN string
	// BIC code of the debitor's bank.
	BIC string
}

// Transaction is one payment order from the debitor to the creditor (aka receiver of the money).
type Transaction struct {
	// Name full name or company name of the creditor.
	Name string
	// Street name of the creditor without the number.
	Street string
	// StreetNr is the street number of the creditor.
	StreetNr int
	// Postecode of the creditor.
	Postcode int
	// Place name of the creditor.
	Place string
	// Country code of the creditor (ex.: "CH").
	Country string
	// IBAN code of the creditors bank account.
	IBAN string
	// Reference is the comment/description the creditor will see upon receiving the transaction.
	Reference string
	// Amount is the amount of money to be transferred.
	Amount string
	// Currency of the transaction (ex: "CHF").
	Currency string
}

// Order represents one or more transactions of one debitor.
type Order struct {
	// ExecuteOn states the date of execution in the `YYYY-MM-DD` format
	ExecuteOn string
	// Debitor the payee of the transactions.
	Debitor Debitor
	// Transactions the payments to be generated.
	Transactions []Transaction
}

// PaymentOrder returns the ISO 20022 Pain.001 XML representation of the order.
func (o Order) PaymentOrder(indent bool) ([]byte, error) {
	doc, err := NewDocument(o)
	if err != nil {
		return []byte{}, err
	}
	data, err := doc.toXml(indent)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// TransactionSum returns the total amount of all transactions as a string formatted floating point number.
func (o Order) transactionSum() (string, error) {
	sum := 0.0
	for i := range o.Transactions {
		val, err := strconv.ParseFloat(o.Transactions[i].Amount, 64)
		if err != nil {
			return "", fmt.Errorf("error converting \"%s\" (amount of transaction \"%s\") to float", o.Transactions[i].Amount, o.Transactions[i].Reference)
		}
		sum += val
	}
	return fmt.Sprintf("%.2f", sum), nil
}
