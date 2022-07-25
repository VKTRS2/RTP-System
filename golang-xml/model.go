package main

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022/head"
)

type BusMsg struct {
	AppHdr *head.BusinessApplicationHeaderV01 `xml:"AppHdr"`
}

type Message interface {
	String() (result string, ok bool)
}

type Document struct {
	Namespaces map[string]string
	BusMsg     BusMsg `xml:"BusMsg"`
}

func (a *Document) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	a.Namespaces = map[string]string{}
	for _, attr := range start.Attr {
		if attr.Name.Space == "xmlns" {
			a.Namespaces[attr.Name.Local] = attr.Value
		}
	}

	// Go on with unmarshalling.
	type app Document
	aa := (*app)(a)
	return d.DecodeElement(aa, &start)
}

type Admn1 struct {
	XMLName        xml.Name `xml:"BusMsg"`
	Text           string   `xml:",chardata"`
	Ns             string   `xml:"ns,attr"`
	Ns1            string   `xml:"ns1,attr"`
	Ns2            string   `xml:"ns2,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	AppHdr         struct {
		Text string `xml:",chardata"`
		Fr   struct {
			Text string `xml:",chardata"`
			FIId struct {
				Text       string `xml:",chardata"`
				FinInstnId struct {
					Text string `xml:",chardata"`
					Othr struct {
						Text string `xml:",chardata"`
						ID   string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"FIId"`
		} `xml:"Fr"`
		To struct {
			Text string `xml:",chardata"`
			FIId struct {
				Text       string `xml:",chardata"`
				FinInstnId struct {
					Text string `xml:",chardata"`
					Othr struct {
						Text string `xml:",chardata"`
						ID   string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"FIId"`
		} `xml:"To"`
		BizMsgIdr string `xml:"BizMsgIdr"`
		MsgDefIdr string `xml:"MsgDefIdr"`
		CreDt     string `xml:"CreDt"`
	} `xml:"AppHdr"`
	Document struct {
		Text    string `xml:",chardata"`
		AdmnReq struct {
			Text   string `xml:",chardata"`
			GrpHdr struct {
				Text    string `xml:",chardata"`
				MsgId   string `xml:"MsgId"`
				CreDtTm string `xml:"CreDtTm"`
			} `xml:"GrpHdr"`
			AdmnTxInf struct {
				Text     string `xml:",chardata"`
				FnctnCd  string `xml:"FnctnCd"`
				InstrId  string `xml:"InstrId"`
				InstgAgt struct {
					Text       string `xml:",chardata"`
					FinInstnId struct {
						Text string `xml:",chardata"`
						Othr struct {
							Text string `xml:",chardata"`
							ID   string `xml:"Id"`
						} `xml:"Othr"`
					} `xml:"FinInstnId"`
				} `xml:"InstgAgt"`
			} `xml:"AdmnTxInf"`
		} `xml:"AdmnReq"`
	} `xml:"Document"`
}

type Admn2 struct {
	XMLName        xml.Name `xml:"ns:BusMsg"`
	Text           string   `xml:",chardata"`
	Ns             string   `xml:"xmlns:ns,attr"`
	Ns1            string   `xml:"xmlns:ns1,attr"`
	Ns2            string   `xml:"xmlns:ns2,attr"`
	Xsi            string   `xml:"xmlns:xsi,attr"`
	SchemaLocation string   `xml:"xsi:schemaLocation,attr"`
	AppHdr         struct {
		Text string `xml:",chardata"`
		Fr   struct {
			Text string `xml:",chardata"`
			FIId struct {
				Text       string `xml:",chardata"`
				FinInstnId struct {
					Text string `xml:",chardata"`
					Othr struct {
						Text string `xml:",chardata"`
						ID   string `xml:"ns1:Id"`
					} `xml:"ns1:Othr"`
				} `xml:"ns1:FinInstnId"`
			} `xml:"ns1:FIId"`
		} `xml:"ns1:Fr"`
		To struct {
			Text string `xml:",chardata"`
			FIId struct {
				Text       string `xml:",chardata"`
				FinInstnId struct {
					Text string `xml:",chardata"`
					Othr struct {
						Text string `xml:",chardata"`
						ID   string `xml:"ns1:Id"`
					} `xml:"ns1:Othr"`
				} `xml:"ns1:FinInstnId"`
			} `xml:"ns1:FIId"`
		} `xml:"ns1:To"`
		BizMsgIdr string `xml:"ns1:BizMsgIdr"`
		MsgDefIdr string `xml:"ns1:MsgDefIdr"`
		CreDt     string `xml:"ns1:CreDt"`
	} `xml:"ns:AppHdr"`
	Document struct {
		Text     string `xml:",chardata"`
		AdmnResp struct {
			Text   string `xml:",chardata"`
			GrpHdr struct {
				Text    string `xml:",chardata"`
				MsgId   string `xml:"ns2:MsgId"`
				CreDtTm string `xml:"ns2:CreDtTm"`
			} `xml:"ns2:GrpHdr"`
			AdmnResponse struct {
				Text     string `xml:",chardata"`
				InstgAgt struct {
					Text       string `xml:",chardata"`
					FinInstnId struct {
						Text string `xml:",chardata"`
						Othr struct {
							Text string `xml:",chardata"`
							ID   string `xml:"ns2:Id"`
						} `xml:"ns2:Othr"`
					} `xml:"ns2:FinInstnId"`
				} `xml:"ns2:InstgAgt"`
				OrgnlInstrId string `xml:"ns2:OrgnlInstrId"`
				FnctnCd      string `xml:"ns2:FnctnCd"`
				TxSts        string `xml:"ns2:TxSts"`
			} `xml:"ns2:AdmnResponse"`
		} `xml:"ns:AdmnResp"`
	} `xml:"ns:Document"`
}
