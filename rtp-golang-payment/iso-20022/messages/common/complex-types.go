// Package common
// Do not Edit. This stuff it's been automatically generated.
package common

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
)

// TaxAuthorisation1 type definition
type TaxAuthorisation1 struct {
	Titl Max35Text  `xml:"Titl,omitempty"`
	Nm   Max140Text `xml:"Nm,omitempty"`
}

// TaxAmount2 type definition
type TaxAmount2 struct {
	Rate         xsdt.Decimal                       `xml:"Rate,omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails2                `xml:"Dtls,omitempty"`
}

// SupplementaryData1 type definition
type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"Envlp"`
}

// Case5 type definition
type Case5 struct {
	Id             Max35Text     `xml:"Id"`
	Cretr          Party40Choice `xml:"Cretr"`
	ReopCaseIndctn xsdt.Boolean  `xml:"ReopCaseIndctn,omitempty"`
}

// PaymentTypeInformation27 type definition
type PaymentTypeInformation27 struct {
	InstrPrty Priority2Code           `xml:"InstrPrty,omitempty"`
	ClrChanl  ClearingChannel2Code    `xml:"ClrChanl,omitempty"`
	SvcLvl    []ServiceLevel8Choice   `xml:"SvcLvl,omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`
	SeqTp     SequenceType3Code       `xml:"SeqTp,omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty"`
}

// DocumentLineType1Choice type definition
type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                     `xml:"Prtry,omitempty"`
}

// GarnishmentType1Choice type definition
type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                    `xml:"Prtry,omitempty"`
}

// Frequency36Choice type definition
type Frequency36Choice struct {
	Tp     Frequency6Code       `xml:"Tp,omitempty"`
	Prd    *FrequencyPeriod1    `xml:"Prd,omitempty"`
	PtInTm *FrequencyAndMoment1 `xml:"PtInTm,omitempty"`
}

// PersonIdentification13 type definition
type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1          `xml:"DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty"`
}

// EquivalentAmount2 type definition
type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CcyOfTrf ActiveOrHistoricCurrencyCode      `xml:"CcyOfTrf"`
}

// SettlementInstruction7 type definition
type SettlementInstruction7 struct {
	SttlmMtd             SettlementMethod1Code                         `xml:"SttlmMtd"`
	SttlmAcct            *CashAccount38                                `xml:"SttlmAcct,omitempty"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"ClrSys,omitempty"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt,omitempty"`
	InstgRmbrsmntAgtAcct *CashAccount38                                `xml:"InstgRmbrsmntAgtAcct,omitempty"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgtAcct *CashAccount38                                `xml:"InstdRmbrsmntAgtAcct,omitempty"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification6 `xml:"ThrdRmbrsmntAgt,omitempty"`
	ThrdRmbrsmntAgtAcct  *CashAccount38                                `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
}

// CreditorReferenceType1Choice type definition
type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"Cd,omitempty"`
	Prtry Max35Text         `xml:"Prtry,omitempty"`
}

// GenericIdentification30 type definition
type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"Id"`
	Issr    Max35Text              `xml:"Issr"`
	SchmeNm Max35Text              `xml:"SchmeNm,omitempty"`
}

// FrequencyPeriod1 type definition
type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"Tp"`
	CntPerPrd xsdt.Decimal   `xml:"CntPerPrd"`
}

// DateAndDateTime2Choice type definition
type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"Dt,omitempty"`
	DtTm ISODateTime `xml:"DtTm,omitempty"`
}

// DiscountAmountAndType1 type definition
type DiscountAmountAndType1 struct {
	Tp  *DiscountAmountType1Choice        `xml:"Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

// TaxParty2 type definition
type TaxParty2 struct {
	TaxId   Max35Text          `xml:"TaxId,omitempty"`
	RegnId  Max35Text          `xml:"RegnId,omitempty"`
	TaxTp   Max35Text          `xml:"TaxTp,omitempty"`
	Authstn *TaxAuthorisation1 `xml:"Authstn,omitempty"`
}

// TaxRecordDetails2 type definition
type TaxRecordDetails2 struct {
	Prd *TaxPeriod2                       `xml:"Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

// DiscountAmountType1Choice type definition
type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                       `xml:"Prtry,omitempty"`
}

// PartyAndSignature3 type definition
type PartyAndSignature3 struct {
	Pty   PartyIdentification135 `xml:"Pty"`
	Sgntr SkipPayload            `xml:"Sgntr"`
}

// OrganisationIdentificationSchemeName1Choice type definition
type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                               `xml:"Prtry,omitempty"`
}

// CreditorReferenceType2 type definition
type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"CdOrPrtry"`
	Issr      Max35Text                    `xml:"Issr,omitempty"`
}

