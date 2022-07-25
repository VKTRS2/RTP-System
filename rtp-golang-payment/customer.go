package model

import (
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"github.com/pjover/sam/internal/domain/model/group_type"
	"github.com/pjover/sam/internal/domain/model/language"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"time"
)

type Child struct {
	Id            int
	Name          string
	Surname       string
	SecondSurname string
	TaxID         string
	BirthDate     time.Time
	Group         group_type.GroupType
	Note          string
	Active        bool
}

type Adult struct {
	Name             string
	Surname          string
	SecondSurname    string
	TaxID            string
	Role             adult_role.AdultRole
	Address          Address
	Email            string
	MobilePhone      string
	HomePhone        string
	GrandMotherPhone string
	GrandParentPhone string
	WorkPhone        string
	BirthDate        time.Time
	Nationality      string
}

type Address struct {
	Street  string
	ZipCode string
	City    string
	State   string
}

type InvoiceHolder struct {
	Name        string
	TaxID       string
	Address     Address
	Email       string
	SendEmail   bool
	PaymentType payment_type.PaymentType
	BankAccount string
	IsBusiness  bool
}

type Customer struct {
	Id            int
	Active        bool
	Children      []Child
	Adults        []Adult
	InvoiceHolder InvoiceHolder
	Note          string
	Language      language.Language
	ChangedOn     time.Time
}
