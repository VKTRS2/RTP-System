package payment_type

import "github.com/pjover/sam/internal/domain/model/sequence_type"

func (p PaymentType) SequenceType() sequence_type.SequenceType {
	switch p {
	case BankDirectDebit:
		return sequence_type.StandardInvoice
	case BankTransfer:
	case Voucher:
	case Cash:
		return sequence_type.SpecialInvoice
	case Rectification:
		return sequence_type.RectificationInvoice
	}
	return sequence_type.Invalid
}
