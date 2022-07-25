# golang Swiss ISO 20022 Pain.001 payment order library [![GoDoc](https://godoc.org/github.com/72nd/go-iso20022-pain001?status.svg)](https://godoc.org/github.com/72nd/go-iso20022-pain001) [![Go Report Card](https://goreportcard.com/badge/github.com/72nd/go-iso20022-pain001)](https://goreportcard.com/report/github.com/72nd/go-iso20022-pain001) 

**Remark:** This code is extracted from another project of mine and fulfills _our_ ([Genossenschaft Solutionsbüro](https://buero.io)) daily needs of automated payment order generation. It does _not_ implement the whole specification of the [Swiss standard](https://www.six-group.com/dam/download/banking-services/interbank-clearing/en/standardization/iso/swiss-recommendations/implementation-guidelines-ct.pdf). Further there is no safety net whatsoever. This comes without warranty of any kind. As this is about your money, think first. 

## Usage

A simple example which should illustrate the way of working with the library. For more information about some data format (like date, country-codes or currency-codes) please refer to the [documentation](https://godoc.org/github.com/72nd/go-iso20022-pain001) or the [offical implementation guidelines](https://www.six-group.com/dam/download/banking-services/interbank-clearing/en/standardization/iso/swiss-recommendations/implementation-guidelines-ct.pdf).

```golang
// Define the debitor (aka originator) of the transactions.
debitor := Debitor {
	Name:     "George Goodman",
	Street:   "Fnord Street",
	StreetNr: 5,
	Postcode: 2323,
	Place:    "Fnord Town",
	Country:  "FN",
	IBAN:     "FN56 1483 6012 4456 8800 0",
	BIC:      "EXP1234"
}

// Each Order can contain one or more transaction to different creditors (receivers of the money).
transactions := []Transaction{
	{
		Name:      "Joe Moon",
		Street:    "Muldon Street",
		StreetNr:  17,
		Postcode:  1717,
		Place:     "Steet Town",
		Country:   "FN",
		IBAN:      "FN56 0483 5012 3456 7800 9",
		Reference: "statuary",
		Amount:    "17.17",
		Currency:  "CFN",
	}
}

// Combine all parts into the actual bank order.
order := Order{
	ExecuteOn:    "2020-01-02",
	Debitor:      debitor,
	Transactions: transactions,
}

// Get the XML payment order and save it to the file.
xml, err := order.PaymentOrder()
if err != nil {
	log.Fatal(err)
}
if err := ioutil.WriteFile("payment-order.xml", xml, 0644); err != nil {
	log.Fatal(err)
}
```


## Sample Output

Example of a output file:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="http://www.six-interbank-clearing.com/de/pain.001.001.03.ch.02.xsd" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.six-interbank-clearing.com/de/pain.001.001.03.ch.02.xsd  pain.001.001.03.ch.02.xsd">
   <CstmrCdtTrfInitn>
      <GrpHdr>
         <MsgId>016dfe3b-6c81-428c</MsgId>
         <CreDtTm>2020-07-16T19:59:25</CreDtTm>
         <NbOfTxs>1</NbOfTxs>
         <CtrlSum>17.17</CtrlSum>
         <InitgPty>
            <Nm>George Goodman</Nm>
         </InitgPty>
      </GrpHdr>
      <PmtInf>
         <PmtInfId>PMTINFID-01</PmtInfId>
         <PmtMtd>TRF</PmtMtd>
         <BtchBookg>true</BtchBookg>
         <ReqdExctnDt>2020-01-02</ReqdExctnDt>
         <Dbtr>
            <Nm>George Goodman</Nm>
            <PstlAdr>
               <StrtNm>Fnord Street</StrtNm>
               <BldgNb>23</BldgNb>
               <PstCd>2323</PstCd>
               <TwnNm>Fnord Town</TwnNm>
               <Ctry>FN</Ctry>
            </PstlAdr>
         </Dbtr>
         <DbtrAcct>
            <Id>
               <IBAN>FN5614836012445688000</IBAN>
            </Id>
         </DbtrAcct>
         <DbtrAgt>
            <FinInstnId>
               <BIC>EXP1234</BIC>
            </FinInstnId>
         </DbtrAgt>
         <CdtTrfTxInf>
            <PmtId>
               <InstrId>INSTRID-01-1</InstrId>
               <EndToEndId>ENDTOENDID-1</EndToEndId>
            </PmtId>
            <PmtTpInf>
               <LclInstrm>
                  <Prtry>CH03</Prtry>
               </LclInstrm>
            </PmtTpInf>
            <Amt>
               <InstdAmt Ccy="CFN">17.17</InstdAmt>
            </Amt>
            <Cdtr>
               <Nm>Joe Moon</Nm>
               <PstlAdr>
                  <StrtNm>Muldon Street</StrtNm>
                  <BldgNb>17</BldgNb>
                  <PstCd>1717</PstCd>
                  <TwnNm>Steet Town</TwnNm>
                  <Ctry>FN</Ctry>
               </PstlAdr>
            </Cdtr>
            <CdtrAcct>
               <Id>
                  <IBAN>FN5604835012345678009</IBAN>
               </Id>
            </CdtrAcct>
            <RmtInf>
               <Ustrd>statuary</Ustrd>
            </RmtInf>
         </CdtTrfTxInf>
      </PmtInf>
   </CstmrCdtTrfInitn>
</Document>
```

