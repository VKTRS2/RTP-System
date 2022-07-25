// Package pain_014_001_07
// Do not Edit. This stuff it's been automatically generated.
package pain_014_001_07

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
)

// GroupHeader87 type definition
type GroupHeader87 struct {
	MsgId    common.Max35Text                                     `xml:"MsgId"`
	CreDtTm  common.ISODateTime                                   `xml:"CreDtTm"`
	InitgPty common.PartyIdentification135                        `xml:"InitgPty"`
	DbtrAgt  *common.BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty"`
	CdtrAgt  *common.BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty"`
}

// StatusReasonInformation12 type definition
type StatusReasonInformation12 struct {
	Orgtr    *common.PartyIdentification135 `xml:"Orgtr,omitempty"`
	Rsn      *StatusReason6Choice           `xml:"Rsn,omitempty"`
	AddtlInf []common.Max105Text            `xml:"AddtlInf,omitempty"`
}

// NumberOfTransactionsPerStatus5 type definition
type NumberOfTransactionsPerStatus5 struct {
	DtldNbOfTxs common.Max15NumericText                      `xml:"DtldNbOfTxs"`
	DtldSts     common.ExternalPaymentTransactionStatus1Code `xml:"DtldSts"`
	DtldCtrlSum xsdt.Decimal                                 `xml:"DtldCtrlSum,omitempty"`
}

// OriginalPaymentInstruction31 type definition
type OriginalPaymentInstruction31 struct {
	OrgnlPmtInfId common.Max35Text                       `xml:"OrgnlPmtInfId"`
	OrgnlNbOfTxs  common.Max15NumericText                `xml:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  xsdt.Decimal                           `xml:"OrgnlCtrlSum,omitempty"`
	PmtInfSts     common.ExternalPaymentGroupStatus1Code `xml:"PmtInfSts,omitempty"`
	StsRsnInf     []StatusReasonInformation12            `xml:"StsRsnInf,omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus5       `xml:"NbOfTxsPerSts,omitempty"`
	TxInfAndSts   []PaymentTransaction104                `xml:"TxInfAndSts,omitempty"`
}

// OriginalGroupInformation30 type definition
type OriginalGroupInformation30 struct {
	OrgnlMsgId    common.Max35Text                       `xml:"OrgnlMsgId"`
	OrgnlMsgNmId  common.Max35Text                       `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm  common.ISODateTime                     `xml:"OrgnlCreDtTm,omitempty"`
	OrgnlNbOfTxs  common.Max15NumericText                `xml:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  xsdt.Decimal                           `xml:"OrgnlCtrlSum,omitempty"`
	GrpSts        common.ExternalPaymentGroupStatus1Code `xml:"GrpSts,omitempty"`
	StsRsnInf     []StatusReasonInformation12            `xml:"StsRsnInf,omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus5       `xml:"NbOfTxsPerSts,omitempty"`
}

// StatusReason6Choice type definition
type StatusReason6Choice struct {
	Cd    common.ExternalStatusReason1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                 `xml:"Prtry,omitempty"`
}

// PaymentConditionStatus1 type definition
type PaymentConditionStatus1 struct {
	AccptdAmt *common.ActiveCurrencyAndAmount `xml:"AccptdAmt,omitempty"`
	GrntedPmt xsdt.Boolean                    `xml:"GrntedPmt"`
	EarlyPmt  xsdt.Boolean                    `xml:"EarlyPmt"`
}

// OriginalTransactionReference29 type definition
type OriginalTransactionReference29 struct {
	Amt         *common.AmountType4Choice                            `xml:"Amt,omitempty"`
	ReqdExctnDt *common.DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty"`
	XpryDt      *common.DateAndDateTime2Choice                       `xml:"XpryDt,omitempty"`
	PmtCond     *common.PaymentCondition1                            `xml:"PmtCond,omitempty"`
	PmtTpInf    *common.PaymentTypeInformation26                     `xml:"PmtTpInf,omitempty"`
	PmtMtd      common.PaymentMethod4Code                            `xml:"PmtMtd,omitempty"`
	RmtInf      *common.RemittanceInformation16                      `xml:"RmtInf,omitempty"`
	NclsdFile   []common.Document12                                  `xml:"NclsdFile,omitempty"`
	UltmtDbtr   *common.PartyIdentification135                       `xml:"UltmtDbtr,omitempty"`
	Dbtr        *common.PartyIdentification135                       `xml:"Dbtr,omitempty"`
	DbtrAcct    *common.CashAccount38                                `xml:"DbtrAcct,omitempty"`
	DbtrAgt     *common.BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty"`
	CdtrAgt     common.BranchAndFinancialInstitutionIdentification6  `xml:"CdtrAgt"`
	Cdtr        common.PartyIdentification135                        `xml:"Cdtr"`
	CdtrAcct    *common.CashAccount38                                `xml:"CdtrAcct,omitempty"`
	UltmtCdtr   *common.PartyIdentification135                       `xml:"UltmtCdtr,omitempty"`
}

// PaymentTransaction104 type definition
type PaymentTransaction104 struct {
	StsId           common.Max35Text                             `xml:"StsId,omitempty"`
	OrgnlInstrId    common.Max35Text                             `xml:"OrgnlInstrId,omitempty"`
	OrgnlEndToEndId common.Max35Text                             `xml:"OrgnlEndToEndId,omitempty"`
	OrgnlUETR       common.UUIDv4Identifier                      `xml:"OrgnlUETR,omitempty"`
	TxSts           common.ExternalPaymentTransactionStatus1Code `xml:"TxSts,omitempty"`
	StsRsnInf       []StatusReasonInformation12                  `xml:"StsRsnInf,omitempty"`
	PmtCondSts      *PaymentConditionStatus1                     `xml:"PmtCondSts,omitempty"`
	ChrgsInf        []common.Charges7                            `xml:"ChrgsInf,omitempty"`
	DbtrDcsnDtTm    common.ISODateTime                           `xml:"DbtrDcsnDtTm,omitempty"`
	AccptncDtTm     common.ISODateTime                           `xml:"AccptncDtTm,omitempty"`
	AcctSvcrRef     common.Max35Text                             `xml:"AcctSvcrRef,omitempty"`
	ClrSysRef       common.Max35Text                             `xml:"ClrSysRef,omitempty"`
	OrgnlTxRef      *OriginalTransactionReference29              `xml:"OrgnlTxRef,omitempty"`
	NclsdFile       []common.Document12                          `xml:"NclsdFile,omitempty"`
	SplmtryData     []common.SupplementaryData1                  `xml:"SplmtryData,omitempty"`
}

// CreditorPaymentActivationRequestStatusReportV07 type definition
type CreditorPaymentActivationRequestStatusReportV07 struct {
	GrpHdr            GroupHeader87                  `xml:"GrpHdr"`
	OrgnlGrpInfAndSts OriginalGroupInformation30     `xml:"OrgnlGrpInfAndSts"`
	OrgnlPmtInfAndSts []OriginalPaymentInstruction31 `xml:"OrgnlPmtInfAndSts,omitempty"`
	SplmtryData       []common.SupplementaryData1    `xml:"SplmtryData,omitempty"`
}
