package pain001

import "testing"

func TestNewCustomerCreditTransferInitiation(t *testing.T) {
	expected := `<CstmrCdtTrfInitn><GrpHdr><MsgId>7efb812b-c34d-4f58</MsgId><CreDtTm>2020-07-16T23:04:52</CreDtTm><NbOfTxs>1</NbOfTxs><CtrlSum>17.17</CtrlSum><InitgPty><Nm>George Goodman</Nm></InitgPty></GrpHdr><PmtInf><PmtInfId>PMTINFID-01</PmtInfId><PmtMtd>TRF</PmtMtd><BtchBookg>true</BtchBookg><ReqdExctnDt>2020-01-02</ReqdExctnDt><Dbtr><Nm>George Goodman</Nm><PstlAdr><StrtNm>Fnord Street</StrtNm><BldgNb>23</BldgNb><PstCd>2323</PstCd><TwnNm>Fnord Town</TwnNm><Ctry>FN</Ctry></PstlAdr></Dbtr><DbtrAcct><Id><IBAN>FN5614836012445688000</IBAN></Id></DbtrAcct><DbtrAgt><FinInstnId><BIC>EXP1234</BIC></FinInstnId></DbtrAgt><CdtTrfTxInf><PmtId><InstrId>INSTRID-01-1</InstrId><EndToEndId>ENDTOENDID-1</EndToEndId></PmtId><PmtTpInf><LclInstrm><Prtry>CH03</Prtry></LclInstrm></PmtTpInf><Amt><InstdAmt Ccy="CFN">17.17</InstdAmt></Amt><Cdtr><Nm>Joe Moon</Nm><PstlAdr><StrtNm>Muldon Street</StrtNm><BldgNb>17</BldgNb><PstCd>1717</PstCd><TwnNm>Steet Town</TwnNm><Ctry>FN</Ctry></PstlAdr></Cdtr><CdtrAcct><Id><IBAN>FN5604835012345678009</IBAN></Id></CdtrAcct><RmtInf><Ustrd>statuary</Ustrd></RmtInf></CdtTrfTxInf></PmtInf></CstmrCdtTrfInitn>`
	data, err := NewCustomerCreditTransferInitiation(Order{
		ExecuteOn: "2020-01-02",
		Debitor: Debitor{
			Name:     "George Goodman",
			Street:   "Fnord Street",
			StreetNr: 23,
			Postcode: 2323,
			Place:    "Fnord Town",
			Country:  "FN",
			IBAN:     "FN56 1483 6012 4456 8800 0",
			BIC:      "EXP1234",
		},
		Transactions: []Transaction{{
			Name:      "Joe Moon",
			Street:    "Muldon Street",
			StreetNr:  17,
			Postcode:  1717,
			Place:     "Steet Town",
			Country:   "FN",
			IBAN:      "FN56 0483 5012 3456 7800 9",
			Reference: "statuary",
			Amount:    "17.17",
			Currency:  "CFN"}}})
	if err != nil {
		t.Error(err)
	}
	data.GroupHeader.MessageIdentification = "7efb812b-c34d-4f58"
	data.GroupHeader.CreationDateTime = "2020-07-16T23:04:52"
	compareXmlResult(t, data, expected)
}
