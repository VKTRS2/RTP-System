package main

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"time"
)

func tes() {
	s := `<?xml version="1.0" encoding="UTF-8"?>
	<ns:BusMsg xmlns:ns="urn:iso" xmlns:ns1="urn:iso:std:iso:20022:tech:xsd:head.001.001.01" xmlns:ns2="urn:iso:std:iso:20022:tech:xsd:admn.001.001.01" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="urn:iso ../../../../xsd/phase1/MainCIHub.xsd ">
	   <ns:AppHdr>
		  <ns1:Fr>
			 <ns1:FIId>
				<ns1:FinInstnId>
				   <ns1:Othr>
					  <ns1:Id>INDOIDJA</ns1:Id>
					  <!-- Sending system -->
				   </ns1:Othr>
				</ns1:FinInstnId>
			 </ns1:FIId>
		  </ns1:Fr>
		  <ns1:To>
			 <ns1:FIId>
				<ns1:FinInstnId>
				   <ns1:Othr>
					  <ns1:Id>FASTIDJA</ns1:Id>
					  <!-- Receiving system -->
				   </ns1:Othr>
				</ns1:FinInstnId>
			 </ns1:FIId>
		  </ns1:To>
		  <ns1:BizMsgIdr>20210409INDOIDJA000ORB12345678</ns1:BizMsgIdr>
		  <ns1:MsgDefIdr>admn.001.001.01</ns1:MsgDefIdr>
		  <ns1:CreDt>2021-04-09T12:00:00Z</ns1:CreDt>
	   </ns:AppHdr>
	   <ns:Document>
		  <ns:AdmnReq>
			 <ns2:GrpHdr>
				<ns2:MsgId>20210409INDOIDJA00012345678</ns2:MsgId>
				<ns2:CreDtTm>2021-04-09T17:44:27.076</ns2:CreDtTm>
			 </ns2:GrpHdr>
			 <ns2:AdmnTxInf>
				<ns2:FnctnCd>1003</ns2:FnctnCd>
				<ns2:InstrId>20210409INDOIDJA00012345678</ns2:InstrId>
				<ns2:InstgAgt>
				   <ns2:FinInstnId>
					  <ns2:Othr>
						 <ns2:Id>INDOIDJA</ns2:Id>
					  </ns2:Othr>
				   </ns2:FinInstnId>
				</ns2:InstgAgt>
			 </ns2:AdmnTxInf>
		  </ns:AdmnReq>
	   </ns:Document>
	</ns:BusMsg>`

	// s2 := `<?xml version="1.0" encoding="UTF-8"?>
	// <ns:BusMsg xmlns:ns="urn:iso" xmlns:ns1="urn:iso:std:iso:20022:tech:xsd:head.001.001.01" xmlns:ns2="urn:iso:std:iso:20022:tech:xsd:admn.002.001.01" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="urn:iso ../../../../xsd/phase1/MainCIHub.xsd ">
	//    <ns:AppHdr>
	// 	  <ns1:Fr>
	// 		 <ns1:FIId>
	// 			<ns1:FinInstnId>
	// 			   <ns1:Othr>
	// 				  <ns1:Id>FASTIDJA</ns1:Id>
	// 				  <!-- Sending system -->
	// 			   </ns1:Othr>
	// 			</ns1:FinInstnId>
	// 		 </ns1:FIId>
	// 	  </ns1:Fr>
	// 	  <ns1:To>
	// 		 <ns1:FIId>
	// 			<ns1:FinInstnId>
	// 			   <ns1:Othr>
	// 				  <ns1:Id>INDOIDJA</ns1:Id>
	// 				  <!-- Receiving system -->
	// 			   </ns1:Othr>
	// 			</ns1:FinInstnId>
	// 		 </ns1:FIId>
	// 	  </ns1:To>
	// 	  <ns1:BizMsgIdr>20210409FASTIDJA000ORB12345678</ns1:BizMsgIdr>
	// 	  <ns1:MsgDefIdr>admn.002.001.01</ns1:MsgDefIdr>
	// 	  <ns1:CreDt>2021-04-09T12:00:00Z</ns1:CreDt>
	//    </ns:AppHdr>
	//    <ns:Document>
	// 	  <ns:AdmnResp>
	// 		 <ns2:GrpHdr>
	// 			<ns2:MsgId>20210409INDOIDJA00012345678</ns2:MsgId>
	// 			<ns2:CreDtTm>2021-04-09T00:07:46.830</ns2:CreDtTm>
	// 		 </ns2:GrpHdr>
	// 		 <ns2:AdmnResponse>
	// 			<ns2:InstgAgt>
	// 			   <ns2:FinInstnId>
	// 				  <ns2:Othr>
	// 					 <ns2:Id>INDOIDJA</ns2:Id>
	// 				  </ns2:Othr>
	// 			   </ns2:FinInstnId>
	// 			</ns2:InstgAgt>
	// 			<ns2:OrgnlInstrId>20210409INDOIDJA00012345678</ns2:OrgnlInstrId>
	// 			<ns2:FnctnCd>1003</ns2:FnctnCd>
	// 			<ns2:TxSts>ACTC</ns2:TxSts>
	// 		 </ns2:AdmnResponse>
	// 	  </ns:AdmnResp>
	//    </ns:Document>
	// </ns:BusMsg>`

	request := Admn1{}
	err := xml.Unmarshal([]byte(s), &request)
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

	data, err := xml.Marshal(response)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(string(data))
}
