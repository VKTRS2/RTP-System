// Package pain_013_001_07
// Do not Edit. This stuff it's been automatically generated.
package pain_013_001_07

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
)

// Cheque11 type definition
type Cheque11 struct {
	ChqTp       common.ChequeType2Code       `xml:"ChqTp,omitempty"`
	ChqNb       common.Max35Text             `xml:"ChqNb,omitempty"`
	ChqFr       *NameAndAddress16            `xml:"ChqFr,omitempty"`
	DlvryMtd    *ChequeDeliveryMethod1Choice `xml:"DlvryMtd,omitempty"`
	DlvrTo      *NameAndAddress16            `xml:"DlvrTo,omitempty"`
	InstrPrty   common.Priority2Code         `xml:"InstrPrty,omitempty"`
	ChqMtrtyDt  common.ISODate               `xml:"ChqMtrtyDt,omitempty"`
	FrmsCd      common.Max35Text             `xml:"FrmsCd,omitempty"`
	MemoFld     []common.Max35Text           `xml:"MemoFld,omitempty"`
	RgnlClrZone common.Max35Text             `xml:"RgnlClrZone,omitempty"`
	PrtLctn     common.Max35Text             `xml:"PrtLctn,omitempty"`
	Sgntr       []common.Max70Text           `xml:"Sgntr,omitempty"`
}

// InstructionForCreditorAgent1 type definition
type InstructionForCreditorAgent1 struct {
	Cd       common.Instruction3Code `xml:"Cd,omitempty"`
	InstrInf common.Max140Text       `xml:"InstrInf,omitempty"`
}

// ChequeDeliveryMethod1Choice type definition
type ChequeDeliveryMethod1Choice struct {
	Cd    common.ChequeDelivery1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text           `xml:"Prtry,omitempty"`
}

// PaymentInstruction31 type definition
type PaymentInstruction31 struct {
	PmtInfId    common.Max35Text                                    `xml:"PmtInfId,omitempty"`
	PmtMtd      common.PaymentMethod7Code                           `xml:"PmtMtd"`
	PmtTpInf    *common.PaymentTypeInformation26                    `xml:"PmtTpInf,omitempty"`
	ReqdExctnDt common.DateAndDateTime2Choice                       `xml:"ReqdExctnDt"`
	XpryDt      *common.DateAndDateTime2Choice                      `xml:"XpryDt,omitempty"`
	PmtCond     *common.PaymentCondition1                           `xml:"PmtCond,omitempty"`
	Dbtr        common.PartyIdentification135                       `xml:"Dbtr"`
	DbtrAcct    *common.CashAccount38                               `xml:"DbtrAcct,omitempty"`
	DbtrAgt     common.BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt"`
	UltmtDbtr   *common.PartyIdentification135                      `xml:"UltmtDbtr,omitempty"`
	ChrgBr      common.ChargeBearerType1Code                        `xml:"ChrgBr,omitempty"`
	CdtTrfTx    []CreditTransferTransaction35                       `xml:"CdtTrfTx"`
}

// StructuredRegulatoryReporting3 type definition
type StructuredRegulatoryReporting3 struct {
	Tp   common.Max35Text                          `xml:"Tp,omitempty"`
	Dt   common.ISODate                            `xml:"Dt,omitempty"`
	Ctry common.CountryCode                        `xml:"Ctry,omitempty"`
	Cd   common.Max10Text                          `xml:"Cd,omitempty"`
	Amt  *common.ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
	Inf  []common.Max35Text                        `xml:"Inf,omitempty"`
}