// DateAndPlaceOfBirth1 type definition
type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"BirthDt"`
	PrvcOfBirth Max35Text   `xml:"PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth CountryCode `xml:"CtryOfBirth"`
}

// AmountType4Choice type definition
type AmountType4Choice struct {
	InstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`
	EqvtAmt  *EquivalentAmount2                 `xml:"EqvtAmt,omitempty"`
}

// GenericAccountIdentification1 type definition
type GenericAccountIdentification1 struct {
	Id      Max34Text                 `xml:"Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    Max35Text                 `xml:"Issr,omitempty"`
}

// CaseAssignment5 type definition
type CaseAssignment5 struct {
	Id      Max35Text     `xml:"Id"`
	Assgnr  Party40Choice `xml:"Assgnr"`
	Assgne  Party40Choice `xml:"Assgne"`
	CreDtTm ISODateTime   `xml:"CreDtTm"`
}

// AmendmentInformationDetails13 type definition
type AmendmentInformationDetails13 struct {
	OrgnlMndtId      Max35Text                                     `xml:"OrgnlMndtId,omitempty"`
	OrgnlCdtrSchmeId *PartyIdentification135                       `xml:"OrgnlCdtrSchmeId,omitempty"`
	OrgnlCdtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlCdtrAgt,omitempty"`
	OrgnlCdtrAgtAcct *CashAccount38                                `xml:"OrgnlCdtrAgtAcct,omitempty"`
	OrgnlDbtr        *PartyIdentification135                       `xml:"OrgnlDbtr,omitempty"`
	OrgnlDbtrAcct    *CashAccount38                                `xml:"OrgnlDbtrAcct,omitempty"`
	OrgnlDbtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlDbtrAgt,omitempty"`
	OrgnlDbtrAgtAcct *CashAccount38                                `xml:"OrgnlDbtrAgtAcct,omitempty"`
	OrgnlFnlColltnDt ISODate                                       `xml:"OrgnlFnlColltnDt,omitempty"`
	OrgnlFrqcy       *Frequency36Choice                            `xml:"OrgnlFrqcy,omitempty"`
	OrgnlRsn         *MandateSetupReason1Choice                    `xml:"OrgnlRsn,omitempty"`
	OrgnlTrckgDays   Exact2NumericText                             `xml:"OrgnlTrckgDays,omitempty"`
}

// OrganisationIdentification29 type definition
type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

// PartyIdentification135 type definition
type PartyIdentification135 struct {
	Nm        Max140Text       `xml:"Nm,omitempty"`
	PstlAdr   *PostalAddress24 `xml:"PstlAdr,omitempty"`
	Id        *Party38Choice   `xml:"Id,omitempty"`
	CtryOfRes CountryCode      `xml:"CtryOfRes,omitempty"`
	CtctDtls  *Contact4        `xml:"CtctDtls,omitempty"`
}

// FinancialInstitutionIdentification18 type definition
type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier               `xml:"BICFI,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                        `xml:"LEI,omitempty"`
	Nm          Max140Text                           `xml:"Nm,omitempty"`
	PstlAdr     *PostalAddress24                     `xml:"PstlAdr,omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty"`
}

// Charges7 type definition
type Charges7 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

// TaxInformation7 type definition
type TaxInformation7 struct {
	Cdtr            *TaxParty1                         `xml:"Cdtr,omitempty"`
	Dbtr            *TaxParty2                         `xml:"Dbtr,omitempty"`
	UltmtDbtr       *TaxParty2                         `xml:"UltmtDbtr,omitempty"`
	AdmstnZone      Max35Text                          `xml:"AdmstnZone,omitempty"`
	RefNb           Max140Text                         `xml:"RefNb,omitempty"`
	Mtd             Max35Text                          `xml:"Mtd,omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`
	Dt              ISODate                            `xml:"Dt,omitempty"`
	SeqNb           xsdt.Decimal                       `xml:"SeqNb,omitempty"`
	Rcrd            []TaxRecord2                       `xml:"Rcrd,omitempty"`
}

