package examples_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/camt_055_001_08"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/pain_013_001_07"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
	"os"
	"testing"
	"time"
)

const (
	NameOfInitiatingParty        = "AcmeCorp"
	InitiatingPartyOrgId         = "01114601006"
	InitiatingPartyOrgIssr       = "Acme CORP"
	PaymentInstructionId         = "Fatt2022-000001-2022-03-22-11:50:45"
	DebitorName                  = "LoremIpsumSPA"
	DebitorAddress               = "ViaLoremIpsum 30 Roma"
	DebitorIBAN                  = "IT60X0760111101000000123456"
	DebitorAgentBIC              = "BPPIITRRXXX"
	InvoiceNumber                = "fatt2022-000001"
	CreditorName                 = "AFC Poste Italiane"
	CreditorAgentBIC             = "BPPIITRRXXX"
	CreditorAddress              = "Via Del Creditore 75 Roma"
	CreditorOrgId                = "0468651441"
	CreditorIBAN                 = "IT60X0760111101000004545561"
	RemittanceInfo               = "AT41/fatt2022-000001/AT05/pagamento fattura"
	RemittanceInfoAT87           = "AT87/testo libero"
	AssignerName                 = "AFC Poste Italiane"
	AssignmentId                 = "str12345"
	AssignerOrgId                = "0468651441"
	AssigneeAgentBic             = "BPPIITRRXXX"
	OriginalPaymentInformationId = "Fatt2022-000001-2022-03-22-11:50:45"

	Example_Camt_055_001_08_1 = "example-document-camt_055_001_08.xml"
	Example_Pain_013_001_07_1 = "example-document-pain_013_001_07-1.xml"
	Example_Pain_013_001_07_2 = "example-document-pain_013_001_07-2.xml"
)

var pain_013_001_07_Document pain_013_001_07.Document
var camt_055_001_08_Document camt_055_001_08.Document

