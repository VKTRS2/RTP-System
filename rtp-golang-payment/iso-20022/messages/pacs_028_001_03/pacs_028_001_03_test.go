// Package pacs_028_001_03_test
// Do not Edit. This stuff it's been automatically generated.
package pacs_028_001_03_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/pacs_028_001_03"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
	"github.com/stretchr/testify/require"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
)

const Example_pacs_028_001_03 = "example-document-pacs_028_001_03.xml"

func TestDocumentpacs_028_001_03(t *testing.T) {

	d := pacs_028_001_03.Document{
		FIToFIPmtStsReq: pacs_028_001_03.FIToFIPaymentStatusRequestV03{
			GrpHdr: pacs_028_001_03.GroupHeader91{
				MsgId:   common.MustToMax35Text(common.Max35TextSample),
				CreDtTm: common.MustToISODateTime(common.ISODateTimeSample),
				InstgAgt: &common.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: common.FinancialInstitutionIdentification18{
						BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
						ClrSysMmbId: &common.ClearingSystemMemberIdentification2{
							ClrSysId: &common.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &common.PostalAddress24{
							AdrTp: &common.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &common.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
						Othr: &common.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &common.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &common.BranchData3{
						Id:  common.MustToMax35Text(common.Max35TextSample),
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &common.PostalAddress24{
							AdrTp: &common.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &common.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
					},
				},
				InstdAgt: &common.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: common.FinancialInstitutionIdentification18{
						BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
						ClrSysMmbId: &common.ClearingSystemMemberIdentification2{
							ClrSysId: &common.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &common.PostalAddress24{
							AdrTp: &common.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &common.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
						Othr: &common.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &common.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &common.BranchData3{
						Id:  common.MustToMax35Text(common.Max35TextSample),
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &common.PostalAddress24{
							AdrTp: &common.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &common.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
					},
				},
			},
			OrgnlGrpInf: []pacs_028_001_03.OriginalGroupInformation27{{
				OrgnlMsgId:   common.MustToMax35Text(common.Max35TextSample),
				OrgnlMsgNmId: common.MustToMax35Text(common.Max35TextSample),
				OrgnlCreDtTm: common.MustToISODateTime(common.ISODateTimeSample),
				OrgnlNbOfTxs: common.MustToMax15NumericText(common.Max15NumericTextSample),
				OrgnlCtrlSum: xsdt.MustToDecimal(xsdt.DecimalSample)},
			},
			TxInf: []pacs_028_001_03.PaymentTransaction113{{
				StsReqId: common.MustToMax35Text(common.Max35TextSample),
				OrgnlGrpInf: &common.OriginalGroupInformation29{
					OrgnlMsgId:   common.MustToMax35Text(common.Max35TextSample),
					OrgnlMsgNmId: common.MustToMax35Text(common.Max35TextSample),
					OrgnlCreDtTm: common.MustToISODateTime(common.ISODateTimeSample),
				},
				OrgnlInstrId:    common.MustToMax35Text(common.Max35TextSample),
				OrgnlEndToEndId: common.MustToMax35Text(common.Max35TextSample),
				OrgnlTxId:       common.MustToMax35Text(common.Max35TextSample),
				OrgnlUETR:       common.MustToUUIDv4Identifier(common.UUIDv4IdentifierSample),
				AccptncDtTm:     common.MustToISODateTime(common.ISODateTimeSample),
				ClrSysRef:       common.MustToMax35Text(common.Max35TextSample),
				InstgAgt: &common.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: common.FinancialInstitutionIdentification18{
						BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
						ClrSysMmbId: &common.ClearingSystemMemberIdentification2{
							ClrSysId: &common.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &common.PostalAddress24{
							AdrTp: &common.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &common.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
						Othr: &common.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &common.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &common.BranchData3{
						Id:  common.MustToMax35Text(common.Max35TextSample),
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &common.PostalAddress24{
							AdrTp: &common.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &common.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
					},
				},
				InstdAgt: &common.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: common.FinancialInstitutionIdentification18{
						BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
						ClrSysMmbId: &common.ClearingSystemMemberIdentification2{
							ClrSysId: &common.ClearingSystemIdentification2Choice{
								Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							MmbId: common.MustToMax35Text(common.Max35TextSample),
						},
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &common.PostalAddress24{
							AdrTp: &common.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &common.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
						Othr: &common.GenericFinancialIdentification1{
							Id: common.MustToMax35Text(common.Max35TextSample),
							SchmeNm: &common.FinancialIdentificationSchemeName1Choice{
								Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Issr: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					BrnchId: &common.BranchData3{
						Id:  common.MustToMax35Text(common.Max35TextSample),
						LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
						Nm:  common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &common.PostalAddress24{
							AdrTp: &common.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &common.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
					},
				},
				OrgnlTxRef: &common.OriginalTransactionReference28{
					IntrBkSttlmAmt: &common.ActiveOrHistoricCurrencyAndAmount{
						Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Value: xsdt.MustToDecimal(xsdt.DecimalSample),
					},
					Amt: &common.AmountType4Choice{
						InstdAmt: &common.ActiveOrHistoricCurrencyAndAmount{
							Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Value: xsdt.MustToDecimal(xsdt.DecimalSample),
						},
						EqvtAmt: &common.EquivalentAmount2{
							Amt: common.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Value: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							CcyOfTrf: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						},
					},
					IntrBkSttlmDt: common.MustToISODate(common.ISODateSample),
					ReqdColltnDt:  common.MustToISODate(common.ISODateSample),
					ReqdExctnDt: &common.DateAndDateTime2Choice{
						Dt:   common.MustToISODate(common.ISODateSample),
						DtTm: common.MustToISODateTime(common.ISODateTimeSample),
					},
					CdtrSchmeId: &common.PartyIdentification135{
						Nm: common.MustToMax140Text(common.Max140TextSample),
						PstlAdr: &common.PostalAddress24{
							AdrTp: &common.AddressType3Choice{
								Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
								Prtry: &common.GenericIdentification30{
									Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
									Issr:    common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Dept:        common.MustToMax70Text(common.Max70TextSample),
							SubDept:     common.MustToMax70Text(common.Max70TextSample),
							StrtNm:      common.MustToMax70Text(common.Max70TextSample),
							BldgNb:      common.MustToMax16Text(common.Max16TextSample),
							BldgNm:      common.MustToMax35Text(common.Max35TextSample),
							Flr:         common.MustToMax70Text(common.Max70TextSample),
							PstBx:       common.MustToMax16Text(common.Max16TextSample),
							Room:        common.MustToMax70Text(common.Max70TextSample),
							PstCd:       common.MustToMax16Text(common.Max16TextSample),
							TwnNm:       common.MustToMax35Text(common.Max35TextSample),
							TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
							DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
							CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
							Ctry:        common.MustToCountryCode(common.CountryCodeSample),
							AdrLine: []common.Max70Text{
								common.MustToMax70Text(common.Max70TextSample),
							},
						},
						Id: &common.Party38Choice{
							OrgId: &common.OrganisationIdentification29{
								AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
								LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Othr: []common.GenericOrganisationIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
							PrvtId: &common.PersonIdentification13{
								DtAndPlcOfBirth: &common.DateAndPlaceOfBirth1{
									BirthDt:     common.MustToISODate(common.ISODateSample),
									PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
									CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
								},
								Othr: []common.GenericPersonIdentification1{{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &common.PersonIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample)},
								},
							},
						},
						CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
						CtctDtls: &common.Contact4{
							NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
							Nm:        common.MustToMax140Text(common.Max140TextSample),
							PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
							MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
							FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
							EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
							EmailPurp: common.MustToMax35Text(common.Max35TextSample),
							JobTitl:   common.MustToMax35Text(common.Max35TextSample),
							Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
							Dept:      common.MustToMax70Text(common.Max70TextSample),
							Othr: []common.OtherContact1{{
								ChanlTp: common.MustToMax4Text(common.Max4TextSample),
								Id:      common.MustToMax128Text(common.Max128TextSample)},
							},
							PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
						},
					},
					SttlmInf: &common.SettlementInstruction7{
						SttlmMtd: common.MustToSettlementMethod1Code(common.SettlementMethod1CodeSample),
						SttlmAcct: &common.CashAccount38{
							Id: common.AccountIdentification4Choice{
								IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
								Othr: &common.GenericAccountIdentification1{
									Id: common.MustToMax34Text(common.Max34TextSample),
									SchmeNm: &common.AccountSchemeName1Choice{
										Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Tp: &common.CashAccountType2Choice{
								Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Nm:  common.MustToMax70Text(common.Max70TextSample),
							Prxy: &common.ProxyAccountIdentification1{
								Tp: &common.ProxyAccountType1Choice{
									Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Id: common.MustToMax2048Text(common.Max2048TextSample),
							},
						},
						ClrSys: &common.ClearingSystemIdentification3Choice{
							Cd:    common.MustToExternalCashClearingSystem1Code(common.ExternalCashClearingSystem1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						InstgRmbrsmntAgt: &common.BranchAndFinancialInstitutionIdentification6{
							FinInstnId: common.FinancialInstitutionIdentification18{
								BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
								ClrSysMmbId: &common.ClearingSystemMemberIdentification2{
									ClrSysId: &common.ClearingSystemIdentification2Choice{
										Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									MmbId: common.MustToMax35Text(common.Max35TextSample),
								},
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Othr: &common.GenericFinancialIdentification1{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &common.FinancialIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							BrnchId: &common.BranchData3{
								Id:  common.MustToMax35Text(common.Max35TextSample),
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
							},
						},
						InstgRmbrsmntAgtAcct: &common.CashAccount38{
							Id: common.AccountIdentification4Choice{
								IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
								Othr: &common.GenericAccountIdentification1{
									Id: common.MustToMax34Text(common.Max34TextSample),
									SchmeNm: &common.AccountSchemeName1Choice{
										Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Tp: &common.CashAccountType2Choice{
								Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Nm:  common.MustToMax70Text(common.Max70TextSample),
							Prxy: &common.ProxyAccountIdentification1{
								Tp: &common.ProxyAccountType1Choice{
									Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Id: common.MustToMax2048Text(common.Max2048TextSample),
							},
						},
						InstdRmbrsmntAgt: &common.BranchAndFinancialInstitutionIdentification6{
							FinInstnId: common.FinancialInstitutionIdentification18{
								BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
								ClrSysMmbId: &common.ClearingSystemMemberIdentification2{
									ClrSysId: &common.ClearingSystemIdentification2Choice{
										Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									MmbId: common.MustToMax35Text(common.Max35TextSample),
								},
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Othr: &common.GenericFinancialIdentification1{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &common.FinancialIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							BrnchId: &common.BranchData3{
								Id:  common.MustToMax35Text(common.Max35TextSample),
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
							},
						},
						InstdRmbrsmntAgtAcct: &common.CashAccount38{
							Id: common.AccountIdentification4Choice{
								IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
								Othr: &common.GenericAccountIdentification1{
									Id: common.MustToMax34Text(common.Max34TextSample),
									SchmeNm: &common.AccountSchemeName1Choice{
										Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Tp: &common.CashAccountType2Choice{
								Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Nm:  common.MustToMax70Text(common.Max70TextSample),
							Prxy: &common.ProxyAccountIdentification1{
								Tp: &common.ProxyAccountType1Choice{
									Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Id: common.MustToMax2048Text(common.Max2048TextSample),
							},
						},
						ThrdRmbrsmntAgt: &common.BranchAndFinancialInstitutionIdentification6{
							FinInstnId: common.FinancialInstitutionIdentification18{
								BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
								ClrSysMmbId: &common.ClearingSystemMemberIdentification2{
									ClrSysId: &common.ClearingSystemIdentification2Choice{
										Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									MmbId: common.MustToMax35Text(common.Max35TextSample),
								},
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Othr: &common.GenericFinancialIdentification1{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &common.FinancialIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							BrnchId: &common.BranchData3{
								Id:  common.MustToMax35Text(common.Max35TextSample),
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
							},
						},
						ThrdRmbrsmntAgtAcct: &common.CashAccount38{
							Id: common.AccountIdentification4Choice{
								IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
								Othr: &common.GenericAccountIdentification1{
									Id: common.MustToMax34Text(common.Max34TextSample),
									SchmeNm: &common.AccountSchemeName1Choice{
										Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							Tp: &common.CashAccountType2Choice{
								Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
							Nm:  common.MustToMax70Text(common.Max70TextSample),
							Prxy: &common.ProxyAccountIdentification1{
								Tp: &common.ProxyAccountType1Choice{
									Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Id: common.MustToMax2048Text(common.Max2048TextSample),
							},
						},
					},
					PmtTpInf: &common.PaymentTypeInformation27{
						InstrPrty: common.MustToPriority2Code(common.Priority2CodeSample),
						ClrChanl:  common.MustToClearingChannel2Code(common.ClearingChannel2CodeSample),
						SvcLvl: []common.ServiceLevel8Choice{{
							Cd:    common.MustToExternalServiceLevel1Code(common.ExternalServiceLevel1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample)},
						},
						LclInstrm: &common.LocalInstrument2Choice{
							Cd:    common.MustToExternalLocalInstrument1Code(common.ExternalLocalInstrument1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						SeqTp: common.MustToSequenceType3Code(common.SequenceType3CodeSample),
						CtgyPurp: &common.CategoryPurpose1Choice{
							Cd:    common.MustToExternalCategoryPurpose1Code(common.ExternalCategoryPurpose1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
					},
					PmtMtd: common.MustToPaymentMethod4Code(common.PaymentMethod4CodeSample),
					MndtRltdInf: &common.MandateRelatedInformation14{
						MndtId:    common.MustToMax35Text(common.Max35TextSample),
						DtOfSgntr: common.MustToISODate(common.ISODateSample),
						AmdmntInd: xsdt.MustToBoolean(xsdt.BooleanSample),
						AmdmntInfDtls: &common.AmendmentInformationDetails13{
							OrgnlMndtId: common.MustToMax35Text(common.Max35TextSample),
							OrgnlCdtrSchmeId: &common.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Id: &common.Party38Choice{
									OrgId: &common.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []common.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &common.PersonIdentification13{
										DtAndPlcOfBirth: &common.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []common.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &common.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &common.Contact4{
									NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
									Nm:        common.MustToMax140Text(common.Max140TextSample),
									PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
									EmailPurp: common.MustToMax35Text(common.Max35TextSample),
									JobTitl:   common.MustToMax35Text(common.Max35TextSample),
									Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
									Dept:      common.MustToMax70Text(common.Max70TextSample),
									Othr: []common.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							OrgnlCdtrAgt: &common.BranchAndFinancialInstitutionIdentification6{
								FinInstnId: common.FinancialInstitutionIdentification18{
									BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
									ClrSysMmbId: &common.ClearingSystemMemberIdentification2{
										ClrSysId: &common.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &common.PostalAddress24{
										AdrTp: &common.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &common.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Othr: &common.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &common.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &common.BranchData3{
									Id:  common.MustToMax35Text(common.Max35TextSample),
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &common.PostalAddress24{
										AdrTp: &common.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &common.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
								},
							},
							OrgnlCdtrAgtAcct: &common.CashAccount38{
								Id: common.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &common.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &common.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &common.CashAccountType2Choice{
									Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
								Prxy: &common.ProxyAccountIdentification1{
									Tp: &common.ProxyAccountType1Choice{
										Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Id: common.MustToMax2048Text(common.Max2048TextSample),
								},
							},
							OrgnlDbtr: &common.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Id: &common.Party38Choice{
									OrgId: &common.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []common.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &common.PersonIdentification13{
										DtAndPlcOfBirth: &common.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []common.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &common.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &common.Contact4{
									NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
									Nm:        common.MustToMax140Text(common.Max140TextSample),
									PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
									EmailPurp: common.MustToMax35Text(common.Max35TextSample),
									JobTitl:   common.MustToMax35Text(common.Max35TextSample),
									Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
									Dept:      common.MustToMax70Text(common.Max70TextSample),
									Othr: []common.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							OrgnlDbtrAcct: &common.CashAccount38{
								Id: common.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &common.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &common.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &common.CashAccountType2Choice{
									Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
								Prxy: &common.ProxyAccountIdentification1{
									Tp: &common.ProxyAccountType1Choice{
										Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Id: common.MustToMax2048Text(common.Max2048TextSample),
								},
							},
							OrgnlDbtrAgt: &common.BranchAndFinancialInstitutionIdentification6{
								FinInstnId: common.FinancialInstitutionIdentification18{
									BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
									ClrSysMmbId: &common.ClearingSystemMemberIdentification2{
										ClrSysId: &common.ClearingSystemIdentification2Choice{
											Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										MmbId: common.MustToMax35Text(common.Max35TextSample),
									},
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &common.PostalAddress24{
										AdrTp: &common.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &common.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Othr: &common.GenericFinancialIdentification1{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &common.FinancialIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								BrnchId: &common.BranchData3{
									Id:  common.MustToMax35Text(common.Max35TextSample),
									LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Nm:  common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &common.PostalAddress24{
										AdrTp: &common.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &common.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
								},
							},
							OrgnlDbtrAgtAcct: &common.CashAccount38{
								Id: common.AccountIdentification4Choice{
									IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
									Othr: &common.GenericAccountIdentification1{
										Id: common.MustToMax34Text(common.Max34TextSample),
										SchmeNm: &common.AccountSchemeName1Choice{
											Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Tp: &common.CashAccountType2Choice{
									Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
								Nm:  common.MustToMax70Text(common.Max70TextSample),
								Prxy: &common.ProxyAccountIdentification1{
									Tp: &common.ProxyAccountType1Choice{
										Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Id: common.MustToMax2048Text(common.Max2048TextSample),
								},
							},
							OrgnlFnlColltnDt: common.MustToISODate(common.ISODateSample),
							OrgnlFrqcy: &common.Frequency36Choice{
								Tp: common.MustToFrequency6Code(common.Frequency6CodeSample),
								Prd: &common.FrequencyPeriod1{
									Tp:        common.MustToFrequency6Code(common.Frequency6CodeSample),
									CntPerPrd: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								PtInTm: &common.FrequencyAndMoment1{
									Tp:     common.MustToFrequency6Code(common.Frequency6CodeSample),
									PtInTm: common.MustToExact2NumericText(common.Exact2NumericTextSample),
								},
							},
							OrgnlRsn: &common.MandateSetupReason1Choice{
								Cd:    common.MustToExternalMandateSetupReason1Code(common.ExternalMandateSetupReason1CodeSample),
								Prtry: common.MustToMax70Text(common.Max70TextSample),
							},
							OrgnlTrckgDays: common.MustToExact2NumericText(common.Exact2NumericTextSample),
						},
						ElctrncSgntr: common.MustToMax1025Text(common.Max1025TextSample),
						FrstColltnDt: common.MustToISODate(common.ISODateSample),
						FnlColltnDt:  common.MustToISODate(common.ISODateSample),
						Frqcy: &common.Frequency36Choice{
							Tp: common.MustToFrequency6Code(common.Frequency6CodeSample),
							Prd: &common.FrequencyPeriod1{
								Tp:        common.MustToFrequency6Code(common.Frequency6CodeSample),
								CntPerPrd: xsdt.MustToDecimal(xsdt.DecimalSample),
							},
							PtInTm: &common.FrequencyAndMoment1{
								Tp:     common.MustToFrequency6Code(common.Frequency6CodeSample),
								PtInTm: common.MustToExact2NumericText(common.Exact2NumericTextSample),
							},
						},
						Rsn: &common.MandateSetupReason1Choice{
							Cd:    common.MustToExternalMandateSetupReason1Code(common.ExternalMandateSetupReason1CodeSample),
							Prtry: common.MustToMax70Text(common.Max70TextSample),
						},
						TrckgDays: common.MustToExact2NumericText(common.Exact2NumericTextSample),
					},
					RmtInf: &common.RemittanceInformation16{
						Ustrd: []common.Max140Text{
							common.MustToMax140Text(common.Max140TextSample),
						},
						Strd: []common.StructuredRemittanceInformation16{{
							RfrdDocInf: []common.ReferredDocumentInformation7{{
								Tp: &common.ReferredDocumentType4{
									CdOrPrtry: common.ReferredDocumentType3Choice{
										Cd:    common.MustToDocumentType6Code(common.DocumentType6CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
								Nb:     common.MustToMax35Text(common.Max35TextSample),
								RltdDt: common.MustToISODate(common.ISODateSample),
								LineDtls: []common.DocumentLineInformation1{{
									Id: []common.DocumentLineIdentification1{{
										Tp: &common.DocumentLineType1{
											CdOrPrtry: common.DocumentLineType1Choice{
												Cd:    common.MustToExternalDocumentLineType1Code(common.ExternalDocumentLineType1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample),
										},
										Nb:     common.MustToMax35Text(common.Max35TextSample),
										RltdDt: common.MustToISODate(common.ISODateSample)},
									},
									Desc: common.MustToMax2048Text(common.Max2048TextSample),
									Amt: &common.RemittanceAmount3{
										DuePyblAmt: &common.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										DscntApldAmt: []common.DiscountAmountAndType1{{
											Tp: &common.DiscountAmountType1Choice{
												Cd:    common.MustToExternalDiscountAmountType1Code(common.ExternalDiscountAmountType1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Amt: common.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											}},
										},
										CdtNoteAmt: &common.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										TaxAmt: []common.TaxAmountAndType1{{
											Tp: &common.TaxAmountType1Choice{
												Cd:    common.MustToExternalTaxAmountType1Code(common.ExternalTaxAmountType1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Amt: common.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											}},
										},
										AdjstmntAmtAndRsn: []common.DocumentAdjustment1{{
											Amt: common.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											},
											CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
											Rsn:       common.MustToMax4Text(common.Max4TextSample),
											AddtlInf:  common.MustToMax140Text(common.Max140TextSample)},
										},
										RmtdAmt: &common.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
									}},
								}},
							},
							RfrdDocAmt: &common.RemittanceAmount2{
								DuePyblAmt: &common.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								DscntApldAmt: []common.DiscountAmountAndType1{{
									Tp: &common.DiscountAmountType1Choice{
										Cd:    common.MustToExternalDiscountAmountType1Code(common.ExternalDiscountAmountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Amt: common.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									}},
								},
								CdtNoteAmt: &common.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								TaxAmt: []common.TaxAmountAndType1{{
									Tp: &common.TaxAmountType1Choice{
										Cd:    common.MustToExternalTaxAmountType1Code(common.ExternalTaxAmountType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Amt: common.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									}},
								},
								AdjstmntAmtAndRsn: []common.DocumentAdjustment1{{
									Amt: common.ActiveOrHistoricCurrencyAndAmount{
										Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
										Value: xsdt.MustToDecimal(xsdt.DecimalSample),
									},
									CdtDbtInd: common.MustToCreditDebitCode(common.CreditDebitCodeSample),
									Rsn:       common.MustToMax4Text(common.Max4TextSample),
									AddtlInf:  common.MustToMax140Text(common.Max140TextSample)},
								},
								RmtdAmt: &common.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
							},
							CdtrRefInf: &common.CreditorReferenceInformation2{
								Tp: &common.CreditorReferenceType2{
									CdOrPrtry: common.CreditorReferenceType1Choice{
										Cd:    common.MustToDocumentType3Code(common.DocumentType3CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
								Ref: common.MustToMax35Text(common.Max35TextSample),
							},
							Invcr: &common.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Id: &common.Party38Choice{
									OrgId: &common.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []common.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &common.PersonIdentification13{
										DtAndPlcOfBirth: &common.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []common.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &common.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &common.Contact4{
									NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
									Nm:        common.MustToMax140Text(common.Max140TextSample),
									PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
									EmailPurp: common.MustToMax35Text(common.Max35TextSample),
									JobTitl:   common.MustToMax35Text(common.Max35TextSample),
									Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
									Dept:      common.MustToMax70Text(common.Max70TextSample),
									Othr: []common.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							Invcee: &common.PartyIdentification135{
								Nm: common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Id: &common.Party38Choice{
									OrgId: &common.OrganisationIdentification29{
										AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
										LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
										Othr: []common.GenericOrganisationIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
									PrvtId: &common.PersonIdentification13{
										DtAndPlcOfBirth: &common.DateAndPlaceOfBirth1{
											BirthDt:     common.MustToISODate(common.ISODateSample),
											PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
											CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
										},
										Othr: []common.GenericPersonIdentification1{{
											Id: common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: &common.PersonIdentificationSchemeName1Choice{
												Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
												Prtry: common.MustToMax35Text(common.Max35TextSample),
											},
											Issr: common.MustToMax35Text(common.Max35TextSample)},
										},
									},
								},
								CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
								CtctDtls: &common.Contact4{
									NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
									Nm:        common.MustToMax140Text(common.Max140TextSample),
									PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
									MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
									EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
									EmailPurp: common.MustToMax35Text(common.Max35TextSample),
									JobTitl:   common.MustToMax35Text(common.Max35TextSample),
									Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
									Dept:      common.MustToMax70Text(common.Max70TextSample),
									Othr: []common.OtherContact1{{
										ChanlTp: common.MustToMax4Text(common.Max4TextSample),
										Id:      common.MustToMax128Text(common.Max128TextSample)},
									},
									PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
								},
							},
							TaxRmt: &common.TaxInformation7{
								Cdtr: &common.TaxParty1{
									TaxId:  common.MustToMax35Text(common.Max35TextSample),
									RegnId: common.MustToMax35Text(common.Max35TextSample),
									TaxTp:  common.MustToMax35Text(common.Max35TextSample),
								},
								Dbtr: &common.TaxParty2{
									TaxId:  common.MustToMax35Text(common.Max35TextSample),
									RegnId: common.MustToMax35Text(common.Max35TextSample),
									TaxTp:  common.MustToMax35Text(common.Max35TextSample),
									Authstn: &common.TaxAuthorisation1{
										Titl: common.MustToMax35Text(common.Max35TextSample),
										Nm:   common.MustToMax140Text(common.Max140TextSample),
									},
								},
								UltmtDbtr: &common.TaxParty2{
									TaxId:  common.MustToMax35Text(common.Max35TextSample),
									RegnId: common.MustToMax35Text(common.Max35TextSample),
									TaxTp:  common.MustToMax35Text(common.Max35TextSample),
									Authstn: &common.TaxAuthorisation1{
										Titl: common.MustToMax35Text(common.Max35TextSample),
										Nm:   common.MustToMax140Text(common.Max140TextSample),
									},
								},
								AdmstnZone: common.MustToMax35Text(common.Max35TextSample),
								RefNb:      common.MustToMax140Text(common.Max140TextSample),
								Mtd:        common.MustToMax35Text(common.Max35TextSample),
								TtlTaxblBaseAmt: &common.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								TtlTaxAmt: &common.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								Dt:    common.MustToISODate(common.ISODateSample),
								SeqNb: xsdt.MustToDecimal(xsdt.DecimalSample),
								Rcrd: []common.TaxRecord2{{
									Tp:       common.MustToMax35Text(common.Max35TextSample),
									Ctgy:     common.MustToMax35Text(common.Max35TextSample),
									CtgyDtls: common.MustToMax35Text(common.Max35TextSample),
									DbtrSts:  common.MustToMax35Text(common.Max35TextSample),
									CertId:   common.MustToMax35Text(common.Max35TextSample),
									FrmsCd:   common.MustToMax35Text(common.Max35TextSample),
									Prd: &common.TaxPeriod2{
										Yr: common.MustToISODate(common.ISODateSample),
										Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
										FrToDt: &common.DatePeriod2{
											FrDt: common.MustToISODate(common.ISODateSample),
											ToDt: common.MustToISODate(common.ISODateSample),
										},
									},
									TaxAmt: &common.TaxAmount2{
										Rate: xsdt.MustToDecimal(xsdt.DecimalSample),
										TaxblBaseAmt: &common.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										TtlAmt: &common.ActiveOrHistoricCurrencyAndAmount{
											Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
											Value: xsdt.MustToDecimal(xsdt.DecimalSample),
										},
										Dtls: []common.TaxRecordDetails2{{
											Prd: &common.TaxPeriod2{
												Yr: common.MustToISODate(common.ISODateSample),
												Tp: common.MustToTaxRecordPeriod1Code(common.TaxRecordPeriod1CodeSample),
												FrToDt: &common.DatePeriod2{
													FrDt: common.MustToISODate(common.ISODateSample),
													ToDt: common.MustToISODate(common.ISODateSample),
												},
											},
											Amt: common.ActiveOrHistoricCurrencyAndAmount{
												Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
												Value: xsdt.MustToDecimal(xsdt.DecimalSample),
											}},
										},
									},
									AddtlInf: common.MustToMax140Text(common.Max140TextSample)},
								},
							},
							GrnshmtRmt: &common.Garnishment3{
								Tp: common.GarnishmentType1{
									CdOrPrtry: common.GarnishmentType1Choice{
										Cd:    common.MustToExternalGarnishmentType1Code(common.ExternalGarnishmentType1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
								Grnshee: &common.PartyIdentification135{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &common.PostalAddress24{
										AdrTp: &common.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &common.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Id: &common.Party38Choice{
										OrgId: &common.OrganisationIdentification29{
											AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
											LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
											Othr: []common.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &common.PersonIdentification13{
											DtAndPlcOfBirth: &common.DateAndPlaceOfBirth1{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []common.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &common.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &common.Contact4{
										NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
										Nm:        common.MustToMax140Text(common.Max140TextSample),
										PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
										EmailPurp: common.MustToMax35Text(common.Max35TextSample),
										JobTitl:   common.MustToMax35Text(common.Max35TextSample),
										Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
										Dept:      common.MustToMax70Text(common.Max70TextSample),
										Othr: []common.OtherContact1{{
											ChanlTp: common.MustToMax4Text(common.Max4TextSample),
											Id:      common.MustToMax128Text(common.Max128TextSample)},
										},
										PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
									},
								},
								GrnshmtAdmstr: &common.PartyIdentification135{
									Nm: common.MustToMax140Text(common.Max140TextSample),
									PstlAdr: &common.PostalAddress24{
										AdrTp: &common.AddressType3Choice{
											Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
											Prtry: &common.GenericIdentification30{
												Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
												Issr:    common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: common.MustToMax35Text(common.Max35TextSample),
											},
										},
										Dept:        common.MustToMax70Text(common.Max70TextSample),
										SubDept:     common.MustToMax70Text(common.Max70TextSample),
										StrtNm:      common.MustToMax70Text(common.Max70TextSample),
										BldgNb:      common.MustToMax16Text(common.Max16TextSample),
										BldgNm:      common.MustToMax35Text(common.Max35TextSample),
										Flr:         common.MustToMax70Text(common.Max70TextSample),
										PstBx:       common.MustToMax16Text(common.Max16TextSample),
										Room:        common.MustToMax70Text(common.Max70TextSample),
										PstCd:       common.MustToMax16Text(common.Max16TextSample),
										TwnNm:       common.MustToMax35Text(common.Max35TextSample),
										TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
										DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
										CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
										Ctry:        common.MustToCountryCode(common.CountryCodeSample),
										AdrLine: []common.Max70Text{
											common.MustToMax70Text(common.Max70TextSample),
										},
									},
									Id: &common.Party38Choice{
										OrgId: &common.OrganisationIdentification29{
											AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
											LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
											Othr: []common.GenericOrganisationIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
										PrvtId: &common.PersonIdentification13{
											DtAndPlcOfBirth: &common.DateAndPlaceOfBirth1{
												BirthDt:     common.MustToISODate(common.ISODateSample),
												PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
												CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
											},
											Othr: []common.GenericPersonIdentification1{{
												Id: common.MustToMax35Text(common.Max35TextSample),
												SchmeNm: &common.PersonIdentificationSchemeName1Choice{
													Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
													Prtry: common.MustToMax35Text(common.Max35TextSample),
												},
												Issr: common.MustToMax35Text(common.Max35TextSample)},
											},
										},
									},
									CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
									CtctDtls: &common.Contact4{
										NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
										Nm:        common.MustToMax140Text(common.Max140TextSample),
										PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
										MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
										FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
										EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
										EmailPurp: common.MustToMax35Text(common.Max35TextSample),
										JobTitl:   common.MustToMax35Text(common.Max35TextSample),
										Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
										Dept:      common.MustToMax70Text(common.Max70TextSample),
										Othr: []common.OtherContact1{{
											ChanlTp: common.MustToMax4Text(common.Max4TextSample),
											Id:      common.MustToMax128Text(common.Max128TextSample)},
										},
										PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
									},
								},
								RefNb: common.MustToMax140Text(common.Max140TextSample),
								Dt:    common.MustToISODate(common.ISODateSample),
								RmtdAmt: &common.ActiveOrHistoricCurrencyAndAmount{
									Ccy:   common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
									Value: xsdt.MustToDecimal(xsdt.DecimalSample),
								},
								FmlyMdclInsrncInd: xsdt.MustToBoolean(xsdt.BooleanSample),
								MplyeeTermntnInd:  xsdt.MustToBoolean(xsdt.BooleanSample),
							},
							AddtlRmtInf: []common.Max140Text{
								common.MustToMax140Text(common.Max140TextSample),
							}},
						},
					},
					UltmtDbtr: &common.Party40Choice{
						Pty: &common.PartyIdentification135{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &common.PostalAddress24{
								AdrTp: &common.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &common.GenericIdentification30{
										Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
										Issr:    common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								BldgNm:      common.MustToMax35Text(common.Max35TextSample),
								Flr:         common.MustToMax70Text(common.Max70TextSample),
								PstBx:       common.MustToMax16Text(common.Max16TextSample),
								Room:        common.MustToMax70Text(common.Max70TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
								DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
							Id: &common.Party38Choice{
								OrgId: &common.OrganisationIdentification29{
									AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
									LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Othr: []common.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &common.PersonIdentification13{
									DtAndPlcOfBirth: &common.DateAndPlaceOfBirth1{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []common.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &common.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &common.Contact4{
								NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
								Nm:        common.MustToMax140Text(common.Max140TextSample),
								PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
								EmailPurp: common.MustToMax35Text(common.Max35TextSample),
								JobTitl:   common.MustToMax35Text(common.Max35TextSample),
								Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
								Dept:      common.MustToMax70Text(common.Max70TextSample),
								Othr: []common.OtherContact1{{
									ChanlTp: common.MustToMax4Text(common.Max4TextSample),
									Id:      common.MustToMax128Text(common.Max128TextSample)},
								},
								PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
							},
						},
						Agt: &common.BranchAndFinancialInstitutionIdentification6{
							FinInstnId: common.FinancialInstitutionIdentification18{
								BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
								ClrSysMmbId: &common.ClearingSystemMemberIdentification2{
									ClrSysId: &common.ClearingSystemIdentification2Choice{
										Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									MmbId: common.MustToMax35Text(common.Max35TextSample),
								},
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Othr: &common.GenericFinancialIdentification1{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &common.FinancialIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							BrnchId: &common.BranchData3{
								Id:  common.MustToMax35Text(common.Max35TextSample),
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
							},
						},
					},
					Dbtr: &common.Party40Choice{
						Pty: &common.PartyIdentification135{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &common.PostalAddress24{
								AdrTp: &common.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &common.GenericIdentification30{
										Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
										Issr:    common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								BldgNm:      common.MustToMax35Text(common.Max35TextSample),
								Flr:         common.MustToMax70Text(common.Max70TextSample),
								PstBx:       common.MustToMax16Text(common.Max16TextSample),
								Room:        common.MustToMax70Text(common.Max70TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
								DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
							Id: &common.Party38Choice{
								OrgId: &common.OrganisationIdentification29{
									AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
									LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Othr: []common.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &common.PersonIdentification13{
									DtAndPlcOfBirth: &common.DateAndPlaceOfBirth1{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []common.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &common.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &common.Contact4{
								NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
								Nm:        common.MustToMax140Text(common.Max140TextSample),
								PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
								EmailPurp: common.MustToMax35Text(common.Max35TextSample),
								JobTitl:   common.MustToMax35Text(common.Max35TextSample),
								Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
								Dept:      common.MustToMax70Text(common.Max70TextSample),
								Othr: []common.OtherContact1{{
									ChanlTp: common.MustToMax4Text(common.Max4TextSample),
									Id:      common.MustToMax128Text(common.Max128TextSample)},
								},
								PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
							},
						},
						Agt: &common.BranchAndFinancialInstitutionIdentification6{
							FinInstnId: common.FinancialInstitutionIdentification18{
								BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
								ClrSysMmbId: &common.ClearingSystemMemberIdentification2{
									ClrSysId: &common.ClearingSystemIdentification2Choice{
										Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									MmbId: common.MustToMax35Text(common.Max35TextSample),
								},
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Othr: &common.GenericFinancialIdentification1{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &common.FinancialIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							BrnchId: &common.BranchData3{
								Id:  common.MustToMax35Text(common.Max35TextSample),
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
							},
						},
					},
					DbtrAcct: &common.CashAccount38{
						Id: common.AccountIdentification4Choice{
							IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
							Othr: &common.GenericAccountIdentification1{
								Id: common.MustToMax34Text(common.Max34TextSample),
								SchmeNm: &common.AccountSchemeName1Choice{
									Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Tp: &common.CashAccountType2Choice{
							Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Nm:  common.MustToMax70Text(common.Max70TextSample),
						Prxy: &common.ProxyAccountIdentification1{
							Tp: &common.ProxyAccountType1Choice{
								Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Id: common.MustToMax2048Text(common.Max2048TextSample),
						},
					},
					DbtrAgt: &common.BranchAndFinancialInstitutionIdentification6{
						FinInstnId: common.FinancialInstitutionIdentification18{
							BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
							ClrSysMmbId: &common.ClearingSystemMemberIdentification2{
								ClrSysId: &common.ClearingSystemIdentification2Choice{
									Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								MmbId: common.MustToMax35Text(common.Max35TextSample),
							},
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &common.PostalAddress24{
								AdrTp: &common.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &common.GenericIdentification30{
										Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
										Issr:    common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								BldgNm:      common.MustToMax35Text(common.Max35TextSample),
								Flr:         common.MustToMax70Text(common.Max70TextSample),
								PstBx:       common.MustToMax16Text(common.Max16TextSample),
								Room:        common.MustToMax70Text(common.Max70TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
								DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
							Othr: &common.GenericFinancialIdentification1{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &common.FinancialIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						BrnchId: &common.BranchData3{
							Id:  common.MustToMax35Text(common.Max35TextSample),
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &common.PostalAddress24{
								AdrTp: &common.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &common.GenericIdentification30{
										Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
										Issr:    common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								BldgNm:      common.MustToMax35Text(common.Max35TextSample),
								Flr:         common.MustToMax70Text(common.Max70TextSample),
								PstBx:       common.MustToMax16Text(common.Max16TextSample),
								Room:        common.MustToMax70Text(common.Max70TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
								DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
						},
					},
					DbtrAgtAcct: &common.CashAccount38{
						Id: common.AccountIdentification4Choice{
							IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
							Othr: &common.GenericAccountIdentification1{
								Id: common.MustToMax34Text(common.Max34TextSample),
								SchmeNm: &common.AccountSchemeName1Choice{
									Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Tp: &common.CashAccountType2Choice{
							Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Nm:  common.MustToMax70Text(common.Max70TextSample),
						Prxy: &common.ProxyAccountIdentification1{
							Tp: &common.ProxyAccountType1Choice{
								Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Id: common.MustToMax2048Text(common.Max2048TextSample),
						},
					},
					CdtrAgt: &common.BranchAndFinancialInstitutionIdentification6{
						FinInstnId: common.FinancialInstitutionIdentification18{
							BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
							ClrSysMmbId: &common.ClearingSystemMemberIdentification2{
								ClrSysId: &common.ClearingSystemIdentification2Choice{
									Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								MmbId: common.MustToMax35Text(common.Max35TextSample),
							},
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &common.PostalAddress24{
								AdrTp: &common.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &common.GenericIdentification30{
										Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
										Issr:    common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								BldgNm:      common.MustToMax35Text(common.Max35TextSample),
								Flr:         common.MustToMax70Text(common.Max70TextSample),
								PstBx:       common.MustToMax16Text(common.Max16TextSample),
								Room:        common.MustToMax70Text(common.Max70TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
								DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
							Othr: &common.GenericFinancialIdentification1{
								Id: common.MustToMax35Text(common.Max35TextSample),
								SchmeNm: &common.FinancialIdentificationSchemeName1Choice{
									Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						BrnchId: &common.BranchData3{
							Id:  common.MustToMax35Text(common.Max35TextSample),
							LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
							Nm:  common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &common.PostalAddress24{
								AdrTp: &common.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &common.GenericIdentification30{
										Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
										Issr:    common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								BldgNm:      common.MustToMax35Text(common.Max35TextSample),
								Flr:         common.MustToMax70Text(common.Max70TextSample),
								PstBx:       common.MustToMax16Text(common.Max16TextSample),
								Room:        common.MustToMax70Text(common.Max70TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
								DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
						},
					},
					CdtrAgtAcct: &common.CashAccount38{
						Id: common.AccountIdentification4Choice{
							IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
							Othr: &common.GenericAccountIdentification1{
								Id: common.MustToMax34Text(common.Max34TextSample),
								SchmeNm: &common.AccountSchemeName1Choice{
									Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Tp: &common.CashAccountType2Choice{
							Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Nm:  common.MustToMax70Text(common.Max70TextSample),
						Prxy: &common.ProxyAccountIdentification1{
							Tp: &common.ProxyAccountType1Choice{
								Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Id: common.MustToMax2048Text(common.Max2048TextSample),
						},
					},
					Cdtr: &common.Party40Choice{
						Pty: &common.PartyIdentification135{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &common.PostalAddress24{
								AdrTp: &common.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &common.GenericIdentification30{
										Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
										Issr:    common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								BldgNm:      common.MustToMax35Text(common.Max35TextSample),
								Flr:         common.MustToMax70Text(common.Max70TextSample),
								PstBx:       common.MustToMax16Text(common.Max16TextSample),
								Room:        common.MustToMax70Text(common.Max70TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
								DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
							Id: &common.Party38Choice{
								OrgId: &common.OrganisationIdentification29{
									AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
									LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Othr: []common.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &common.PersonIdentification13{
									DtAndPlcOfBirth: &common.DateAndPlaceOfBirth1{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []common.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &common.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &common.Contact4{
								NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
								Nm:        common.MustToMax140Text(common.Max140TextSample),
								PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
								EmailPurp: common.MustToMax35Text(common.Max35TextSample),
								JobTitl:   common.MustToMax35Text(common.Max35TextSample),
								Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
								Dept:      common.MustToMax70Text(common.Max70TextSample),
								Othr: []common.OtherContact1{{
									ChanlTp: common.MustToMax4Text(common.Max4TextSample),
									Id:      common.MustToMax128Text(common.Max128TextSample)},
								},
								PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
							},
						},
						Agt: &common.BranchAndFinancialInstitutionIdentification6{
							FinInstnId: common.FinancialInstitutionIdentification18{
								BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
								ClrSysMmbId: &common.ClearingSystemMemberIdentification2{
									ClrSysId: &common.ClearingSystemIdentification2Choice{
										Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									MmbId: common.MustToMax35Text(common.Max35TextSample),
								},
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Othr: &common.GenericFinancialIdentification1{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &common.FinancialIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							BrnchId: &common.BranchData3{
								Id:  common.MustToMax35Text(common.Max35TextSample),
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
							},
						},
					},
					CdtrAcct: &common.CashAccount38{
						Id: common.AccountIdentification4Choice{
							IBAN: common.MustToIBAN2007Identifier(common.IBAN2007IdentifierSample),
							Othr: &common.GenericAccountIdentification1{
								Id: common.MustToMax34Text(common.Max34TextSample),
								SchmeNm: &common.AccountSchemeName1Choice{
									Cd:    common.MustToExternalAccountIdentification1Code(common.ExternalAccountIdentification1CodeSample),
									Prtry: common.MustToMax35Text(common.Max35TextSample),
								},
								Issr: common.MustToMax35Text(common.Max35TextSample),
							},
						},
						Tp: &common.CashAccountType2Choice{
							Cd:    common.MustToExternalCashAccountType1Code(common.ExternalCashAccountType1CodeSample),
							Prtry: common.MustToMax35Text(common.Max35TextSample),
						},
						Ccy: common.MustToActiveOrHistoricCurrencyCode(common.ActiveOrHistoricCurrencyCodeSample),
						Nm:  common.MustToMax70Text(common.Max70TextSample),
						Prxy: &common.ProxyAccountIdentification1{
							Tp: &common.ProxyAccountType1Choice{
								Cd:    common.MustToExternalProxyAccountType1Code(common.ExternalProxyAccountType1CodeSample),
								Prtry: common.MustToMax35Text(common.Max35TextSample),
							},
							Id: common.MustToMax2048Text(common.Max2048TextSample),
						},
					},
					UltmtCdtr: &common.Party40Choice{
						Pty: &common.PartyIdentification135{
							Nm: common.MustToMax140Text(common.Max140TextSample),
							PstlAdr: &common.PostalAddress24{
								AdrTp: &common.AddressType3Choice{
									Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
									Prtry: &common.GenericIdentification30{
										Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
										Issr:    common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: common.MustToMax35Text(common.Max35TextSample),
									},
								},
								Dept:        common.MustToMax70Text(common.Max70TextSample),
								SubDept:     common.MustToMax70Text(common.Max70TextSample),
								StrtNm:      common.MustToMax70Text(common.Max70TextSample),
								BldgNb:      common.MustToMax16Text(common.Max16TextSample),
								BldgNm:      common.MustToMax35Text(common.Max35TextSample),
								Flr:         common.MustToMax70Text(common.Max70TextSample),
								PstBx:       common.MustToMax16Text(common.Max16TextSample),
								Room:        common.MustToMax70Text(common.Max70TextSample),
								PstCd:       common.MustToMax16Text(common.Max16TextSample),
								TwnNm:       common.MustToMax35Text(common.Max35TextSample),
								TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
								DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
								CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
								Ctry:        common.MustToCountryCode(common.CountryCodeSample),
								AdrLine: []common.Max70Text{
									common.MustToMax70Text(common.Max70TextSample),
								},
							},
							Id: &common.Party38Choice{
								OrgId: &common.OrganisationIdentification29{
									AnyBIC: common.MustToAnyBICDec2014Identifier(common.AnyBICDec2014IdentifierSample),
									LEI:    common.MustToLEIIdentifier(common.LEIIdentifierSample),
									Othr: []common.GenericOrganisationIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalOrganisationIdentification1Code(common.ExternalOrganisationIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
								PrvtId: &common.PersonIdentification13{
									DtAndPlcOfBirth: &common.DateAndPlaceOfBirth1{
										BirthDt:     common.MustToISODate(common.ISODateSample),
										PrvcOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CityOfBirth: common.MustToMax35Text(common.Max35TextSample),
										CtryOfBirth: common.MustToCountryCode(common.CountryCodeSample),
									},
									Othr: []common.GenericPersonIdentification1{{
										Id: common.MustToMax35Text(common.Max35TextSample),
										SchmeNm: &common.PersonIdentificationSchemeName1Choice{
											Cd:    common.MustToExternalPersonIdentification1Code(common.ExternalPersonIdentification1CodeSample),
											Prtry: common.MustToMax35Text(common.Max35TextSample),
										},
										Issr: common.MustToMax35Text(common.Max35TextSample)},
									},
								},
							},
							CtryOfRes: common.MustToCountryCode(common.CountryCodeSample),
							CtctDtls: &common.Contact4{
								NmPrfx:    common.MustToNamePrefix2Code(common.NamePrefix2CodeSample),
								Nm:        common.MustToMax140Text(common.Max140TextSample),
								PhneNb:    common.MustToPhoneNumber(common.PhoneNumberSample),
								MobNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								FaxNb:     common.MustToPhoneNumber(common.PhoneNumberSample),
								EmailAdr:  common.MustToMax2048Text(common.Max2048TextSample),
								EmailPurp: common.MustToMax35Text(common.Max35TextSample),
								JobTitl:   common.MustToMax35Text(common.Max35TextSample),
								Rspnsblty: common.MustToMax35Text(common.Max35TextSample),
								Dept:      common.MustToMax70Text(common.Max70TextSample),
								Othr: []common.OtherContact1{{
									ChanlTp: common.MustToMax4Text(common.Max4TextSample),
									Id:      common.MustToMax128Text(common.Max128TextSample)},
								},
								PrefrdMtd: common.MustToPreferredContactMethod1Code(common.PreferredContactMethod1CodeSample),
							},
						},
						Agt: &common.BranchAndFinancialInstitutionIdentification6{
							FinInstnId: common.FinancialInstitutionIdentification18{
								BICFI: common.MustToBICFIDec2014Identifier(common.BICFIDec2014IdentifierSample),
								ClrSysMmbId: &common.ClearingSystemMemberIdentification2{
									ClrSysId: &common.ClearingSystemIdentification2Choice{
										Cd:    common.MustToExternalClearingSystemIdentification1Code(common.ExternalClearingSystemIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									MmbId: common.MustToMax35Text(common.Max35TextSample),
								},
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
								Othr: &common.GenericFinancialIdentification1{
									Id: common.MustToMax35Text(common.Max35TextSample),
									SchmeNm: &common.FinancialIdentificationSchemeName1Choice{
										Cd:    common.MustToExternalFinancialInstitutionIdentification1Code(common.ExternalFinancialInstitutionIdentification1CodeSample),
										Prtry: common.MustToMax35Text(common.Max35TextSample),
									},
									Issr: common.MustToMax35Text(common.Max35TextSample),
								},
							},
							BrnchId: &common.BranchData3{
								Id:  common.MustToMax35Text(common.Max35TextSample),
								LEI: common.MustToLEIIdentifier(common.LEIIdentifierSample),
								Nm:  common.MustToMax140Text(common.Max140TextSample),
								PstlAdr: &common.PostalAddress24{
									AdrTp: &common.AddressType3Choice{
										Cd: common.MustToAddressType2Code(common.AddressType2CodeSample),
										Prtry: &common.GenericIdentification30{
											Id:      common.MustToExact4AlphaNumericText(common.Exact4AlphaNumericTextSample),
											Issr:    common.MustToMax35Text(common.Max35TextSample),
											SchmeNm: common.MustToMax35Text(common.Max35TextSample),
										},
									},
									Dept:        common.MustToMax70Text(common.Max70TextSample),
									SubDept:     common.MustToMax70Text(common.Max70TextSample),
									StrtNm:      common.MustToMax70Text(common.Max70TextSample),
									BldgNb:      common.MustToMax16Text(common.Max16TextSample),
									BldgNm:      common.MustToMax35Text(common.Max35TextSample),
									Flr:         common.MustToMax70Text(common.Max70TextSample),
									PstBx:       common.MustToMax16Text(common.Max16TextSample),
									Room:        common.MustToMax70Text(common.Max70TextSample),
									PstCd:       common.MustToMax16Text(common.Max16TextSample),
									TwnNm:       common.MustToMax35Text(common.Max35TextSample),
									TwnLctnNm:   common.MustToMax35Text(common.Max35TextSample),
									DstrctNm:    common.MustToMax35Text(common.Max35TextSample),
									CtrySubDvsn: common.MustToMax35Text(common.Max35TextSample),
									Ctry:        common.MustToCountryCode(common.CountryCodeSample),
									AdrLine: []common.Max70Text{
										common.MustToMax70Text(common.Max70TextSample),
									},
								},
							},
						},
					},
					Purp: &common.Purpose2Choice{
						Cd:    common.MustToExternalPurpose1Code(common.ExternalPurpose1CodeSample),
						Prtry: common.MustToMax35Text(common.Max35TextSample),
					},
				},
				SplmtryData: []common.SupplementaryData1{{
					PlcAndNm: common.MustToMax350Text(common.Max350TextSample),
					Envlp: common.SupplementaryDataEnvelope1{
						Item: xsdt.MustToAnyType(xsdt.AnyTypeSample),
					}},
				}},
			},
			SplmtryData: []common.SupplementaryData1{{
				PlcAndNm: common.MustToMax350Text(common.Max350TextSample),
				Envlp: common.SupplementaryDataEnvelope1{
					Item: xsdt.MustToAnyType(xsdt.AnyTypeSample),
				}},
			},
		},
	}

	b, err := d.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_pacs_028_001_03, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_pacs_028_001_03)

}
