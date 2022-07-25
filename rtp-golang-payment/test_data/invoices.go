package test_data

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/payment_type"
)

var lines = []model.Line{
	{
		ProductId:     "TST",
		Units:         1,
		ProductPrice:  11,
		TaxPercentage: 0,
		ChildId:       1850,
	},
	{
		ProductId:     "XXX",
		Units:         3,
		ProductPrice:  5.5,
		TaxPercentage: 0.1,
		ChildId:       1850,
	},
	{
		ProductId:     "YYY",
		Units:         1.5,
		ProductPrice:  5,
		TaxPercentage: 0,
		ChildId:       1850,
	},
}

var InvoiceF100 = model.Invoice{
	Id:          "F-100",
	CustomerId:  148,
	Date:        Date,
	YearMonth:   "2019-05",
	ChildrenIds: []int{1800, 1801},
	Lines:       lines,
	PaymentType: payment_type.BankDirectDebit,
	Note:        "Invoice note",
	Emailed:     false,
	Printed:     false,
	SentToBank:  false,
}

var InvoiceF101 = model.Invoice{
	Id:          "F-101",
	CustomerId:  148,
	Date:        Date,
	YearMonth:   "2019-05",
	ChildrenIds: []int{1801, 1802},
	Lines:       lines,
	PaymentType: payment_type.BankDirectDebit,
	Note:        "Invoice note",
	Emailed:     false,
	Printed:     false,
	SentToBank:  false,
}

var InvoiceF102 = model.Invoice{
	Id:          "F-102",
	CustomerId:  149,
	Date:        Date,
	YearMonth:   "2019-05",
	ChildrenIds: []int{1800, 1801, 1802},
	Lines:       lines,
	PaymentType: payment_type.BankDirectDebit,
	Note:        "Invoice note",
	Emailed:     false,
	Printed:     false,
	SentToBank:  false,
}

var InvoiceF103 = model.Invoice{
	Id:          "F-103",
	CustomerId:  149,
	Date:        Date,
	YearMonth:   "2019-05",
	ChildrenIds: []int{1800},
	Lines:       lines,
	PaymentType: payment_type.BankDirectDebit,
	Note:        "Invoice note",
	Emailed:     false,
	Printed:     false,
	SentToBank:  false,
}