// SupplementaryDataEnvelope1 type definition
type SupplementaryDataEnvelope1 struct {
	Item xsdt.AnyType `xml:",any"`
}

// RemittanceInformation16 type definition
type RemittanceInformation16 struct {
	Ustrd []Max140Text                        `xml:"Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"Strd,omitempty"`
}

// RemittanceAmount2 type definition
type RemittanceAmount2 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1           `xml:"DscntApldAmt,omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1                `xml:"TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1              `xml:"AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
}

// GarnishmentType1 type definition
type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"CdOrPrtry"`
	Issr      Max35Text              `xml:"Issr,omitempty"`
}

// GenericPersonIdentification1 type definition
type GenericPersonIdentification1 struct {
	Id      Max35Text                              `xml:"Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    Max35Text                              `xml:"Issr,omitempty"`
}

// ActiveOrHistoricCurrencyAndAmount type definition
type ActiveOrHistoricCurrencyAndAmount struct {
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
	Value xsdt.Decimal                 `xml:",chardata"`
}

// CashAccount38 type definition
type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"Id"`
	Tp   *CashAccountType2Choice      `xml:"Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`
	Nm   Max70Text                    `xml:"Nm,omitempty"`
	Prxy *ProxyAccountIdentification1 `xml:"Prxy,omitempty"`
}

// ClearingSystemIdentification2Choice type definition
type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                                 `xml:"Prtry,omitempty"`
}

// TaxPeriod2 type definition
type TaxPeriod2 struct {
	Yr     ISODate              `xml:"Yr,omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"Tp,omitempty"`
	FrToDt *DatePeriod2         `xml:"FrToDt,omitempty"`
}

// OtherContact1 type definition
type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"ChanlTp"`
	Id      Max128Text `xml:"Id,omitempty"`
}

// LocalInstrument2Choice type definition
type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                    `xml:"Prtry,omitempty"`
}

// AmountOrRate1Choice type definition
type AmountOrRate1Choice struct {
	Amt  *ActiveCurrencyAndAmount `xml:"Amt,omitempty"`
	Rate xsdt.Decimal             `xml:"Rate,omitempty"`
}

// Document12 type definition
type Document12 struct {
	Tp        DocumentType1Choice    `xml:"Tp"`
	Id        Max35Text              `xml:"Id"`
	IsseDt    DateAndDateTime2Choice `xml:"IsseDt"`
	Nm        Max140Text             `xml:"Nm,omitempty"`
	LangCd    xsdt.String            `xml:"LangCd,omitempty"`
	Frmt      DocumentFormat1Choice  `xml:"Frmt"`
	FileNm    Max140Text             `xml:"FileNm,omitempty"`
	DgtlSgntr *PartyAndSignature3    `xml:"DgtlSgntr,omitempty"`
	Nclsr     Max10MbBinary          `xml:"Nclsr"`
}

// BranchAndFinancialInstitutionIdentification6 type definition
type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"FinInstnId"`
	BrnchId    *BranchData3                         `xml:"BrnchId,omitempty"`
}

// FinancialIdentificationSchemeName1Choice type definition
type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                                       `xml:"Prtry,omitempty"`
}

// ProxyAccountType1Choice type definition
type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                     `xml:"Prtry,omitempty"`
}

// MandateRelatedInformation14 type definition
type MandateRelatedInformation14 struct {
	MndtId        Max35Text                      `xml:"MndtId,omitempty"`
	DtOfSgntr     ISODate                        `xml:"DtOfSgntr,omitempty"`
	AmdmntInd     xsdt.Boolean                   `xml:"AmdmntInd,omitempty"`
	AmdmntInfDtls *AmendmentInformationDetails13 `xml:"AmdmntInfDtls,omitempty"`
	ElctrncSgntr  Max1025Text                    `xml:"ElctrncSgntr,omitempty"`
	FrstColltnDt  ISODate                        `xml:"FrstColltnDt,omitempty"`
	FnlColltnDt   ISODate                        `xml:"FnlColltnDt,omitempty"`
	Frqcy         *Frequency36Choice             `xml:"Frqcy,omitempty"`
	Rsn           *MandateSetupReason1Choice     `xml:"Rsn,omitempty"`
	TrckgDays     Exact2NumericText              `xml:"TrckgDays,omitempty"`
}

// TaxAmountType1Choice type definition
type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                  `xml:"Prtry,omitempty"`
}

