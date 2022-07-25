package rtp

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/rs/zerolog/log"
)

type Status struct {
	Code          string `yaml:"cd" mapstructure:"cd" json:"cd"`
	ReasonCode    string `yaml:"rsn,omitempty" mapstructure:"rsn,omitempty" json:"rsn,omitempty"`
	Text          string `yaml:"text,omitempty" mapstructure:"text,omitempty" json:"text,omitempty"`
	RfCInProgress bool   `yaml:"rfc-in-progress,omitempty" mapstructure:"rfc-in-progress,omitempty" json:"rfc-in-progress,omitempty"`
}

func (st *Status) IsFinal() bool {
	return IsStatusFinal(st.Code)
}

func (st *Status) UpdateText(dataset string) string {
	st.Text = LookupStsCdRsnText(dataset, st.Code, st.ReasonCode)
	return st.Text
}

func (st *Status) ShowInfo() {
	log.Trace().Str("tx-sts", st.Code).Str("rsn", st.ReasonCode).Str("text", st.Text).Msg("status info")
}

const (
	DataSetDS01 = "DS01"
	DataSetDS02 = "DS02"
	DataSetDS03 = "DS03"
	DataSetDS04 = "DS04"
	DataSetDS05 = "DS05"
	DataSetDS06 = "DS06"
	DataSetDS07 = "DS07"
	DataSetDS08 = "DS08"
	DataSetDS09 = "DS09"
	DataSetDS10 = "DS10"
	DataSetDS11 = "DS11"
	DataSetDS12 = "DS12"
	DataSetDS13 = "DS13"
	DataSetDS14 = "DS14"
	DataSetDS15 = "DS15"
	DataSetDS16 = "DS16"
	DataSetDS17 = "DS17"

	// IG 2.0 Pag. 64 (DS-04a), 109 (DS-04b+), 159 (DS-05, DS-06), 203 (DS-08, DS-09 positive), 257 (DS-08, DS-09 negative), 495 (DS-17)

	StatusZero            = "RCRE" // Non ISO. Represent a non status.
	RTPRejectInvalidState = "ISTS" // Non Iso20022 Represent an invalid transition.

	StatusRejected            = "RJCT" // applies to: DS04, DS08, DS09 (negative)
	StatusTechnicallyAccepted = "ACTC" // applies to: DS05, DS06
	StatusAccepted            = "ACCP" // applies to: DS08, DS09 (positive)
	StatusAcceptedWithChange  = "ACWC" // applies to: DS08, DS09 (positive)

	StatusReasonAlreadyExpiredRTP         = "AEXR" // applies to: DS17 StatusUpdate pain.014
	StatusReasonAlreadyAcceptedRTP        = "ALAC" // applies to: DS10b, DS13, DS17 StatusUpdate pain.014
	StatusReasonAlreadyRefusedRTP         = "ARFR" // applies to: DS10b, DS13  StatusUpdate pain.014
	StatusReasonAlreadyRejectedRTP        = "ARJR" // applies to: DS10b, DS13   StatusUpdate pain.014
	StatusReasonInitialRTPNeverReceived   = "IRNR" // applies to: DS17 StatusUpdate pain.014
	StatusReasonRTPReceivedCanBeProcessed = "REPR" // applies to: DS17 StatusUpdate pain.014

	RejectReasonInvalidDebtorAccountNumber            = "AC02" // applies to: DS04
	RejectReasonNotAllowedCurrency                    = "AM03" // applies to: DS04, DS08, DS09
	RejectReasonDuplication                           = "AM05" // applies to: DS04, DS08, DS09
	RejectReasonWrongAmount                           = "AM09" // applies to: DS08, DS09
	RejectReasonAttachmentsNotSupported               = "ATNS" // applies to: DS04
	RejectReasonInvalidDebtorIdentificationCode       = "BE16" // applies to: DS04
	RejectReasonExpiryDateTooLong                     = "EDTL" // applies to: DS04
	RejectReasonExpiryDateTimeReached                 = "EDTR" // applies to: DS04
	RejectReasonInvalidFileFormat                     = "FF01" // applies to: DS04
	RejectReasonFraudulentOrigin                      = "FRAD" // applies to: DS04
	RejectReasonIncorrectExpiryDateTime               = "IEDT" // applies to: DS08, DS09
	RejectReasonNotSpecifiedReasonCustomerGenerated   = "MS02" // applies to: DS08, DS09
	RejectReasonNotSpecifiedReasonAgentGenerated      = "MS03" // applies to: DS04
	RejectReasonNonAgreedRTP                          = "NOAR" // applies to: DS08, DS09
	RejectReasonPayerOrPayerRTPSPNotReachable         = "NRCH" // applies to: DS04
	RejectReasonTypeOfPaymentInstrumentNotSupported   = "PINS" // applies to: DS04
	RejectReasonRegulatoryReason                      = "RR04" // applies to: DS04, DS10b, DS13
	RejectReasonRTPNotSupportedForDebtor              = "RTNS" // applies to: DS04
	RejectReasonRTPServiceProviderIdentifierIncorrect = "SPII" // applies to: DS04
	RejectReasonUnknownCreditor                       = "UCRD" // applies to: DS08, DS09

	// IG 2.0 Pag. 348 (DS-10b), 380 (DS-12, DS-13 positive), 411 (DS-12, DS-13 negative), 520 (DS-16, DS-17 status update)

	StatusRejectedCancellationRequest                          = "RJCR" // applies to: DS10b, DS13 (negative)
	StatusCancelledAsPerRequest                                = "CNCL" // applies to: DS12, DS13 (positive)
	RejectCancellationReasonAlreadyCancelledRTP                = "ACLR" // applies to: DS10b, DS13
	RejectCancellationReasonPaymentAlreadyTransmittedExecution = "PATE" // applies to: DS10b, DS13
	RejectCancellationReasonUnknownRTP                         = "URTP" // applies to: DS10b, DS13
	RfCStatusUpdateAlreadyRejected                             = "RCAR" // applies to: DS17
	RfCStatusUpdateNeverReceived                               = "RCNR" // applies to: DS17
	RfCStatusUpdateReceivedAndProcessed                        = "RCPR" // applies to: DS17
)

