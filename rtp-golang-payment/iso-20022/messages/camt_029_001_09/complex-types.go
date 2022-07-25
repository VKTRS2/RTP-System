// Package camt_029_001_09
// Do not Edit. This stuff it's been automatically generated.
package camt_029_001_09

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
)

// PaymentTransaction103 type definition
type PaymentTransaction103 struct {
	CxlStsId          common.Max35Text                          `xml:"CxlStsId,omitempty"`
	RslvdCase         *common.Case5                             `xml:"RslvdCase,omitempty"`
	OrgnlInstrId      common.Max35Text                          `xml:"OrgnlInstrId,omitempty"`
	OrgnlEndToEndId   common.Max35Text                          `xml:"OrgnlEndToEndId,omitempty"`
	UETR              common.UUIDv4Identifier                   `xml:"UETR,omitempty"`
	TxCxlSts          common.CancellationIndividualStatus1Code  `xml:"TxCxlSts,omitempty"`
	CxlStsRsnInf      []CancellationStatusReason4               `xml:"CxlStsRsnInf,omitempty"`
	OrgnlInstdAmt     *common.ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`
	OrgnlReqdExctnDt  *common.DateAndDateTime2Choice            `xml:"OrgnlReqdExctnDt,omitempty"`
	OrgnlReqdColltnDt common.ISODate                            `xml:"OrgnlReqdColltnDt,omitempty"`
	OrgnlTxRef        *common.OriginalTransactionReference28    `xml:"OrgnlTxRef,omitempty"`
}

// ResolutionOfInvestigationV09 type definition
type ResolutionOfInvestigationV09 struct {
	Assgnmt       common.CaseAssignment5        `xml:"Assgnmt"`
	RslvdCase     *common.Case5                 `xml:"RslvdCase,omitempty"`
	Sts           InvestigationStatus5Choice    `xml:"Sts"`
	CxlDtls       []UnderlyingTransaction22     `xml:"CxlDtls,omitempty"`
	ModDtls       *PaymentTransaction107        `xml:"ModDtls,omitempty"`
	ClmNonRctDtls *ClaimNonReceipt2Choice       `xml:"ClmNonRctDtls,omitempty"`
	StmtDtls      *StatementResolutionEntry4    `xml:"StmtDtls,omitempty"`
	CrrctnTx      *CorrectiveTransaction4Choice `xml:"CrrctnTx,omitempty"`
	RsltnRltdInf  *ResolutionData1              `xml:"RsltnRltdInf,omitempty"`
	SplmtryData   []common.SupplementaryData1   `xml:"SplmtryData,omitempty"`
}

// Compensation2 type definition
type Compensation2 struct {
	Amt     common.ActiveCurrencyAndAmount                      `xml:"Amt"`
	DbtrAgt common.BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt"`
	CdtrAgt common.BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt"`
	Rsn     CompensationReason1Choice                           `xml:"Rsn"`
}

// NumberOfTransactionsPerStatus1 type definition
type NumberOfTransactionsPerStatus1 struct {
	DtldNbOfTxs common.Max15NumericText                 `xml:"DtldNbOfTxs"`
	DtldSts     common.TransactionIndividualStatus1Code `xml:"DtldSts"`
	DtldCtrlSum xsdt.Decimal                            `xml:"DtldCtrlSum,omitempty"`
}

// ClaimNonReceipt2Choice type definition
type ClaimNonReceipt2Choice struct {
	Accptd *ClaimNonReceipt2                   `xml:"Accptd,omitempty"`
	Rjctd  *ClaimNonReceiptRejectReason1Choice `xml:"Rjctd,omitempty"`
}

// CorrectiveGroupInformation1 type definition
type CorrectiveGroupInformation1 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	MsgNmId common.Max35Text   `xml:"MsgNmId"`
	CreDtTm common.ISODateTime `xml:"CreDtTm,omitempty"`
}

// CompensationReason1Choice type definition
type CompensationReason1Choice struct {
	Cd    common.ExternalPaymentCompensationReason1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                              `xml:"Prtry,omitempty"`
}

// CancellationStatusReason4 type definition
type CancellationStatusReason4 struct {
	Orgtr    *common.PartyIdentification135   `xml:"Orgtr,omitempty"`
	Rsn      *CancellationStatusReason3Choice `xml:"Rsn,omitempty"`
	AddtlInf []common.Max105Text              `xml:"AddtlInf,omitempty"`
}