// CashAccountType2Choice type definition
type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                    `xml:"Prtry,omitempty"`
}

// ActiveCurrencyAndAmount type definition
type ActiveCurrencyAndAmount struct {
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
	Value xsdt.Decimal       `xml:",chardata"`
}

// PaymentCondition1 type definition
type PaymentCondition1 struct {
	AmtModAllwd   xsdt.Boolean         `xml:"AmtModAllwd"`
	EarlyPmtAllwd xsdt.Boolean         `xml:"EarlyPmtAllwd"`
	DelyPnlty     Max140Text           `xml:"DelyPnlty,omitempty"`
	ImdtPmtRbt    *AmountOrRate1Choice `xml:"ImdtPmtRbt,omitempty"`
	GrntedPmtReqd xsdt.Boolean         `xml:"GrntedPmtReqd"`
}

// AccountSchemeName1Choice type definition
type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                          `xml:"Prtry,omitempty"`
}

// Garnishment3 type definition
type Garnishment3 struct {
	Tp                GarnishmentType1                   `xml:"Tp"`
	Grnshee           *PartyIdentification135            `xml:"Grnshee,omitempty"`
	GrnshmtAdmstr     *PartyIdentification135            `xml:"GrnshmtAdmstr,omitempty"`
	RefNb             Max140Text                         `xml:"RefNb,omitempty"`
	Dt                ISODate                            `xml:"Dt,omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
	FmlyMdclInsrncInd xsdt.Boolean                       `xml:"FmlyMdclInsrncInd,omitempty"`
	MplyeeTermntnInd  xsdt.Boolean                       `xml:"MplyeeTermntnInd,omitempty"`
}

// Contact4 type definition
type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"Dept,omitempty"`
	Othr      []OtherContact1             `xml:"Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"PrefrdMtd,omitempty"`
}

// ProxyAccountIdentification1 type definition
type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"Tp,omitempty"`
	Id Max2048Text              `xml:"Id"`
}

// CategoryPurpose1Choice type definition
type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                    `xml:"Prtry,omitempty"`
}

// FrequencyAndMoment1 type definition
type FrequencyAndMoment1 struct {
	Tp     Frequency6Code    `xml:"Tp"`
	PtInTm Exact2NumericText `xml:"PtInTm"`
}

// DocumentAdjustment1 type definition
type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"CdtDbtInd,omitempty"`
	Rsn       Max4Text                          `xml:"Rsn,omitempty"`
	AddtlInf  Max140Text                        `xml:"AddtlInf,omitempty"`
}

// ClearingSystemIdentification3Choice type definition
type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                       `xml:"Prtry,omitempty"`
}

// ReferredDocumentType3Choice type definition
type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"Cd,omitempty"`
	Prtry Max35Text         `xml:"Prtry,omitempty"`
}

// Party40Choice type definition
type Party40Choice struct {
	Pty *PartyIdentification135                       `xml:"Pty,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"Agt,omitempty"`
}

// MandateSetupReason1Choice type definition
type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"Cd,omitempty"`
	Prtry Max70Text                       `xml:"Prtry,omitempty"`
}

// AddressType3Choice type definition
type AddressType3Choice struct {
	Cd    AddressType2Code         `xml:"Cd,omitempty"`
	Prtry *GenericIdentification30 `xml:"Prtry,omitempty"`
}

// StructuredRemittanceInformation16 type definition
type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty"`
	RfrdDocAmt  *RemittanceAmount2             `xml:"RfrdDocAmt,omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty"`
	Invcr       *PartyIdentification135        `xml:"Invcr,omitempty"`
	Invcee      *PartyIdentification135        `xml:"Invcee,omitempty"`
	TaxRmt      *TaxInformation7               `xml:"TaxRmt,omitempty"`
	GrnshmtRmt  *Garnishment3                  `xml:"GrnshmtRmt,omitempty"`
	AddtlRmtInf []Max140Text                   `xml:"AddtlRmtInf,omitempty"`
}

// ClearingSystemMemberIdentification2 type definition
type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty"`
	MmbId    Max35Text                            `xml:"MmbId"`
}