type StsRsnCodeDescription struct {
	Code        string
	Description string
}

var StsRsnCodeDescriptionRegistry = map[string]StsRsnCodeDescription{
	StatusZero: {
		Code:        StatusZero,
		Description: "Created",
	},
	RTPRejectInvalidState: {
		Code:        RTPRejectInvalidState,
		Description: "Invalid state",
	},
	StatusRejected: {
		Code:        StatusRejected,
		Description: "Rejected",
	},
	StatusTechnicallyAccepted: {
		Code:        StatusTechnicallyAccepted,
		Description: "Technically Accepted",
	},
	StatusAccepted: {
		Code:        StatusAccepted,
		Description: "Accepted",
	},
	StatusAcceptedWithChange: {
		Code:        StatusAcceptedWithChange,
		Description: "Accepted With Change",
	},
	StatusReasonAlreadyExpiredRTP: {
		Code:        StatusReasonAlreadyExpiredRTP,
		Description: "Already Expired",
	},
	StatusReasonAlreadyAcceptedRTP: {
		Code:        StatusReasonAlreadyAcceptedRTP,
		Description: "Already Accepted",
	},
	StatusReasonAlreadyRefusedRTP: {
		Code:        StatusReasonAlreadyRefusedRTP,
		Description: "Already Refused",
	},
	StatusReasonAlreadyRejectedRTP: {
		Code:        StatusReasonAlreadyRejectedRTP,
		Description: "Already Rejected",
	},
	StatusReasonInitialRTPNeverReceived: {
		Code:        StatusReasonInitialRTPNeverReceived,
		Description: "Initial RTP Never Received",
	},
	StatusReasonRTPReceivedCanBeProcessed: {
		Code:        StatusReasonRTPReceivedCanBeProcessed,
		Description: "RTP Received Can Be Processed",
	},
	RejectReasonInvalidDebtorAccountNumber: {
		Code:        RejectReasonInvalidDebtorAccountNumber,
		Description: "Invalid Debtor Account Number",
	},
	RejectReasonNotAllowedCurrency: {
		Code:        RejectReasonNotAllowedCurrency,
		Description: "Not Allowed Currency",
	},
	RejectReasonDuplication: {
		Code:        RejectReasonDuplication,
		Description: "Duplication",
	},
	RejectReasonWrongAmount: {
		Code:        RejectReasonWrongAmount,
		Description: "Wrong Amount",
	},
	RejectReasonAttachmentsNotSupported: {
		Code:        RejectReasonAttachmentsNotSupported,
		Description: "Attachments Not Supported",
	},
	RejectReasonInvalidDebtorIdentificationCode: {
		Code:        RejectReasonInvalidDebtorIdentificationCode,
		Description: "Invalid Debtor Identification Code",
	},
	RejectReasonExpiryDateTooLong: {
		Code:        RejectReasonExpiryDateTooLong,
		Description: "Expiry Date Too Long",
	},
	RejectReasonExpiryDateTimeReached: {
		Code:        RejectReasonExpiryDateTimeReached,
		Description: "Expiry Date Time Reached",
	},
	RejectReasonInvalidFileFormat: {
		Code:        RejectReasonInvalidFileFormat,
		Description: "Invalid File Format",
	},
	RejectReasonFraudulentOrigin: {
		Code:        RejectReasonFraudulentOrigin,
		Description: "Fraudulent Origin",
	},
	RejectReasonIncorrectExpiryDateTime: {
		Code:        RejectReasonIncorrectExpiryDateTime,
		Description: "Incorrect Expiry DateTime",
	},
	RejectReasonNotSpecifiedReasonCustomerGenerated: {
		Code:        RejectReasonNotSpecifiedReasonCustomerGenerated,
		Description: "Not Specified Reason Customer Generated",
	},
	RejectReasonNotSpecifiedReasonAgentGenerated: {
		Code:        RejectReasonNotSpecifiedReasonAgentGenerated,
		Description: "Reason has not been specified by agent",
	},
	RejectReasonNonAgreedRTP: {
		Code:        RejectReasonNonAgreedRTP,
		Description: "Non Agreed RTP",
	},
	RejectReasonPayerOrPayerRTPSPNotReachable: {
		Code:        RejectReasonPayerOrPayerRTPSPNotReachable,
		Description: "Payer Or Payer RTPSP Not Reachable",
	},
	RejectReasonTypeOfPaymentInstrumentNotSupported: {
		Code:        RejectReasonTypeOfPaymentInstrumentNotSupported,
		Description: "Type Of Payment Instrument Not Supported",
	},
	RejectReasonRegulatoryReason: {
		Code:        RejectReasonRegulatoryReason,
		Description: "Regulatory Reason",
	},
	RejectReasonRTPNotSupportedForDebtor: {
		Code:        RejectReasonRTPNotSupportedForDebtor,
		Description: "RTP Not Supported ForDebtor",
	},
	RejectReasonRTPServiceProviderIdentifierIncorrect: {
		Code:        RejectReasonRTPServiceProviderIdentifierIncorrect,
		Description: "RTP Service Provider Identifier Incorrect",
	},
	RejectReasonUnknownCreditor: {
		Code:        RejectReasonUnknownCreditor,
		Description: "Unknown creditor",
	},
	StatusRejectedCancellationRequest: {
		Code:        StatusRejectedCancellationRequest,
		Description: "Rejected Cancellation Request",
	},
	StatusCancelledAsPerRequest: {
		Code:        StatusCancelledAsPerRequest,
		Description: "Cancelled as per request",
	},
	RejectCancellationReasonAlreadyCancelledRTP: {
		Code:        RejectCancellationReasonAlreadyCancelledRTP,
		Description: "Request-to-pay has already been cancelled",
	},

	RejectCancellationReasonPaymentAlreadyTransmittedExecution: {
		Code:        RejectCancellationReasonPaymentAlreadyTransmittedExecution,
		Description: "Payment related to the request-to-pay has already been transmitted for execution",
	},
	RejectCancellationReasonUnknownRTP: {
		Code:        RejectCancellationReasonUnknownRTP,
		Description: "Request-to-pay is unknown",
	},
	RfCStatusUpdateAlreadyRejected: {
		Code:        RfCStatusUpdateAlreadyRejected,
		Description: "Already Rejected",
	},
	RfCStatusUpdateNeverReceived: {
		Code:        RfCStatusUpdateNeverReceived,
		Description: "Never received",
	},
	RfCStatusUpdateReceivedAndProcessed: {
		Code:        RfCStatusUpdateReceivedAndProcessed,
		Description: "Already Processed",
	},
}

