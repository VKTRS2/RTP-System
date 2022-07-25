package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//Transfers hey
type Transfers struct {
	TotalAmount float64    `xml:"CstmrCdtTrfInitn>GrpHdr>CtrlSum"`
	Iban        string     `xml:"CstmrCdtTrfInitn>PmtInf>DbtrAcct>Id>IBAN"`
	Transfers   []Transfer `xml:"CstmrCdtTrfInitn>PmtInf>CdtTrfTxInf"`
}

// Transfer hey
type Transfer struct {
	Amount          float64 `xml:"Amt>InstdAmt"`
	BeneficiaryName string  `xml:"Cdtr>Nm"`
	Description     string  `xml:"RmtInf>Ustrd"`
	Bic             string  `xml:"CdtrAgt>FinInstnId>BIC"`
	Iban            string  `xml:"CdtrAcct>Id>IBAN"`
}

func (t Transfer) String() string {
	return fmt.Sprintf("%f - %s - %s - %s - %s", t.Amount, t.BeneficiaryName, t.Description, t.Bic, t.Iban)
}

func readFile(f string) []byte {

	xmlFile, err := os.Open(f)
	checkErr(err)
	defer xmlFile.Close()

	fmt.Println("Successfully Opened: ", f)

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var transfers Transfers

	xml.Unmarshal(byteValue, &transfers)

	fmt.Printf("Total amount: %f - From %s \n", transfers.TotalAmount, transfers.Iban)
	for _, transfer := range transfers.Transfers {
		fmt.Printf("\t%s\n", transfer)
	}

	data, _ := json.Marshal(transfers)

	return data
}
