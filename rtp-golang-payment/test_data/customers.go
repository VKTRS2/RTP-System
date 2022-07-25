package test_data

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"github.com/pjover/sam/internal/domain/model/group_type"
	"github.com/pjover/sam/internal/domain/model/language"
	"github.com/pjover/sam/internal/domain/model/payment_type"
)

var Child1850 = model.Child{
	Id:            1850,
	Name:          "Laura",
	Surname:       "Llull",
	SecondSurname: "Bibiloni",
	BirthDate:     Date,
	Group:         group_type.EI_1,
	Active:        true,
}

var Child1851 = model.Child{
	Id:            1851,
	Name:          "Aina",
	Surname:       "Llull",
	SecondSurname: "Bibiloni",
	TaxID:         "60235657Z",
	BirthDate:     Date,
	Group:         group_type.EI_1,
	Active:        true,
}

var AdultMare = model.Adult{
	Name:          "Nom de la mare",
	Surname:       "1er llinatge_mare",
	SecondSurname: "2on llinatge_mare",
	TaxID:         "36361882D",
	Role:          adult_role.Mother,
}

var AdultPare = model.Adult{
	Name:          "Nom de la pare",
	Surname:       "1er llinatge_pare",
	SecondSurname: "2on llinatge_pare",
	TaxID:         "71032204Q",
	Role:          adult_role.Father,
}

var InvoiceHolder148 = model.InvoiceHolder{
	Name:  "Nom de la mare 1er llinatge_mare 2on llinatge_mare",
	TaxID: "36361882D",
	Address: model.Address{
		Street:  "Address first line",
		ZipCode: "07007",
		City:    "Palma",
		State:   "Illes Balears",
	},
	Email:       "email@gmail.com",
	PaymentType: payment_type.BankDirectDebit,
	BankAccount: "ES2830668859978258529057",
}

var InvoiceHolder149 = model.InvoiceHolder{
	Name:  "Nom empresa",
	TaxID: "37866397W",
	Address: model.Address{
		Street:  "Address first line",
		ZipCode: "07007",
		City:    "Palma",
		State:   "Illes Balears",
	},
	Email:       "email@gmail.com",
	PaymentType: payment_type.BankTransfer,
	BankAccount: "ES2830668859978258529057",
	IsBusiness:  true,
}

var Customer148 = model.Customer{
	Id:     148,
	Active: true,
	Children: []model.Child{
		Child1850,
		Child1851,
	},
	Adults: []model.Adult{
		AdultMare,
		AdultPare,
	},
	InvoiceHolder: InvoiceHolder148,
	Note:          "Nota del client",
	Language:      language.Catalan,
}

var Customer149 = model.Customer{
	Id:     149,
	Active: true,
	Children: []model.Child{
		Child1850,
		Child1851,
	},
	Adults: []model.Adult{
		AdultMare,
		AdultPare,
	},
	InvoiceHolder: InvoiceHolder149,
	Note:          "Nota del client",
	Language:      language.Catalan,
}