// DocumentLineType1 type definition
type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"CdOrPrtry"`
	Issr      Max35Text               `xml:"Issr,omitempty"`
}

// TaxRecord2 type definition
type TaxRecord2 struct {
	Tp       Max35Text   `xml:"Tp,omitempty"`
	Ctgy     Max35Text   `xml:"Ctgy,omitempty"`
	CtgyDtls Max35Text   `xml:"CtgyDtls,omitempty"`
	DbtrSts  Max35Text   `xml:"DbtrSts,omitempty"`
	CertId   Max35Text   `xml:"CertId,omitempty"`
	FrmsCd   Max35Text   `xml:"FrmsCd,omitempty"`
	Prd      *TaxPeriod2 `xml:"Prd,omitempty"`
	TaxAmt   *TaxAmount2 `xml:"TaxAmt,omitempty"`
	AddtlInf Max140Text  `xml:"AddtlInf,omitempty"`
}

// DocumentType1Choice type definition
type DocumentType1Choice struct {
	Cd    ExternalDocumentType1Code `xml:"Cd,omitempty"`
	Prtry *GenericIdentification1   `xml:"Prtry,omitempty"`
}

// PaymentTypeInformation26 type definition
type PaymentTypeInformation26 struct {
	InstrPrty Priority2Code           `xml:"InstrPrty,omitempty"`
	SvcLvl    []ServiceLevel8Choice   `xml:"SvcLvl,omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty"`
}

// PersonIdentificationSchemeName1Choice type definition
type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                         `xml:"Prtry,omitempty"`
}

// GenericFinancialIdentification1 type definition
type GenericFinancialIdentification1 struct {
	Id      Max35Text                                 `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    Max35Text                                 `xml:"Issr,omitempty"`
}

// OriginalTransactionReference28 type definition
type OriginalTransactionReference28 struct {
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty"`
	Amt            *AmountType4Choice                            `xml:"Amt,omitempty"`
	IntrBkSttlmDt  ISODate                                       `xml:"IntrBkSttlmDt,omitempty"`
	ReqdColltnDt   ISODate                                       `xml:"ReqdColltnDt,omitempty"`
	ReqdExctnDt    *DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty"`
	CdtrSchmeId    *PartyIdentification135                       `xml:"CdtrSchmeId,omitempty"`
	SttlmInf       *SettlementInstruction7                       `xml:"SttlmInf,omitempty"`
	PmtTpInf       *PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty"`
	PmtMtd         PaymentMethod4Code                            `xml:"PmtMtd,omitempty"`
	MndtRltdInf    *MandateRelatedInformation14                  `xml:"MndtRltdInf,omitempty"`
	RmtInf         *RemittanceInformation16                      `xml:"RmtInf,omitempty"`
	UltmtDbtr      *Party40Choice                                `xml:"UltmtDbtr,omitempty"`
	Dbtr           *Party40Choice                                `xml:"Dbtr,omitempty"`
	DbtrAcct       *CashAccount38                                `xml:"DbtrAcct,omitempty"`
	DbtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty"`
	DbtrAgtAcct    *CashAccount38                                `xml:"DbtrAgtAcct,omitempty"`
	CdtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty"`
	CdtrAgtAcct    *CashAccount38                                `xml:"CdtrAgtAcct,omitempty"`
	Cdtr           *Party40Choice                                `xml:"Cdtr,omitempty"`
	CdtrAcct       *CashAccount38                                `xml:"CdtrAcct,omitempty"`
	UltmtCdtr      *Party40Choice                                `xml:"UltmtCdtr,omitempty"`
	Purp           *Purpose2Choice                               `xml:"Purp,omitempty"`
}

// ReferredDocumentType4 type definition
type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"CdOrPrtry"`
	Issr      Max35Text                   `xml:"Issr,omitempty"`
}

// ReferredDocumentInformation7 type definition
type ReferredDocumentInformation7 struct {
	Tp       *ReferredDocumentType4     `xml:"Tp,omitempty"`
	Nb       Max35Text                  `xml:"Nb,omitempty"`
	RltdDt   ISODate                    `xml:"RltdDt,omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"LineDtls,omitempty"`
}

// Purpose2Choice type definition
type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"Cd,omitempty"`
	Prtry Max35Text            `xml:"Prtry,omitempty"`
}

// Party38Choice type definition
type Party38Choice struct {
	OrgId  *OrganisationIdentification29 `xml:"OrgId,omitempty"`
	PrvtId *PersonIdentification13       `xml:"PrvtId,omitempty"`
}

// TaxAmountAndType1 type definition
type TaxAmountAndType1 struct {
	Tp  *TaxAmountType1Choice             `xml:"Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