// ClaimNonReceipt2 type definition
type ClaimNonReceipt2 struct {
	DtPrcd      common.ISODate                                       `xml:"DtPrcd"`
	OrgnlNxtAgt *common.BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlNxtAgt,omitempty"`
}

// GenericIdentification3 type definition
type GenericIdentification3 struct {
	Id   common.Max35Text `xml:"Id"`
	Issr common.Max35Text `xml:"Issr,omitempty"`
}

// CorrectivePaymentInitiation4 type definition
type CorrectivePaymentInitiation4 struct {
	GrpHdr       *CorrectiveGroupInformation1             `xml:"GrpHdr,omitempty"`
	PmtInfId     common.Max35Text                         `xml:"PmtInfId,omitempty"`
	InstrId      common.Max35Text                         `xml:"InstrId,omitempty"`
	EndToEndId   common.Max35Text                         `xml:"EndToEndId,omitempty"`
	UETR         common.UUIDv4Identifier                  `xml:"UETR,omitempty"`
	InstdAmt     common.ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`
	ReqdExctnDt  *common.DateAndDateTime2Choice           `xml:"ReqdExctnDt,omitempty"`
	ReqdColltnDt common.ISODate                           `xml:"ReqdColltnDt,omitempty"`
}

// CorrectiveInterbankTransaction2 type definition
type CorrectiveInterbankTransaction2 struct {
	GrpHdr         *CorrectiveGroupInformation1             `xml:"GrpHdr,omitempty"`
	InstrId        common.Max35Text                         `xml:"InstrId,omitempty"`
	EndToEndId     common.Max35Text                         `xml:"EndToEndId,omitempty"`
	TxId           common.Max35Text                         `xml:"TxId,omitempty"`
	UETR           common.UUIDv4Identifier                  `xml:"UETR,omitempty"`
	IntrBkSttlmAmt common.ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt"`
	IntrBkSttlmDt  common.ISODate                           `xml:"IntrBkSttlmDt"`
}

// ChargesRecord3 type definition
type ChargesRecord3 struct {
	Amt         common.ActiveOrHistoricCurrencyAndAmount             `xml:"Amt"`
	CdtDbtInd   common.CreditDebitCode                               `xml:"CdtDbtInd,omitempty"`
	ChrgInclInd xsdt.Boolean                                         `xml:"ChrgInclInd,omitempty"`
	Tp          *ChargeType3Choice                                   `xml:"Tp,omitempty"`
	Rate        xsdt.Decimal                                         `xml:"Rate,omitempty"`
	Br          common.ChargeBearerType1Code                         `xml:"Br,omitempty"`
	Agt         *common.BranchAndFinancialInstitutionIdentification6 `xml:"Agt,omitempty"`
	Tax         *TaxCharges2                                         `xml:"Tax,omitempty"`
}

// ResolutionData1 type definition
type ResolutionData1 struct {
	EndToEndId     common.Max35Text                          `xml:"EndToEndId,omitempty"`
	TxId           common.Max35Text                          `xml:"TxId,omitempty"`
	UETR           common.UUIDv4Identifier                   `xml:"UETR,omitempty"`
	IntrBkSttlmAmt *common.ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty"`
	IntrBkSttlmDt  common.ISODate                            `xml:"IntrBkSttlmDt,omitempty"`
	ClrChanl       common.ClearingChannel2Code               `xml:"ClrChanl,omitempty"`
	Compstn        *Compensation2                            `xml:"Compstn,omitempty"`
	Chrgs          []common.Charges7                         `xml:"Chrgs,omitempty"`
}

// OriginalGroupHeader14 type definition
type OriginalGroupHeader14 struct {
	OrgnlGrpCxlId    common.Max35Text                    `xml:"OrgnlGrpCxlId,omitempty"`
	RslvdCase        *common.Case5                       `xml:"RslvdCase,omitempty"`
	OrgnlMsgId       common.Max35Text                    `xml:"OrgnlMsgId"`
	OrgnlMsgNmId     common.Max35Text                    `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm     common.ISODateTime                  `xml:"OrgnlCreDtTm,omitempty"`
	OrgnlNbOfTxs     common.Max15NumericText             `xml:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum     xsdt.Decimal                        `xml:"OrgnlCtrlSum,omitempty"`
	GrpCxlSts        common.GroupCancellationStatus1Code `xml:"GrpCxlSts,omitempty"`
	CxlStsRsnInf     []CancellationStatusReason4         `xml:"CxlStsRsnInf,omitempty"`
	NbOfTxsPerCxlSts []NumberOfTransactionsPerStatus1    `xml:"NbOfTxsPerCxlSts,omitempty"`
}