// CreditorPaymentActivationRequestV07 type definition
type CreditorPaymentActivationRequestV07 struct {
	GrpHdr      GroupHeader78               `xml:"GrpHdr"`
	PmtInf      []PaymentInstruction31      `xml:"PmtInf"`
	SplmtryData []common.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

// GroupHeader78 type definition
type GroupHeader78 struct {
	MsgId    common.Max35Text              `xml:"MsgId"`
	CreDtTm  common.ISODateTime            `xml:"CreDtTm"`
	NbOfTxs  common.Max15NumericText       `xml:"NbOfTxs"`
	CtrlSum  xsdt.Decimal                  `xml:"CtrlSum,omitempty"`
	InitgPty common.PartyIdentification135 `xml:"InitgPty"`
}

// RemittanceLocationData1 type definition
type RemittanceLocationData1 struct {
	Mtd        common.RemittanceLocationMethod2Code `xml:"Mtd"`
	ElctrncAdr common.Max2048Text                   `xml:"ElctrncAdr,omitempty"`
	PstlAdr    *NameAndAddress16                    `xml:"PstlAdr,omitempty"`
}

// RemittanceLocation7 type definition
type RemittanceLocation7 struct {
	RmtId       common.Max35Text          `xml:"RmtId,omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `xml:"RmtLctnDtls,omitempty"`
}

// TaxInformation8 type definition
type TaxInformation8 struct {
	Cdtr            *common.TaxParty1                         `xml:"Cdtr,omitempty"`
	Dbtr            *common.TaxParty2                         `xml:"Dbtr,omitempty"`
	AdmstnZone      common.Max35Text                          `xml:"AdmstnZone,omitempty"`
	RefNb           common.Max140Text                         `xml:"RefNb,omitempty"`
	Mtd             common.Max35Text                          `xml:"Mtd,omitempty"`
	TtlTaxblBaseAmt *common.ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       *common.ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`
	Dt              common.ISODate                            `xml:"Dt,omitempty"`
	SeqNb           xsdt.Decimal                              `xml:"SeqNb,omitempty"`
	Rcrd            []common.TaxRecord2                       `xml:"Rcrd,omitempty"`
}

// PaymentIdentification6 type definition
type PaymentIdentification6 struct {
	InstrId    common.Max35Text        `xml:"InstrId,omitempty"`
	EndToEndId common.Max35Text        `xml:"EndToEndId"`
	UETR       common.UUIDv4Identifier `xml:"UETR,omitempty"`
}

// NameAndAddress16 type definition
type NameAndAddress16 struct {
	Nm  common.Max140Text      `xml:"Nm"`
	Adr common.PostalAddress24 `xml:"Adr"`
}

// RegulatoryReporting3 type definition
type RegulatoryReporting3 struct {
	DbtCdtRptgInd common.RegulatoryReportingType1Code `xml:"DbtCdtRptgInd,omitempty"`
	Authrty       *RegulatoryAuthority2               `xml:"Authrty,omitempty"`
	Dtls          []StructuredRegulatoryReporting3    `xml:"Dtls,omitempty"`
}

// RegulatoryAuthority2 type definition
type RegulatoryAuthority2 struct {
	Nm   common.Max140Text  `xml:"Nm,omitempty"`
	Ctry common.CountryCode `xml:"Ctry,omitempty"`
}

// CreditTransferTransaction35 type definition
type CreditTransferTransaction35 struct {
	PmtId           PaymentIdentification6                               `xml:"PmtId"`
	PmtTpInf        *common.PaymentTypeInformation26                     `xml:"PmtTpInf,omitempty"`
	PmtCond         *common.PaymentCondition1                            `xml:"PmtCond,omitempty"`
	Amt             common.AmountType4Choice                             `xml:"Amt"`
	ChrgBr          common.ChargeBearerType1Code                         `xml:"ChrgBr"`
	ChqInstr        *Cheque11                                            `xml:"ChqInstr,omitempty"`
	UltmtDbtr       *common.PartyIdentification135                       `xml:"UltmtDbtr,omitempty"`
	IntrmyAgt1      *common.BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty"`
	IntrmyAgt2      *common.BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty"`
	IntrmyAgt3      *common.BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty"`
	CdtrAgt         common.BranchAndFinancialInstitutionIdentification6  `xml:"CdtrAgt"`
	Cdtr            common.PartyIdentification135                        `xml:"Cdtr"`
	CdtrAcct        *common.CashAccount38                                `xml:"CdtrAcct,omitempty"`
	UltmtCdtr       *common.PartyIdentification135                       `xml:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt []InstructionForCreditorAgent1                       `xml:"InstrForCdtrAgt,omitempty"`
	Purp            *common.Purpose2Choice                               `xml:"Purp,omitempty"`
	RgltryRptg      []RegulatoryReporting3                               `xml:"RgltryRptg,omitempty"`
	Tax             *TaxInformation8                                     `xml:"Tax,omitempty"`
	RltdRmtInf      []RemittanceLocation7                                `xml:"RltdRmtInf,omitempty"`
	RmtInf          *common.RemittanceInformation16                      `xml:"RmtInf,omitempty"`
	NclsdFile       []common.Document12                                  `xml:"NclsdFile,omitempty"`
	SplmtryData     []common.SupplementaryData1                          `xml:"SplmtryData,omitempty"`
}
