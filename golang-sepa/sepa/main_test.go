package sepa

import (
	"strings"
	"testing"
)

func TestCumul(t *testing.T) {
	var s = &Document{}
	s.InitDoc("", "2017-05-01T22:45:03", "2017-05-01", "", "FR1420041010050500013M02606", "")
	TTest := []float64{55, 140, 77, 105, 140, 76.3, 164.8, 62.3, 29.3, 125.3, 70, 78.22, 252.9, 35, 70, 173.6, 60.9, 63, 126, 215.6, 12.5, 35, 257.6, 75, 30, 72.5, 259.5, 302.62, 120.4, 35, 173.6, 104.54, 119, 22.5, 80.5, 135.8, 161.85, 1199.86, 32.5, 70, 140, 633.92, 159.6, 35, 196, 97.3, 90.3, 144.9, 258.7, 374.13, 27.5, 1575, 282.1, 56, 105, 57.4, 51.8, 56, 801.5, 66.99, 98.5, 212.8, 35, 109.9, 35, 269.5, 327.6, 224, 38.5, 35, 266, 256.2, 102.9, 201.6, 0.34, 35, 35, 341.6, 21, 217, 35.1, 19, 114, 25, 277.9, 70, 140, 21, 67.5, 41.3, 134.4, 143.36, 74, 21, 24, 27.07, 208.6, 43.75, 70, 58.8, 38.15, 61.5, 147, 378.8, 16.5, 52.5, 24.5, 60.2, 72.84, 175, 17.5, 70, 231.6, 161, 49, 70, 45.5, 291.2, 41.3, 35, 186.2, 154, 70, 35, 70, 35, 230, 119, 70, 20, 70, 175, 36.5, 217, 35, 52, 31.3, 109.2, 35, 24.5, 13.5, 63.5, 111.3, 60.2, 103, 203, 143.5, 35, 57.5, 35, 125.3, 175, 138.6, 153.82, 120.4, 62.5, 35.52, 63.5, 129.5, 70, 175, 224, 70, 126, 140, 35, 140, 25.5, 7.98, 70, 35, 65.2, 105, 77, 35, 98, 225.5, 38.5, 35, 158, 72.8, 147, 50, 210, 385, 28, 202.3, 128.8, 39.2, 117.6, 326, 30}
	cumul := float64(24443.66)
	//TTest := []float64{593.3, 164.8}
	//cumul := float64(758.1)
	for _, m := range TTest {
		s.AddTransaction("", m, "EUR", "", "GB29NWBK60161331926819")
	}
	if s.GroupheaderCtrlSum != cumul {
		t.Error("Expected GroupheaderCtrlSum", cumul, "got", s.GroupheaderCtrlSum)
	}
}
func TestDecimalsNumber(t *testing.T) {
	suite := []struct {
		f float64
		n int
	}{
		{0, 0},
		{123.0, 0},
		{144.2, 1},
		{1.123456789, 9},
		{3.1415900000, 5},
		{-1250, 0},
		{-252123.123, 3},
	}
	for _, s := range suite {
		received := DecimalsNumber(s.f)
		expected := s.n
		if received != expected {
			t.Errorf("Expected %v received %v", expected, received)
		}
	}
}
func TestGenerateSepaXML(t *testing.T) {

	var err error

	// targetDoc is a verified valid sepa xml file
	var targetDoc = `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pain.001.001.03" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><CstmrCdtTrfInitn><GrpHdr><MsgId>VIR201705</MsgId><CreDtTm>2017-05-01T12:00:00</CreDtTm><NbOfTxs>5</NbOfTxs><CtrlSum>170000</CtrlSum><InitgPty><Nm>Franz Holzapfel GMBH</Nm></InitgPty></GrpHdr><PmtInf><PmtInfId>VIR201705</PmtInfId><PmtMtd>TRF</PmtMtd><NbOfTxs>5</NbOfTxs><CtrlSum>170000</CtrlSum><PmtTpInf><SvcLvl><Cd>SEPA</Cd></SvcLvl></PmtTpInf><ReqdExctnDt>2017-05-03</ReqdExctnDt><Dbtr><Nm>Franz Holzapfel GMBH</Nm></Dbtr><DbtrAcct><Id><IBAN>AT611904300234573201</IBAN></Id></DbtrAcct><DbtrAgt><FinInstnId><BIC>BKAUATWW</BIC></FinInstnId></DbtrAgt><ChrgBr>SLEV</ChrgBr><CdtTrfTxInf><PmtId><InstrId>F201705</InstrId><EndToEndId>F201705</EndToEndId></PmtId><Amt><InstdAmt Ccy="EUR">70000</InstdAmt></Amt><Cdtr><Nm>DEF Electronics</Nm></Cdtr><CdtrAcct><Id><IBAN>GB29NWBK60161331926819</IBAN></Id></CdtrAcct><RgltryRptg><Dtls><Cd>150</Cd></Dtls></RgltryRptg><RmtInf><Ustrd>F201705</Ustrd></RmtInf></CdtTrfTxInf><CdtTrfTxInf><PmtId><InstrId>F201706</InstrId><EndToEndId>F201706</EndToEndId></PmtId><Amt><InstdAmt Ccy="EUR">10000</InstdAmt></Amt><Cdtr><Nm>D1F Electronics</Nm></Cdtr><CdtrAcct><Id><IBAN>AT611904300234573201</IBAN></Id></CdtrAcct><RgltryRptg><Dtls><Cd>150</Cd></Dtls></RgltryRptg><RmtInf><Ustrd>F201706</Ustrd></RmtInf></CdtTrfTxInf><CdtTrfTxInf><PmtId><InstrId>F201707</InstrId><EndToEndId>F201707</EndToEndId></PmtId><Amt><InstdAmt Ccy="EUR">20000</InstdAmt></Amt><Cdtr><Nm>D2F Electronics</Nm></Cdtr><CdtrAcct><Id><IBAN>BE62510007547061</IBAN></Id></CdtrAcct><RgltryRptg><Dtls><Cd>150</Cd></Dtls></RgltryRptg><RmtInf><Ustrd>F201707</Ustrd></RmtInf></CdtTrfTxInf><CdtTrfTxInf><PmtId><InstrId>F201708</InstrId><EndToEndId>F201708</EndToEndId></PmtId><Amt><InstdAmt Ccy="EUR">30000</InstdAmt></Amt><Cdtr><Nm>D3F Electronics</Nm></Cdtr><CdtrAcct><Id><IBAN>BG80BNBG96611020345678</IBAN></Id></CdtrAcct><RgltryRptg><Dtls><Cd>150</Cd></Dtls></RgltryRptg><RmtInf><Ustrd>F201708</Ustrd></RmtInf></CdtTrfTxInf><CdtTrfTxInf><PmtId><InstrId>F201709</InstrId><EndToEndId>F201709</EndToEndId></PmtId><Amt><InstdAmt Ccy="EUR">40000</InstdAmt></Amt><Cdtr><Nm>D4F Electronics</Nm></Cdtr><CdtrAcct><Id><IBAN>EE382200221020145685</IBAN></Id></CdtrAcct><RgltryRptg><Dtls><Cd>150</Cd></Dtls></RgltryRptg><RmtInf><Ustrd>F201709</Ustrd></RmtInf></CdtTrfTxInf></PmtInf></CstmrCdtTrfInitn></Document>`

	// our doc
	var sepaDoc = &Document{}

	// Bad format for creation date, expecting AAAA-MM-JJTHH:HH:SS
	err = sepaDoc.InitDoc("", "2017-05-01", "", "", "", "")
	if err == nil {
		t.Error("Expected InitDoc return an error for bad creation date format", "got", err)
	}

	// Bad format for execution date, expecting AAAA-MM-JJ
	err = sepaDoc.InitDoc("", "2017-05-01T22:45:03", "2017-05-01T22:45:03", "", "", "")
	if err == nil {
		t.Error("Expected InitDoc return an error for bad execution date format", "got", err)
	}

	// Bad IBAN
	err = sepaDoc.InitDoc("", "2017-05-01T22:45:03", "2017-05-01", "", "XX12345678901234567", "")
	if err == nil {
		t.Error("Expected InitDoc return an error for bad IBAN", "got", err)
	}

	// Good IBAN
	err = sepaDoc.InitDoc("", "2017-05-01T22:45:03", "2017-05-01", "", "FR1420041010050500013M02606", "")
	if err != nil {
		t.Error("Expected InitDoc return nil for good IBAN", "got", err)
	}

	// Initialize doc test
	err = sepaDoc.InitDoc("VIR201705", "2017-05-01T12:00:00", "2017-05-03", "Franz Holzapfel GMBH", "AT611904300234573201", "BKAUATWW")
	if err != nil {
		t.Error("Expected InitDoc return nil", "got", err)
	}

	// Add Transaction with incorrect IBAN
	err = sepaDoc.AddTransaction("XXX", 0, "XXX", "XXX", "ZZ382200221020145685")
	if err == nil {
		t.Error("Exepected AddTransaction return an error for bad IBAN", "got", err)
	}

	// Add Transaction with incorrect amount (>2 decimals)
	err = sepaDoc.AddTransaction("XXX", 1.234, "XXX", "XXX", "EE382200221020145685")
	if err == nil {
		t.Error("Exepected AddTransaction return an error for bad amount", "got", err)
	}

	// Transactions Test Array
	type testTransac struct {
		id         string
		amount     float64
		currency   string
		debtorName string
		debtorIban string
	}
	TTest := []testTransac{
		{"F201705", 70000, "EUR", "DEF Electronics", "GB29NWBK60161331926819"},
		{"F201706", 10000, "EUR", "D1F Electronics", "AT611904300234573201"},
		{"F201707", 20000, "EUR", "D2F Electronics", "BE62510007547061"},
		{"F201708", 30000, "EUR", "D3F Electronics", "BG80BNBG96611020345678"},
		{"F201709", 40000, "EUR", "D4F Electronics", "EE382200221020145685"},
	}

	// For each transaction, we check that the cumul amount and number of transactions remain correct in header and payment block
	var cumul = float64(0)

	var nb = 0
	for count, transac := range TTest {
		err = sepaDoc.AddTransaction(transac.id, transac.amount, transac.currency, transac.debtorName, transac.debtorIban)
		if err != nil {
			t.Error("Expected AddTransaction return nil", "got", err)
		}
		cumul += transac.amount
		nb = 1 + count
		if sepaDoc.GroupheaderCtrlSum != cumul {
			t.Error("Expected GroupheaderCtrlSum", cumul, "got", sepaDoc.GroupheaderCtrlSum)
		}
		if sepaDoc.PaymentInfoCtrlSum != cumul {
			t.Error("Expected PaymentInfoCtrlSum", cumul, "got", sepaDoc.PaymentInfoCtrlSum)
		}
		if sepaDoc.GroupheaderTransacNb != nb {
			t.Error("Expected GroupheaderTransacNb", nb, "got", sepaDoc.GroupheaderTransacNb)
		}
		if sepaDoc.PaymentInfoTransacNb != nb {
			t.Error("Expected PaymentInfoTransacNb", nb, "got", sepaDoc.PaymentInfoTransacNb)
		}
	}

	// Get the result
	str, err := sepaDoc.Serialize()
	if err != nil {
		t.Error("Expected xml in []byte, got ", err)
	}
	// Ultimate test : compare the all generated doc with the predefined doc
	res := strings.Compare(string(str), targetDoc)
	if res != 0 {
		t.Error("Expected", string(targetDoc), "got", string(str))
	}
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