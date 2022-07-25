package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/j03hanafi/bankiso/iso20022/pacs"

	"github.com/antchfx/xmlquery"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Found error in log ", err)
	}
	log.SetOutput(file)
	log.Println("Log setup")

	path := pathHandler()

	address := ":6067"
	log.Printf("Biller started at %v", address)
	err = http.ListenAndServe(address, path)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func pathHandler() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/biller", biller).Methods("POST")
	router.HandleFunc("/biller3", biller3).Methods("POST")

	return router
}

func biller3(w http.ResponseWriter, r *http.Request) {

	log.Println("New Request from BIFast Connector XML")
	fmt.Println("New Request from BIFast Connector XML")

	body, _ := ioutil.ReadAll(r.Body)
	rawRequest := string(body)
	log.Println(rawRequest)

	request := BusMsg{}
	err := xml.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
	}
	fmt.Println(request)
	appHdr := request.AppHdr
	bizMsgIdr := fmt.Sprintf("%v", *appHdr.BusinessMessageIdentifier)
	trxType := bizMsgIdr[16:19]
	fmt.Println(trxType)
	var fileName string

	switch trxType {
	case "510":
		type Doc struct {
			XMLName  xml.Name              `xml:"BusMsg"`
			Document pacs.Document00800108 `xml:"Document"`
		}
		document := Doc{}
		err := xml.Unmarshal(body, &document)
		if err != nil {
			log.Printf("Error unmarshal JSON: %s", err.Error())
		}

		CrAccId := *document.Document.Message.CreditTransferTransactionInformation[0].CdtrAcct.Id.Other.Identification
		fmt.Println(CrAccId)
		fileName = "accEnqRes.xml"
	}
	fileName = "xml/" + fileName
	fmt.Println("filename:", fileName)
	file, _ := os.Open(fileName)
	defer file.Close()

	response, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	responseFormatter2(w, response, 200)

}

func biller(w http.ResponseWriter, r *http.Request) {
	log.Println("New Request from BIFast Connector")
	fmt.Println("New Request from BIFast Connector")

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

	requestRaw := BusMsg{}
	err := xml.Unmarshal(body, &requestRaw)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
	}
	fmt.Println("request: ", requestRaw)

	// var msgID string
	var fileName string
	bzMsgID := fmt.Sprintf("%v", *requestRaw.AppHdr.BusinessMessageIdentifier)
	trxType := bzMsgID[16:19]
	fmt.Println("trxType:", trxType)

	switch trxType {
	case "000":
		payload := networkRequest(body)
		responseFormatter(w, payload, 200)
	// ##################### Account Enquiry ##################################
	case "510":
		var CrAccId string
		doc, err := xmlquery.Parse(strings.NewReader(string(body)))
		if err != nil {
			panic(err)
		}

		root := xmlquery.FindOne(doc, "//ns:BusMsg")
		if n := root.SelectElement("//ns:Document/ns:FIToFICstmrCdtTrf/ns2:CdtTrfTxInf/ns2:CdtrAcct/ns2:Id/ns2:Othr/ns2:Id"); n != nil {
			fmt.Printf("Name #%s\n", n.InnerText())
			CrAccId = n.InnerText()
		}

		switch CrAccId {
		case "510654300":
			fileName = "accountEnquiryResponse.xml"
		case "511654182":
			fileName = "accountEnquiryResponse2.xml"
		}

	//##################### Credit Transfer ###################################
	case "010": // Credit Transfer

		var CrAccId string
		doc, err := xmlquery.Parse(strings.NewReader(string(body)))
		if err != nil {
			panic(err)
		}

		root := xmlquery.FindOne(doc, "//ns:BusMsg")
		if n := root.SelectElement("//ns:Document/ns:FIToFICstmrCdtTrf/ns2:CdtTrfTxInf/ns2:CdtrAcct/ns2:Id/ns2:Othr/ns2:Id"); n != nil {
			fmt.Printf("Name #%s\n", n.InnerText())
			CrAccId = n.InnerText()
		}
		switch CrAccId {
		case "0102345600":
			fileName = "creditTransferResponse.xml"
		case "0102345184":
			fileName = "creditTransferResponse2.xml"
		}
	case "012":
		fileName = "sampleCreditTransferResponse012.json"
		fmt.Println("012")
	case "110":
		fileName = "sampleCreditTransferResponsewithProxy.json"
		fmt.Println("110")
	//==========================================================================

	case "019":
		fileName = "sampleFItoFICreditTransfer.json"
		fmt.Println("019")
	case "011":
		fileName = "sampleReverseCreditTransfer.json"
		fmt.Println("011")

	// ################# Proxy Resolution #####################################
	case "610":

		var PxValue string
		doc, err := xmlquery.Parse(strings.NewReader(string(body)))
		if err != nil {
			panic(err)
		}

		root := xmlquery.FindOne(doc, "//ns:BusMsg")
		if n := root.SelectElement("//ns:Document/ns:PrxyLookUp/ns2:LookUp/ns2:PrxyOnly/ns2:PrxyRtrvl/ns2:Val"); n != nil {
			fmt.Printf("Name #%s\n", n.InnerText())
			PxValue = n.InnerText()
		}
		switch PxValue {
		case "086102345000":
			fileName = "proxyResolutionResponse.xml"
		case "086112345101":
			fileName = "proxyResolutionResponse2.xml"
		case "086112345804":
			fileName = "proxyResolutionResponse3.xml"
		case "086132345600":
			fileName = "proxyResolutionResponse4.xml"
		case "086142345804":
			fileName = "proxyResolutionResponse5.xml"
		case "08615234804":
			fileName = "proxyResolutionResponse6.xml"
		case "08616234811":
			fileName = "proxyResolutionResponse7.xml"
		case "08617234805":
			fileName = "proxyResolutionResponse8.xml"
		}
	case "611":
		fileName = "sampleProxyResolution611.json"
		fmt.Println("611")
	case "612":
		fileName = "sampleProxyResolution612.json"
		fmt.Println(("612"))
	// =========================================================================

	// ################# Proxy Registration Inquiry ############################
	case "620":

		var PxRegId string
		doc, err := xmlquery.Parse(strings.NewReader(string(body)))
		if err != nil {
			panic(err)
		}

		root := xmlquery.FindOne(doc, "//ns:BusMsg")
		if n := root.SelectElement("//ns:Document/ns:PrxyNqryReq/ns2:GrpHdr/ns2:MsgSndr/ns2:Acct/ns2:Id/ns2:Othr/ns2:Id"); n != nil {
			fmt.Printf("Name #%s\n", n.InnerText())
			PxRegId = n.InnerText()
		}
		fmt.Println(PxRegId)
		switch PxRegId {
		case "6202345600":
			fileName = "proxyRegistrationResponse.xml"
		case "6212345101":
			fileName = "proxyRegistrationResponse2.xml"
		case "6222345808":
			fileName = "proxyRegistrationResponse3.xml"
		case "6232345600":
			fileName = "proxyRegistrationResponse4.xml"
		case "6242345600":
			fileName = "proxyRegistrationResponse5.xml"
		case "6252345808":
			fileName = "proxyRegistrationResponse6.xml"
		case "6262345806":
			fileName = "proxyRegistrationResponse7.xml"
		}
	case "621":
		fileName = "sampleProxyRegistrationInquiry621.json"
		fmt.Println("621")
	case "622":
		fileName = "sampleProxyRegistrationInquiry622.json"
		fmt.Println("622")
	// =========================================================================

	case "710":
		fileName = "sampleRegisterNewProxy.json"
		fmt.Println("710")

	//#################### Proxy Maintenance ###################################
	case "720":

		var PxValue string
		doc, err := xmlquery.Parse(strings.NewReader(string(body)))
		if err != nil {
			panic(err)
		}

		root := xmlquery.FindOne(doc, "//ns:BusMsg")
		if n := root.SelectElement("//ns:Document/ns:PrxyRegn/ns2:Regn/ns2:Prxy/ns2:Val"); n != nil {
			fmt.Printf("Name #%s\n", n.InnerText())
			PxValue = n.InnerText()
		}
		switch PxValue {
		case "7202345600":
			fileName = "proxyMaintenanceResponse.xml"
		case "7212345101":
			fileName = "proxyMaintenanceResponse2.xml"
		}
	case "721":
		fileName = "sampleProxyMaintenance721.json"
		fmt.Println("721")
		//============================================================================
	}
	if trxType != "000" {
		fileName = "xml/" + fileName
		fmt.Println("filename:", fileName)

		file, _ := os.Open(fileName)
		defer file.Close()

		response, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		responseFormatter(w, response, 200)
	}

}