func TestMain(m *testing.M) {
	pain_013_001_07_Document = pain_013_001_07.Document{
		CdtrPmtActvtnReq: pain_013_001_07.CreditorPaymentActivationRequestV07{

			GrpHdr: pain_013_001_07.GroupHeader78{
				MsgId:   common.MustToMax35Text("pain013-DS01-20220322"),
				CreDtTm: common.MustToISODateTime(time.Now()),
				NbOfTxs: common.MustToMax15NumericText("1"),
				InitgPty: common.PartyIdentification135{
					Nm: common.MustToMax140Text(NameOfInitiatingParty),
					Id: &common.Party38Choice{
						OrgId: &common.OrganisationIdentification29{
							Othr: []common.GenericOrganisationIdentification1{
								{
									Id:   common.MustToMax35Text(InitiatingPartyOrgId),
									Issr: common.MustToMax35Text(InitiatingPartyOrgIssr),
								},
							},
						},
					},
				},
			},
			PmtInf: []pain_013_001_07.PaymentInstruction31{
				{
					PmtInfId: common.MustToMax35Text(PaymentInstructionId),
					PmtMtd:   common.MustToPaymentMethod7Code(common.PaymentMethod7CodeTRF),
					PmtTpInf: &common.PaymentTypeInformation26{
						SvcLvl: []common.ServiceLevel8Choice{
							{
								Cd: common.MustToExternalServiceLevel1Code("SRTP"),
							},
						},
						LclInstrm: &common.LocalInstrument2Choice{
							Prtry: common.MustToMax35Text("NOTPROVIDED"),
						},
						CtgyPurp: &common.CategoryPurpose1Choice{
							Cd: common.MustToExternalCategoryPurpose1Code("OTHR"),
						},
					},
					ReqdExctnDt: common.DateAndDateTime2Choice{
						Dt: common.MustToISODate(time.Now()),
					},
					XpryDt: &common.DateAndDateTime2Choice{
						Dt: common.MustToISODate(time.Now()),
					},
					Dbtr: common.PartyIdentification135{
						Nm: common.MustToMax140Text(DebitorName),
						PstlAdr: &common.PostalAddress24{
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(DebitorAddress),
							},
						},
						Id: &common.Party38Choice{
							PrvtId: &common.PersonIdentification13{
								Othr: []common.GenericPersonIdentification1{
									{
										Id: common.MustToMax35Text("123456789"),
										SchmeNm: &common.PersonIdentificationSchemeName1Choice{
											Cd: common.MustToExternalPersonIdentification1Code("POID"),
										},
									},
								},
							},
						},
					},
					DbtrAcct: &common.CashAccount38{
						Id: common.AccountIdentification4Choice{
							IBAN: common.MustToIBAN2007Identifier(DebitorIBAN),
						},
					},
					DbtrAgt: common.BranchAndFinancialInstitutionIdentification6{
						FinInstnId: common.FinancialInstitutionIdentification18{
							BICFI: common.MustToBICFIDec2014Identifier(DebitorAgentBIC),
						},
					},
					CdtTrfTx: []pain_013_001_07.CreditTransferTransaction35{
						{
							PmtId: pain_013_001_07.PaymentIdentification6{
								EndToEndId: common.MustToMax35Text(InvoiceNumber),
							},
							PmtTpInf: &common.PaymentTypeInformation26{
								SvcLvl: []common.ServiceLevel8Choice{
									{
										Cd: common.MustToExternalServiceLevel1Code("SRTP"),
									},
								},
								LclInstrm: &common.LocalInstrument2Choice{
									Prtry: common.MustToMax35Text("NOTPROVIDED"),
								},
								CtgyPurp: &common.CategoryPurpose1Choice{
									Cd: common.MustToExternalCategoryPurpose1Code("OTHR"),
								},
							},
							PmtCond: &common.PaymentCondition1{
								AmtModAllwd:   false,
								EarlyPmtAllwd: false,
								GrntedPmtReqd: false,
							},
							Amt: common.AmountType4Choice{
								InstdAmt: &common.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode("EUR"),
									Value: xsdt.MustToDecimal(535.35),
								},
							},
							ChrgBr: common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSLEV),
							CdtrAgt: common.BranchAndFinancialInstitutionIdentification6{
								FinInstnId: common.FinancialInstitutionIdentification18{
									BICFI: common.MustToBICFIDec2014Identifier(CreditorAgentBIC),
								},
							},
							Cdtr: common.PartyIdentification135{
								Nm: common.MustToMax140Text(CreditorName),
								PstlAdr: &common.PostalAddress24{
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(CreditorAddress),
									},
								},
								Id: &common.Party38Choice{
									OrgId: &common.OrganisationIdentification29{
										Othr: []common.GenericOrganisationIdentification1{
											{
												Id: common.MustToMax35Text(CreditorOrgId),
												SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
													Cd: common.MustToExternalOrganisationIdentification1Code("BOID"),
												},
											},
										},
									},
								},
							},
							CdtrAcct: &common.CashAccount38{
								Id: common.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(CreditorIBAN),
								},
							},
							RmtInf: &common.RemittanceInformation16{
								Ustrd: []common.Max140Text{
									common.MustToMax140Text(RemittanceInfo),
									common.MustToMax140Text(RemittanceInfoAT87),
								},
							},
						},
					},
				},
			},
		},
	}

	camt_055_001_08_Document = camt_055_001_08.Document{
		CstmrPmtCxlReq: camt_055_001_08.CustomerPaymentCancellationRequestV08{
			Assgnmt: common.CaseAssignment5{
				Id: common.MustToMax35Text(AssignmentId),
				Assgnr: common.Party40Choice{
					Pty: &common.PartyIdentification135{
						Nm: common.MustToMax140Text(AssignerName),
						Id: &common.Party38Choice{
							OrgId: &common.OrganisationIdentification29{
								Othr: []common.GenericOrganisationIdentification1{
									{
										Id: common.MustToMax35Text(AssignerOrgId),
										SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
											Cd: common.MustToExternalOrganisationIdentification1Code("BOID"),
										},
									},
								},
							},
						},
					},
				},
				Assgne: common.Party40Choice{
					Agt: &common.BranchAndFinancialInstitutionIdentification6{
						FinInstnId: common.FinancialInstitutionIdentification18{
							BICFI: common.MustToBICFIDec2014Identifier(AssigneeAgentBic),
						},
					},
				},
				CreDtTm: common.MustToISODateTime(time.Now()),
			},
			Undrlyg: []camt_055_001_08.UnderlyingTransaction24{
				{
					OrgnlGrpInfAndCxl: &camt_055_001_08.OriginalGroupHeader15{
						GrpCxlId:     common.MustToMax35Text("camt055-DS10-20220330"),
						OrgnlMsgId:   common.MustToMax35Text("pacs013-DS01-20220322"),
						OrgnlMsgNmId: common.MustToMax35Text("pain.013.001.07"),
						OrgnlCreDtTm: common.MustToISODateTime(time.Now()),
						NbOfTxs:      common.MustToMax15NumericText("1"),
						CtrlSum:      xsdt.MustToDecimal(535.25),
					},
					OrgnlPmtInfAndCxl: []camt_055_001_08.OriginalPaymentInstruction34{
						{
							OrgnlPmtInfId: common.MustToMax35Text(OriginalPaymentInformationId),
							TxInf: []camt_055_001_08.PaymentTransaction109{
								{
									CxlId:           common.MustToMax35Text("CancellationId1234"),
									OrgnlEndToEndId: common.MustToMax35Text(InvoiceNumber),
									CxlRsnInf: []camt_055_001_08.PaymentCancellationReason5{
										{
											Orgtr: &common.PartyIdentification135{
												Nm: common.MustToMax140Text(AssignerName),
												Id: &common.Party38Choice{
													OrgId: &common.OrganisationIdentification29{
														Othr: []common.GenericOrganisationIdentification1{
															{
																Id: common.MustToMax35Text(AssignerOrgId),
																SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
																	Cd: common.MustToExternalOrganisationIdentification1Code("BOID"),
																},
															},
														},
													},
												},
											},
											Rsn: &camt_055_001_08.CancellationReason33Choice{
												Cd: common.MustToExternalCancellationReason1Code("AM09"),
											},
											AddtlInf: []common.Max105Text{common.MustToMax105Text("error in previous invoice's amount")},
										},
									},
									OrgnlTxRef: &common.OriginalTransactionReference28{
										Amt: &common.AmountType4Choice{
											InstdAmt: &common.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode("EUR"),
												Value: xsdt.MustToDecimal(535.25),
											},
										},
										ReqdExctnDt: &common.DateAndDateTime2Choice{
											Dt: common.MustToISODate(time.Now()),
										},
										PmtTpInf: &common.PaymentTypeInformation27{
											SvcLvl: []common.ServiceLevel8Choice{
												{
													Cd: common.MustToExternalServiceLevel1Code("SEPA"),
												},
											},
											LclInstrm: &common.LocalInstrument2Choice{
												Prtry: common.MustToMax35Text("NOTPROVIDED"),
											},
										},
										RmtInf: &common.RemittanceInformation16{
											Ustrd: []common.Max140Text{
												common.MustToMax140Text(RemittanceInfo),
											},
										},
										Cdtr: &common.Party40Choice{
											Pty: &common.PartyIdentification135{
												Nm: common.MustToMax140Text(CreditorName),
												Id: &common.Party38Choice{
													OrgId: &common.OrganisationIdentification29{
														Othr: []common.GenericOrganisationIdentification1{
															{
																Id: common.MustToMax35Text(CreditorOrgId),
																SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
																	Cd: common.MustToExternalOrganisationIdentification1Code("BOID"),
																},
															},
														},
													},
												},
											},
										},
										CdtrAcct: &common.CashAccount38{
											Id: common.AccountIdentification4Choice{
												IBAN: common.MustToIBAN2007Identifier(CreditorIBAN),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}
