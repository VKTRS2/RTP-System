// Package camt_055_001_08
// Do not Edit. This stuff it's been automatically generated.
package camt_055_001_08

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
)

// CancellationReason33Choice type definition
type CancellationReason33Choice struct {
	Cd    common.ExternalCancellationReason1Code `xml:"Cd,omitempty"`
	Prtry common.Max35Text                       `xml:"Prtry,omitempty"`
}

// OriginalGroupHeader15 type definition
type OriginalGroupHeader15 struct {
	GrpCxlId     common.Max35Text             `xml:"GrpCxlId,omitempty"`
	Case         *common.Case5                `xml:"Case,omitempty"`
	OrgnlMsgId   common.Max35Text             `xml:"OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text             `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm common.ISODateTime           `xml:"OrgnlCreDtTm,omitempty"`
	NbOfTxs      common.Max15NumericText      `xml:"NbOfTxs,omitempty"`
	CtrlSum      xsdt.Decimal                 `xml:"CtrlSum,omitempty"`
	GrpCxl       xsdt.Boolean                 `xml:"GrpCxl,omitempty"`
	CxlRsnInf    []PaymentCancellationReason5 `xml:"CxlRsnInf,omitempty"`
}

// ControlData1 type definition
type ControlData1 struct {
	NbOfTxs common.Max15NumericText `xml:"NbOfTxs"`
	CtrlSum xsdt.Decimal            `xml:"CtrlSum,omitempty"`
}

// CustomerPaymentCancellationRequestV08 type definition
type CustomerPaymentCancellationRequestV08 struct {
	Assgnmt     common.CaseAssignment5      `xml:"Assgnmt"`
	Case        *common.Case5               `xml:"Case,omitempty"`
	CtrlData    *ControlData1               `xml:"CtrlData,omitempty"`
	Undrlyg     []UnderlyingTransaction24   `xml:"Undrlyg"`
	SplmtryData []common.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

// OriginalPaymentInstruction34 type definition
type OriginalPaymentInstruction34 struct {
	PmtCxlId      common.Max35Text                   `xml:"PmtCxlId,omitempty"`
	Case          *common.Case5                      `xml:"Case,omitempty"`
	OrgnlPmtInfId common.Max35Text                   `xml:"OrgnlPmtInfId"`
	OrgnlGrpInf   *common.OriginalGroupInformation29 `xml:"OrgnlGrpInf,omitempty"`
	NbOfTxs       common.Max15NumericText            `xml:"NbOfTxs,omitempty"`
	CtrlSum       xsdt.Decimal                       `xml:"CtrlSum,omitempty"`
	PmtInfCxl     xsdt.Boolean                       `xml:"PmtInfCxl,omitempty"`
	CxlRsnInf     []PaymentCancellationReason5       `xml:"CxlRsnInf,omitempty"`
	TxInf         []PaymentTransaction109            `xml:"TxInf,omitempty"`
}

// PaymentCancellationReason5 type definition
type PaymentCancellationReason5 struct {
	Orgtr    *common.PartyIdentification135 `xml:"Orgtr,omitempty"`
	Rsn      *CancellationReason33Choice    `xml:"Rsn,omitempty"`
	AddtlInf []common.Max105Text            `xml:"AddtlInf,omitempty"`
}

// PaymentTransaction109 type definition
type PaymentTransaction109 struct {
	CxlId             common.Max35Text                          `xml:"CxlId,omitempty"`
	Case              *common.Case5                             `xml:"Case,omitempty"`
	OrgnlInstrId      common.Max35Text                          `xml:"OrgnlInstrId,omitempty"`
	OrgnlEndToEndId   common.Max35Text                          `xml:"OrgnlEndToEndId,omitempty"`
	OrgnlUETR         common.UUIDv4Identifier                   `xml:"OrgnlUETR,omitempty"`
	OrgnlInstdAmt     *common.ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`
	OrgnlReqdExctnDt  *common.DateAndDateTime2Choice            `xml:"OrgnlReqdExctnDt,omitempty"`
	OrgnlReqdColltnDt common.ISODate                            `xml:"OrgnlReqdColltnDt,omitempty"`
	CxlRsnInf         []PaymentCancellationReason5              `xml:"CxlRsnInf,omitempty"`
	OrgnlTxRef        *common.OriginalTransactionReference28    `xml:"OrgnlTxRef,omitempty"`
	SplmtryData       []common.SupplementaryData1               `xml:"SplmtryData,omitempty"`
}

// UnderlyingTransaction24 type definition
type UnderlyingTransaction24 struct {
	OrgnlGrpInfAndCxl *OriginalGroupHeader15         `xml:"OrgnlGrpInfAndCxl,omitempty"`
	OrgnlPmtInfAndCxl []OriginalPaymentInstruction34 `xml:"OrgnlPmtInfAndCxl,omitempty"`
}