// Charges6 type definition
type Charges6 struct {
	TtlChrgsAndTaxAmt *common.ActiveOrHistoricCurrencyAndAmount `xml:"TtlChrgsAndTaxAmt,omitempty"`
	Rcrd              []ChargesRecord3                          `xml:"Rcrd,omitempty"`
}

// StatementResolutionEntry4 type definition
type StatementResolutionEntry4 struct {
	OrgnlGrpInf *common.OriginalGroupInformation29        `xml:"OrgnlGrpInf,omitempty"`
	OrgnlStmtId common.Max35Text                          `xml:"OrgnlStmtId,omitempty"`
	UETR        common.UUIDv4Identifier                   `xml:"UETR,omitempty"`
	AcctSvcrRef common.Max35Text                          `xml:"AcctSvcrRef,omitempty"`
	CrrctdAmt   *common.ActiveOrHistoricCurrencyAndAmount `xml:"CrrctdAmt,omitempty"`
	Chrgs       []Charges6                                `xml:"Chrgs,omitempty"`
	Purp        *common.Purpose2Choice                    `xml:"Purp,omitempty"`
}

// ClaimNonReceiptRejectReason1Choice type definition
type ClaimNonReceiptRejectReason1Choice struct {
	Cd    common.ExternalClaimNonReceiptRejection1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                             `xml:"Prtry,omitempty"`
}

// ModificationStatusReason2 type definition
type ModificationStatusReason2 struct {
	Orgtr    *common.PartyIdentification135   `xml:"Orgtr,omitempty"`
	Rsn      *ModificationStatusReason1Choice `xml:"Rsn,omitempty"`
	AddtlInf []common.Max105Text              `xml:"AddtlInf,omitempty"`
}

// PaymentTransaction107 type definition
type PaymentTransaction107 struct {
	ModStsId            common.Max35Text                          `xml:"ModStsId,omitempty"`
	RslvdCase           *common.Case5                             `xml:"RslvdCase,omitempty"`
	OrgnlGrpInf         common.OriginalGroupInformation29         `xml:"OrgnlGrpInf"`
	OrgnlPmtInfId       common.Max35Text                          `xml:"OrgnlPmtInfId,omitempty"`
	OrgnlInstrId        common.Max35Text                          `xml:"OrgnlInstrId,omitempty"`
	OrgnlEndToEndId     common.Max35Text                          `xml:"OrgnlEndToEndId,omitempty"`
	OrgnlTxId           common.Max35Text                          `xml:"OrgnlTxId,omitempty"`
	OrgnlClrSysRef      common.Max35Text                          `xml:"OrgnlClrSysRef,omitempty"`
	OrgnlUETR           common.UUIDv4Identifier                   `xml:"OrgnlUETR,omitempty"`
	ModStsRsnInf        []ModificationStatusReason2               `xml:"ModStsRsnInf,omitempty"`
	RsltnRltdInf        *ResolutionData1                          `xml:"RsltnRltdInf,omitempty"`
	OrgnlIntrBkSttlmAmt *common.ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`
	OrgnlIntrBkSttlmDt  common.ISODate                            `xml:"OrgnlIntrBkSttlmDt,omitempty"`
	Assgnr              *common.Party40Choice                     `xml:"Assgnr,omitempty"`
	Assgne              *common.Party40Choice                     `xml:"Assgne,omitempty"`
	OrgnlTxRef          *common.OriginalTransactionReference28    `xml:"OrgnlTxRef,omitempty"`
}

// ModificationStatusReason1Choice type definition
type ModificationStatusReason1Choice struct {
	Cd    common.ExternalPaymentModificationRejection1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                                 `xml:"Prtry,omitempty"`
}

// OriginalPaymentInstruction30 type definition
type OriginalPaymentInstruction30 struct {
	OrgnlPmtInfCxlId common.Max35Text                    `xml:"OrgnlPmtInfCxlId,omitempty"`
	RslvdCase        *common.Case5                       `xml:"RslvdCase,omitempty"`
	OrgnlPmtInfId    common.Max35Text                    `xml:"OrgnlPmtInfId"`
	OrgnlGrpInf      *common.OriginalGroupInformation29  `xml:"OrgnlGrpInf,omitempty"`
	OrgnlNbOfTxs     common.Max15NumericText             `xml:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum     xsdt.Decimal                        `xml:"OrgnlCtrlSum,omitempty"`
	PmtInfCxlSts     common.GroupCancellationStatus1Code `xml:"PmtInfCxlSts,omitempty"`
	CxlStsRsnInf     []CancellationStatusReason4         `xml:"CxlStsRsnInf,omitempty"`
	NbOfTxsPerCxlSts []NumberOfCancellationsPerStatus1   `xml:"NbOfTxsPerCxlSts,omitempty"`
	TxInfAndSts      []PaymentTransaction103             `xml:"TxInfAndSts,omitempty"`
}

