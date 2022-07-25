package sequence_type

type SequenceType uint

const (
	Invalid              SequenceType = iota
	StandardInvoice                   // Standard invoices with bank direct debit payment type
	SpecialInvoice                    // Special invoices, with other payments types
	RectificationInvoice              // Rectification invoices, correct standard or special invoice
	Customer                          // Customers
)

var values = []string{
	"Indefinit",
	"Factura (rebut)",
	"Factura (no rebut)",
	"Rectificació",
	"Cliente",
}

func (s SequenceType) String() string {
	return values[s]
}