func LookupStsCdRsnText(dataset string, cd string, rsn string) string {

	if rsn != "" {
		if s, ok := StsRsnCodeDescriptionRegistry[rsn]; ok {
			return s.Description
		}
	}

	if s, ok := StsRsnCodeDescriptionRegistry[cd]; ok {
		return s.Description
	}
	return ""
}

func IsTransactionStatusValidForDataset(ds string, txSts string) bool {

	var rc bool
	switch ds {
	case DataSetDS04:
		if util.StringArrayContains(listOfDS04ValidTransactionStatus, txSts) {
			rc = true
		}
	case DataSetDS05:
		if util.StringArrayContains(listOfDS05ValidTransactionStatus, txSts) {
			rc = true
		}
	case DataSetDS08:
		if util.StringArrayContains(listOfDS08ValidTransactionStatus, txSts) {
			rc = true
		}
	case DataSetDS12:
		if util.StringArrayContains(listOfDS12ValidTransactionStatus, txSts) {
			rc = true
		}
	case DataSetDS16:
		if util.StringArrayContains(listOfDS16ValidTransactionStatus, txSts) {
			rc = true
		}
	default:
		log.Error().Str("dataset", ds).Msg("unsupported dataset")
	}

	return rc
}

func IsTransactionStatusReasonValidForDataset(ds string, st Status) bool {

	var rc bool
	switch ds {
	case DataSetDS04:
		if util.StringArrayContains(listOfDS04ValidNegativeTransactionReasons, st.ReasonCode) {
			rc = true
		}
	case DataSetDS05:
		if st.ReasonCode == "" {
			rc = true
		}
	case DataSetDS08:
		if st.Code == StatusRejected {
			if util.StringArrayContains(listOfDS08ValidNegativeTransactionReasons, st.ReasonCode) {
				rc = true
			}
		} else if st.ReasonCode == "" {
			rc = true
		}
	case DataSetDS12:
		if st.Code == StatusRejectedCancellationRequest {
			if util.StringArrayContains(listOfDS12ValidNegativeTransactionReasons, st.ReasonCode) {
				rc = true
			}
		} else if st.ReasonCode == "" {
			rc = true
		}
	case DataSetDS16:
		if util.StringArrayContains(listOfDS16ValidTransactionReasons, st.ReasonCode) {
			rc = true
		}
	default:
		log.Error().Str("dataset", ds).Msg("unsupported dataset")
	}

	return rc
}