// ChargeType3Choice type definition
type ChargeType3Choice struct {
	Cd    common.ExternalChargeType1Code `xml:"Cd,omitempty"`
	Prtry *GenericIdentification3        `xml:"Prtry,omitempty"`
}

// CancellationStatusReason3Choice type definition
type CancellationStatusReason3Choice struct {
	Cd    common.ExternalPaymentCancellationRejection1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                                 `xml:"Prtry,omitempty"`
}

// NumberOfCancellationsPerStatus1 type definition
type NumberOfCancellationsPerStatus1 struct {
	DtldNbOfTxs common.Max15NumericText                  `xml:"DtldNbOfTxs"`
	DtldSts     common.CancellationIndividualStatus1Code `xml:"DtldSts"`
	DtldCtrlSum xsdt.Decimal                             `xml:"DtldCtrlSum,omitempty"`
}

// CorrectiveTransaction4Choice type definition
type CorrectiveTransaction4Choice struct {
	Initn  *CorrectivePaymentInitiation4    `xml:"Initn,omitempty"`
	IntrBk *CorrectiveInterbankTransaction2 `xml:"IntrBk,omitempty"`
}

// UnderlyingTransaction22 type definition
type UnderlyingTransaction22 struct {
	OrgnlGrpInfAndSts *OriginalGroupHeader14         `xml:"OrgnlGrpInfAndSts,omitempty"`
	OrgnlPmtInfAndSts []OriginalPaymentInstruction30 `xml:"OrgnlPmtInfAndSts,omitempty"`
	TxInfAndSts       []PaymentTransaction102        `xml:"TxInfAndSts,omitempty"`
}

// InvestigationStatus5Choice type definition
type InvestigationStatus5Choice struct {
	Conf           common.ExternalInvestigationExecutionConfirmation1Code `xml:"Conf,omitempty"`
	RjctdMod       []ModificationStatusReason1Choice                      `xml:"RjctdMod,omitempty"`
	DplctOf        *common.Case5                                          `xml:"DplctOf,omitempty"`
	AssgnmtCxlConf xsdt.Boolean                                           `xml:"AssgnmtCxlConf,omitempty"`
}

// PaymentTransaction102 type definition
type PaymentTransaction102 struct {
	CxlStsId            common.Max35Text                          `xml:"CxlStsId,omitempty"`
	RslvdCase           *common.Case5                             `xml:"RslvdCase,omitempty"`
	OrgnlGrpInf         *common.OriginalGroupInformation29        `xml:"OrgnlGrpInf,omitempty"`
	OrgnlInstrId        common.Max35Text                          `xml:"OrgnlInstrId,omitempty"`
	OrgnlEndToEndId     common.Max35Text                          `xml:"OrgnlEndToEndId,omitempty"`
	OrgnlTxId           common.Max35Text                          `xml:"OrgnlTxId,omitempty"`
	OrgnlClrSysRef      common.Max35Text                          `xml:"OrgnlClrSysRef,omitempty"`
	OrgnlUETR           common.UUIDv4Identifier                   `xml:"OrgnlUETR,omitempty"`
	TxCxlSts            common.CancellationIndividualStatus1Code  `xml:"TxCxlSts,omitempty"`
	CxlStsRsnInf        []CancellationStatusReason4               `xml:"CxlStsRsnInf,omitempty"`
	RsltnRltdInf        *ResolutionData1                          `xml:"RsltnRltdInf,omitempty"`
	OrgnlIntrBkSttlmAmt *common.ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`
	OrgnlIntrBkSttlmDt  common.ISODate                            `xml:"OrgnlIntrBkSttlmDt,omitempty"`
	Assgnr              *common.Party40Choice                     `xml:"Assgnr,omitempty"`
	Assgne              *common.Party40Choice                     `xml:"Assgne,omitempty"`
	OrgnlTxRef          *common.OriginalTransactionReference28    `xml:"OrgnlTxRef,omitempty"`
}

// TaxCharges2 type definition
type TaxCharges2 struct {
	Id   common.Max35Text                          `xml:"Id,omitempty"`
	Rate xsdt.Decimal                              `xml:"Rate,omitempty"`
	Amt  *common.ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
}