func responseFormatter(w http.ResponseWriter, data []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func responseFormatter2(w http.ResponseWriter, data []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func networkRequest(data []byte) []byte {
	request := Admn1{}
	err := xml.Unmarshal(data, &request)
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println(request.Document.AdmnReq.AdmnTxInf.FnctnCd)

	response := Admn2{}
	response.Ns = "urn:iso"
	response.Ns1 = "urn:iso:std:iso:20022:tech:xsd:head.001.001.01"
	response.Ns2 = "urn:iso:std:iso:20022:tech:xsd:admn.002.001.01"
	response.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	response.SchemaLocation = "urn:iso ../../../../xsd/phase1/MainCIHub.xsd "

	response.AppHdr.Fr.FIId.FinInstnId.Othr.ID = request.AppHdr.Fr.FIId.FinInstnId.Othr.ID
	response.AppHdr.To.FIId.FinInstnId.Othr.ID = request.AppHdr.To.FIId.FinInstnId.Othr.ID
	response.AppHdr.BizMsgIdr = request.AppHdr.BizMsgIdr
	response.AppHdr.MsgDefIdr = "admn.002.001.01"
	response.AppHdr.CreDt = time.Now().Format("2006-01-02T15:04:05.000")

	response.Document.AdmnResp.GrpHdr.MsgId = time.Now().Format("2006-01-02T15:04:05Z07:00")[0:4] + time.Now().Format("2006-01-02T15:04:05Z07:00")[5:7] + time.Now().Format("2006-01-02T15:04:05Z07:00")[8:10] + ("INDOIDJA" + fmt.Sprint(rand.Intn(9)) + fmt.Sprint(rand.Intn(9)) + fmt.Sprint(rand.Intn(9)) + fmt.Sprint(rand.Intn(9)) +
		fmt.Sprint(rand.Intn(9)) + fmt.Sprint(rand.Intn(9)) + fmt.Sprint(rand.Intn(9)) + fmt.Sprint(rand.Intn(9)))
	response.Document.AdmnResp.GrpHdr.CreDtTm = time.Now().Format("2006-01-02T15:04:05.000")
	response.Document.AdmnResp.AdmnResponse.InstgAgt.FinInstnId.Othr.ID = "NOBUIDJA"
	response.Document.AdmnResp.AdmnResponse.OrgnlInstrId = request.Document.AdmnReq.AdmnTxInf.InstrId
	response.Document.AdmnResp.AdmnResponse.FnctnCd = request.Document.AdmnReq.AdmnTxInf.FnctnCd
	response.Document.AdmnResp.AdmnResponse.TxSts = "ACTC"

	payload, err := xml.Marshal(response)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(string(payload))

	return payload
}

// func main() {
// 	tes()
// }