var listOfFinalTransactionStatus = []string{
	StatusRejected, StatusRejectedCancellationRequest, StatusAccepted, StatusCancelledAsPerRequest, StatusAcceptedWithChange,
}

func IsStatusFinal(txSts string) bool {
	return util.StringArrayContains(listOfFinalTransactionStatus, txSts)
}

func NextStatus(dataset string, current Status, event Status) (Status, bool) {

	var rc bool
	var st = current
	switch dataset {
	case DataSetDS04:
		if current.Code == StatusZero {
			rc = true
			st = event
		}
	case DataSetDS05:
		if current.Code == StatusZero {
			rc = true
			st = event
		}
	case DataSetDS08:
		if current.Code == StatusZero || current.Code == StatusTechnicallyAccepted {
			rc = true
			st = event
		}
	case DataSetDS12:
		// This condition refers to the status of the DS10 request....
		// Different issue when there is the need to propagate to the original rtp status.
		if current.Code == StatusZero {
			rc = true
			st = event
		}
	case DataSetDS16:
		switch event.ReasonCode {
		case StatusReasonAlreadyExpiredRTP:
			if current.Code == "" || current.Code == StatusTechnicallyAccepted || current.Code == StatusRejected {
				rc = true
				if current.Code != StatusRejected {
					st = Status{Code: StatusRejected, ReasonCode: event.ReasonCode}
				}
			}
		case StatusReasonAlreadyAcceptedRTP:
			if current.Code == StatusZero || current.Code == StatusTechnicallyAccepted || current.Code == StatusAccepted || current.Code == StatusAcceptedWithChange {
				rc = true
				if current.Code == StatusZero || current.Code == StatusTechnicallyAccepted {
					st = Status{Code: StatusAccepted, ReasonCode: ""}
				}
			}
		case StatusReasonAlreadyRefusedRTP:
			if current.Code == StatusZero || current.Code == StatusTechnicallyAccepted || current.Code == StatusRejected {
				rc = true
				if current.Code != StatusRejected {
					st = Status{Code: StatusRejected, ReasonCode: ""}
				}
			}

		case StatusReasonAlreadyRejectedRTP:
			if current.Code == StatusZero || current.Code == StatusTechnicallyAccepted || current.Code == StatusRejected {
				rc = true
				if current.Code != StatusRejected {
					st = Status{Code: StatusRejected, ReasonCode: ""}
				}
			}
		case StatusReasonInitialRTPNeverReceived:
			if current.Code == StatusZero {
				rc = true
				st = Status{Code: StatusRejected, ReasonCode: event.ReasonCode}
			}
		case StatusReasonRTPReceivedCanBeProcessed:
			if current.Code == StatusZero || current.Code == StatusTechnicallyAccepted {
				rc = true
			}
		}
	default:
		log.Error().Str("dataset", dataset).Str("current", current.Code).Str("event", event.Code).Msg("cannot compute wf-status on unrecognized dataset")

	}

	_ = st.UpdateText(dataset)
	return st, rc
}

