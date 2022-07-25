package pain001

import (
	"encoding/xml"
)

// CustomerCreditTransferInitiation contains all the needed information.
type CustomerCreditTransferInitiation struct {
	XMLName            xml.Name           `xml:"CstmrCdtTrfInitn"`
	GroupHeader        GroupHeader        `xml:"GrpHdr"`
	PaymentInformation PaymentInformation `xml:"PmtInf"`
}

// NewCustomerCreditTransferInitiation takes the name of the sender, the number of transactions and the sum
// of all transactions and returns a CustomerCreditTransferInitiation.
func NewCustomerCreditTransferInitiation(order Order) (CustomerCreditTransferInitiation, error) {
	sum, err := order.transactionSum()
	if err != nil {
		return CustomerCreditTransferInitiation{}, err
	}
	return CustomerCreditTransferInitiation{
		GroupHeader:        NewGroupHeader(order.Debitor.Name, len(order.Transactions), sum),
		PaymentInformation: NewPaymentInformation(order),
	}, nil
}