// TaxParty1 type definition
type TaxParty1 struct {
	TaxId  Max35Text `xml:"TaxId,omitempty"`
	RegnId Max35Text `xml:"RegnId,omitempty"`
	TaxTp  Max35Text `xml:"TaxTp,omitempty"`
}

// PostalAddress24 type definition
type PostalAddress24 struct {
	AdrTp       *AddressType3Choice `xml:"AdrTp,omitempty"`
	Dept        Max70Text           `xml:"Dept,omitempty"`
	SubDept     Max70Text           `xml:"SubDept,omitempty"`
	StrtNm      Max70Text           `xml:"StrtNm,omitempty"`
	BldgNb      Max16Text           `xml:"BldgNb,omitempty"`
	BldgNm      Max35Text           `xml:"BldgNm,omitempty"`
	Flr         Max70Text           `xml:"Flr,omitempty"`
	PstBx       Max16Text           `xml:"PstBx,omitempty"`
	Room        Max70Text           `xml:"Room,omitempty"`
	PstCd       Max16Text           `xml:"PstCd,omitempty"`
	TwnNm       Max35Text           `xml:"TwnNm,omitempty"`
	TwnLctnNm   Max35Text           `xml:"TwnLctnNm,omitempty"`
	DstrctNm    Max35Text           `xml:"DstrctNm,omitempty"`
	CtrySubDvsn Max35Text           `xml:"CtrySubDvsn,omitempty"`
	Ctry        CountryCode         `xml:"Ctry,omitempty"`
	AdrLine     []Max70Text         `xml:"AdrLine,omitempty"`
}

// CreditorReferenceInformation2 type definition
type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `xml:"Tp,omitempty"`
	Ref Max35Text               `xml:"Ref,omitempty"`
}

// DocumentFormat1Choice type definition
type DocumentFormat1Choice struct {
	Cd    ExternalDocumentFormat1Code `xml:"Cd,omitempty"`
	Prtry *GenericIdentification1     `xml:"Prtry,omitempty"`
}

// RemittanceAmount3 type definition
type RemittanceAmount3 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1           `xml:"DscntApldAmt,omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1                `xml:"TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1              `xml:"AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
}

// OriginalGroupInformation29 type definition
type OriginalGroupInformation29 struct {
	OrgnlMsgId   Max35Text   `xml:"OrgnlMsgId"`
	OrgnlMsgNmId Max35Text   `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm ISODateTime `xml:"OrgnlCreDtTm,omitempty"`
}

// DatePeriod2 type definition
type DatePeriod2 struct {
	FrDt ISODate `xml:"FrDt"`
	ToDt ISODate `xml:"ToDt"`
}

// GenericIdentification1 type definition
type GenericIdentification1 struct {
	Id      Max35Text `xml:"Id"`
	SchmeNm Max35Text `xml:"SchmeNm,omitempty"`
	Issr    Max35Text `xml:"Issr,omitempty"`
}

// GenericOrganisationIdentification1 type definition
type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                    `xml:"Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    Max35Text                                    `xml:"Issr,omitempty"`
}

// BranchData3 type definition
type BranchData3 struct {
	Id      Max35Text        `xml:"Id,omitempty"`
	LEI     LEIIdentifier    `xml:"LEI,omitempty"`
	Nm      Max140Text       `xml:"Nm,omitempty"`
	PstlAdr *PostalAddress24 `xml:"PstlAdr,omitempty"`
}

// AccountIdentification4Choice type definition
type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier             `xml:"IBAN,omitempty"`
	Othr *GenericAccountIdentification1 `xml:"Othr,omitempty"`
}

// DocumentLineInformation1 type definition
type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"Id"`
	Desc Max2048Text                   `xml:"Desc,omitempty"`
	Amt  *RemittanceAmount3            `xml:"Amt,omitempty"`
}

// DocumentLineIdentification1 type definition
type DocumentLineIdentification1 struct {
	Tp     *DocumentLineType1 `xml:"Tp,omitempty"`
	Nb     Max35Text          `xml:"Nb,omitempty"`
	RltdDt ISODate            `xml:"RltdDt,omitempty"`
}

// ServiceLevel8Choice type definition
type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"Cd,omitempty"`
	Prtry Max35Text                 `xml:"Prtry,omitempty"`
}

// SkipPayload type definition
type SkipPayload struct {
	Item xsdt.AnyType `xml:",any"`
}