func paymentInstruct(doc *Document) (paymentInstruct, error) {
	output := paymentInstruct{
		AdrLine:                doc.AdrLine.Max70Text,                                     // "<AdrLine>"
		Agt:                    doc.Agt.BranchAndFinancialInstitutionIdentification5,      // "<Agt>"
		Assgne:                 doc.Assgne.Party35Choice,                                  // "<Assgne>"
		Assgnr:                 doc.Assgnr.Party7Choice,                                   // "<Assgnr>"
		BICFI:                  doc.BICFI.BankIdentificationCode,                          // "<BICFI>"
		BizMsgIdr:              doc.BizMsgIdr.BusinessMessageIdentifier,                   // "<BizMsgIdr>"
		BldgNb:                 doc.BldgNb.Max16Text,                                      // "<BldgNb>"
		// CanonicalizationMethod https://docs.microsoft.com/en-us/dotnet/api/system.security.cryptography.xml.signedinfo.canonicalizationmethod?view=dotnet-plat-ext-6.0
		// CanonicalizationMethod https://docs.oracle.com/javase/8/docs/api/javax/xml/crypto/dsig/CanonicalizationMethod.html
		// CanonicalizationMethod https://docs.oracle.com/en/java/javase/13/docs/api/java.xml.crypto/javax/xml/crypto/dsig/CanonicalizationMethod.html
		// CanonicalizationMethod https://www.di-mgt.com.au/xmldsig-c14n.html
		CanonicalizationMethod: doc.CanonicalizationMethod.CanonicalizationMethod,         // "<CanonicalizationMethod>"
		CdtrAgt:                doc.CdtrAgt.BranchAndFinancialInstitutionIdentification4,  // "<CdtrAgt>"
		CdtTrfTxInf:            doc.CdtTrfTxInf.CreditTransferTransactionInformation11,    // "<CdtTrfTxInf>"
		ChrgBr:                 doc.ChrgBr.ChargeBearerType1Code,                          // "<ChrgBr>"
		ChrgsInf:               doc.ChrgsInf.ChargesInformation5,                          // "<ChrgsInf>"
		Conf:                   doc.Conf.ExternalInvestigationExecutionConfirmation1Co,    // "<Conf>"
		CreDt:                  doc.CreDt.ISONormalisedDateTime,                           // "<CreDt>"
		CreDtTm:                doc.CreDtTm.ISODateTime,                                   // "<CreDtTm>"
		Cretr:                  doc.Cretr.Party35Choice,                                   // "<Cretr>"
		Ctry:                   doc.Ctry.CountryCode,                                      // "<Ctry>"
		CxlRsnInf:              doc.CxlRsnInf.CancellationReasonInformation3,              // "<CxlRsnInf>"
		Dbtr:                   doc.Dbtr.PartyIdentification32,                            // "<Dbtr>"
		DbtrAgt:                doc.DbtrAgt.BranchAndFinancialInstitutionIdentification4,  // "<DbtrAgt>"
		// EndToEndId https://wiki.xmldation.com/Support/ISO20022/General_Rules/EndToEndId
		// EndToEndId https://www.jam-software.com/sepa-transfer/end-to-end-id.shtml
		// EndToEndId https://answers.sap.com/questions/12267089/element-endtoendid-not-filled-in-xml-payment-file.html
		// EndToEndId https://answers.sap.com/questions/10275743/dmee-%E2%80%93-endtoendid-with-paymantorder.html
		EndToEndId:             doc.EndToEndId.EndToEndId,                                 // "<EndToEndId>"
		Envlp:                  doc.Envlp.SupplementaryDataEnvelope1,                      // "<Envlp>"
		// FIId https://www.iso.org/iso-22000-food-safety-management.html
		// FIId https://www.qyriel.com/FullCatalogue/ISO_HEAD/out/ProtocolReport/xsd_head/head.001.001.01.xsd.html
		// FIId Financial Institution Identification
		// FIId AppHdr/Fr [Choice]
		FIId:                   doc.FIId.FinancialInstitutionIdentification,               // "<FIId>"
		// FinInstnId EPC limits the usage of Debtor Agent (DbtrAgt) and Creditor Agent CdtrAgt to allow only BIC and nothing else.
		// FinInstnId https://wiki.xmldation.com/Support/EPC/FinInstnId
		// FinInstnId https://wiki.xmldation.com/Support/RBS/CT_Rules/Global_Rules/CdtTrfTxInf%2F%2FCdtrAgt%2F%2FFinInstnId%2F%2FPstlAdr
		// FinInstnId CdtTrfTxInf/CdtrAgt/FinInstnId/PstlAdr Following fields from CreditorAgent / FinancialInstitutionIdentification / PostalAddress / Department '<CdtrAgt><FinInstnId><PstlAdr><Dept>'
		FinInstnId:             doc.FinInstnId.FinancialInstitutionIdentification7,        // "<FinInstnId>"
		FIToFICstmrCdtTrf:      doc.FIToFICstmrCdtTrf.FIToFICustomerCreditTransferV02,     // "<FIToFICstmrCdtTrf>"
		FIToFIPmtCxlReq:        doc.FIToFIPmtCxlReq.FIToFIPaymentCancellationRequestV01,   // "<FIToFIPmtCxlReq>"
		FIToFIPmtStsRpt:        doc.FIToFIPmtStsRpt.FIToFIPaymentStatusReportV03,          // "<FIToFIPmtStsRpt>"
		Fr:                     doc.Fr.Party9Choice,                                       // "<Fr>"
		GrpHdr:                 doc.GrpHdr.GroupHeader33,                                  // "<GrpHdr>"
		Id:                     doc.Id.Max35Text,                                          // "<Id>"
		InstdAgt:               doc.InstdAgt.BranchAndFinancialInstitutionIdentification4, // "<InstdAgt>"
		InstdAmt:               doc.InstdAmt.ActiveOrHistoricCurrencyAndAmount,            // "<InstdAmt>"
		// InstgAgt https://www.swift.com/swift-resource/248686/download
		// InstgAgt https://community.oracle.com/tech/developers/discussion/4327286/ora-00904-error-outer-join-19c
		// InstgAgt https://www.nacha.org/content/iso-20022-ach-mapping-guide
		// InstgAgt https://www.iso20022.org/sites/default/files/documents/D7/ISO20022_RTPG_pacs00800106_July_2017_v1_1.pdf
		InstgAgt:               doc.InstgAgt.BranchAndFinancialInstitutionIdentification4, // "<InstgAgt>"
		// InstrId https://wiki.xmldation.com/Support/ISO20022/General_Rules/InstrId
		// InstrId https://www.mathworks.com/help/instrument/instrid.html
		// InstrId https://wiki.xmldation.com/Support/Sampo/InstrId
		// InstrId https://docs.oracle.com/cd/E16582_01/doc.91/e15104/fields_sepa_pay_file_appx.htm#EOAEL01692
		InstrId:                doc.InstrId.InstructionIdentification,                     // "<InstrId>"
		// IntrBkSttlmAmt https://www.ecb.europa.eu/paym/groups/shared/docs/75299-tips-_cg_2017-09-28_presentation_udfs.pdf
		// IntrBkSttlmAmt https://wiki.xmldation.com/General_Information/ISO_20022/Difference_between_InstdAmt_and_EqvtAmt
		// IntrBkSttlmAmt https://www.iotafinance.com/en/SWIFT-ISO15022-Message-type-MT202-COV.html
		// IntrBkSttlmAmt https://www.bnymellon.com/content/dam/bnymellon/documents/pdf/iso-20022/Module%201_September%202020_Demystifying%20ISO20022.pdf
		IntrBkSttlmAmt:         doc.IntrBkSttlmAmt.ActiveOrHistoricCurrencyAndAmount,      // "<IntrBkSttlmAmt>"
		// IntrBkSttlmDt https://www.citibank.com/tts/sa/flippingbook/2021/ISO-20022-Citi-Mini-Series-and-Reference-Guide-Part-2/10/
		// IntrBkSttlmDt https://www.citibank.com/tts/sa/flippingbook/2021/ISO-20022-Citi-Mini-Series-and-Reference-Guide-Part-2/26/
		// IntrBkSttlmDt https://www.paymentstandards.ch/dam/mapping-rules_pacs008_esr.pdf
		// IntrBkSttlmDt https://www.payments.ca/sites/default/files/part_a_of_5_fitofi_customer_credit_transfers.pdf
		IntrBkSttlmDt:          doc.IntrBkSttlmDt.InterbankSettlementDate,                 // "<IntrBkSttlmDt>"
		Issr:                   doc.Issr.Issuer,                                           // "<Issr>"
		Justfn:                 doc.Justfn.CaseForwardingNotification3Code,                // "<Justfn>"
		KeyInfo:                doc.KeyInfo.KeyInfo,                                       // "<KeyInfo>"
		Mod:                    doc.Mod.RequestedModification7,                            // "<Mod>"
		MsgDefIdr:              doc.MsgDefIdr.MessageDefinitionIdentifier,                 // "<MsgDefIdr>"
		MsgId:                  doc.MsgId.MessageIdentification,                           // "<MsgId>"
		MssngOrIncrrctInf:      doc.MssngOrIncrrctInf.MissingOrIncorrectInformation3,      // "<MssngOrIncrrctInf>"
		// NbOfTxs https://wiki.xmldation.com/Support/RBS/DD_Rules/Global_Rules/NbOfTxs
		// NbOfTxs https://support.oracle.com/knowledge/Oracle%20E-Business%20Suite/1571592_1.html
		// NbOfTxs https://docs.oracle.com/cd/E16582_01/doc.91/e15104/fields_sepa_pay_file_appx.htm#EOAEL01692
		// NbOfTxs https://wiki.xmldation.com/Support/ISO20022/General_Rules/NbOfTxs
		NbOfTxs:                doc.NbOfTxs.Max15NumericText,                              // "<NbOfTxs>"
		NtfctnOfCaseAssgnmt:    doc.NtfctnOfCaseAssgnmt.NotificationOfCaseAssignmentV05,   // "<NtfctnOfCaseAssgnmt>"
		OrgnlCreDtTm:           doc.OrgnlCreDtTm.ISODateTime,                              // "<OrgnlCreDtTm>"
		// OrgnlEndToEndId https://wiki.xmldation.com/Support/ISO20022/General_Rules/EndToEndId
		// OrgnlEndToEndId https://paymentcomponents.atlassian.net/wiki/spaces/AH/pages/479428560/Sample+SEPA+messages+for+Testing
		// OrgnlEndToEndId https://answers.sap.com/questions/10275743/dmee-%E2%80%93-endtoendid-with-paymantorder.html
		// OrgnlEndToEndId https://blogs.sap.com/2021/07/30/pain.002-payment-rejections-processing-via-rfebka00/
		// OrgnlEndToEndId https://docs.crbcos.com/unicorncrb/docs/unicorn-output-files
		OrgnlEndToEndId:        doc.OrgnlEndToEndId.OriginalEndToEndIdentification,        // "<OrgnlEndToEndId>"
		// OrgnlGrpInf https://www.payments.ca/sites/default/files/part_c_of_5_payment_return.pdf
		// OrgnlGrpInf https://wiki.xmldation.com/Support/Nordea/CancellationRequest/Cancellation_Request_%2f%2f_CancellationReason2Code
		// OrgnlGrpInf https://www.iso20022.org/sites/default/files/documents/D7/Pacs004%20Real%20Time%20Payment%20Sep2018_v0.1.pdf
		// OrgnlGrpInf https://www.nacha.org/content/iso-20022-ach-mapping-guide
		// OrgnlGrpInf https://www.iso20022.org/sites/default/files/documents/D7/ISO20022_RTPG_pacs00200108_July_2017_v1_1.pdf
		OrgnlGrpInf:            doc.OrgnlGrpInf.OriginalGroupInformation3,                 // "<OrgnlGrpInf>"
		OrgnlGrpInfAndCxl:      doc.OrgnlGrpInfAndCxl.OriginalGroupInformation23,          // "<OrgnlGrpInfAndCxl>"
		OrgnlGrpInfAndSts:      doc.OrgnlGrpInfAndSts.OriginalGroupInformation20,          // "<OrgnlGrpInfAndSts>"
		OrgnlInstdAmt:          doc.OrgnlInstdAmt.OriginalInstructedAmount,                // "<OrgnlInstdAmt>"
		// OrgnlInstrId https://www.iso20022.org/sites/default/files/documents/D7/Pacs004%20Real%20Time%20Payment%20Sep2018_v0.1.pdf
		// OrgnlInstrId https://paymentcomponents.atlassian.net/wiki/spaces/AH/pages/479428560/Sample+SEPA+messages+for+Testing
		// OrgnlInstrId https://stackoverflow.com/questions/65199828/parsing-xml-in-c-sharp-with-xsd-file
		// OrgnlInstrId https://github.com/FasterXML/jackson-dataformat-xml/issues/217
		OrgnlInstrId:           doc.OrgnlInstrId.OriginalInstructionIdentification,        // "<OrgnlInstrId>"
		OrgnlIntrBkSttlmAmt:    doc.OrgnlIntrBkSttlmAmt.ActiveOrHistoricCurrencyAndAmount, // "<OrgnlIntrBkSttlmAmt>"
		OrgnlMsgId:             doc.OrgnlMsgId.OriginalMessageIdentification,              // "<OrgnlMsgId>"
		OrgnlMsgNmId:           doc.OrgnlMsgNmId.OriginalMessageNameIdentification,        // "<OrgnlMsgNmId>"
		OrgnlTxId:              doc.OrgnlTxId.OriginalTransactionIdentification,           // "<OrgnlTxId>"
		OrgnlTxRef:             doc.OrgnlTxRef.OriginalTransactionReference13,             // "<OrgnlTxRef>"
		Orgtr:                  doc.Orgtr.PartyIdentification32,                           // "<Orgtr>"
		PlcAndNm:               doc.PlcAndNm.PlcAndNm,                                     // "<PlcAndNm>"
		PmtTpInf:               doc.PmtTpInf.PaymentTypeInformation21,                     // "<PmtTpInf>"
		PstCd:                  doc.PstCd.Max16Text,                                       // "<PstCd>"
		PstlAdr:                doc.PstlAdr.PostalAddress6,                                // "<PstlAdr>"
		ReqToModfyPmt:          doc.ReqToModfyPmt.RequestToModifyPaymentV05,               // "<ReqToModfyPmt>"
		RsltnOfInvstgtn:        doc.RsltnOfInvstgtn.ResolutionOfInvestigationV08,          // "<RsltnOfInvstgtn>"
		RtrdInstdAmt:           doc.RtrdInstdAmt.ActiveOrHistoricCurrencyAndAmount,        // "<RtrdInstdAmt>"
		RtrdIntrBkSttlmAmt:     doc.RtrdIntrBkSttlmAmt.ActiveCurrencyAndAmount,            // "<RtrdIntrBkSttlmAmt>"
		RtrId:                  doc.RtrId.Max35Text,                                       // "<RtrId>"
		RtrRsnInf:              doc.RtrRsnInf.ReturnReasonInformation9,                    // "<RtrRsnInf>"
		Signature:              doc.Signature.Signature,                                   // "<Signature>"
		SignatureMethod:        doc.SignatureMethod.SignatureMethod,                       // "<SignatureMethod>"
		SplmtryData:            doc.SplmtryData.SupplementaryData1,                        // "<SplmtryData>"
		StrtNm:                 doc.StrtNm.Max70Text,                                      // "<StrtNm>"
		SttlmAcct:              doc.SttlmAcct.CashAccount16,                               // "<SttlmAcct>"
		SttlmInf:               doc.SttlmInf.SettlementInformation13,                      // "<SttlmInf>"
		SttlmMtd:               doc.SttlmMtd.SettlementMethod1Code,                        // "<SttlmMtd>"
		SvcLvl:                 doc.SvcLvl.ServiceLevel8Choice,                            // "<SvcLvl>"
		To:                     doc.To.Party9Choice,                                       // "<To>"
		TwnNm:                  doc.TwnNm.Max35Text,                                       // "<TwnNm>"
		TxId:                   doc.TxId.TransactionIdentification,                        // "<TxId>"
		TxInfAndSts:            doc.TxInfAndSts.PaymentTransactionInformation26,           // "<TxInfAndSts>"
		TxSts:                  doc.TxSts.TransactionIndividualStatus3Code,                // "<TxSts>"
		UblToApply:             doc.UblToApply.UnableToApplyV07,                           // "<UblToApply>"
		UltmtCdtr:              doc.UltmtCdtr.PartyIdentification32,                       // "<UltmtCdtr>"
		Undrlyg:                doc.Undrlyg.UnderlyingTransaction4Choice,                  // "<Undrlyg>"
		X509Data:               doc.X509Data.KeyInfoX509Data,                              // "<X509Data>"
		XchgRate:               doc.XchgRate.BaseOneRate,                                  // "<XchgRate>"
	}