// Package pacs_028_001_03
// Do not Edit. This stuff it's been automatically generated.
package pacs_028_001_03

// IsValid checks if OriginalGroupInformation27 is valid
func (s OriginalGroupInformation27) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.OrgnlMsgId.IsValid(false)
	valid = valid && s.OrgnlMsgNmId.IsValid(false)
	valid = valid && s.OrgnlCreDtTm.IsValid(true)
	valid = valid && s.OrgnlNbOfTxs.IsValid(true)
	valid = valid && s.OrgnlCtrlSum.IsValid(true)

	return valid
}

// IsValid checks if PaymentTransaction113 is valid
func (s PaymentTransaction113) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.StsReqId.IsValid(true)
	valid = valid && (s.OrgnlGrpInf == nil || (s.OrgnlGrpInf != nil && s.OrgnlGrpInf.IsValid(true)))

	valid = valid && s.OrgnlInstrId.IsValid(true)
	valid = valid && s.OrgnlEndToEndId.IsValid(true)
	valid = valid && s.OrgnlTxId.IsValid(true)
	valid = valid && s.OrgnlUETR.IsValid(true)
	valid = valid && s.AccptncDtTm.IsValid(true)
	valid = valid && s.ClrSysRef.IsValid(true)
	valid = valid && (s.InstgAgt == nil || (s.InstgAgt != nil && s.InstgAgt.IsValid(true)))

	valid = valid && (s.InstdAgt == nil || (s.InstdAgt != nil && s.InstdAgt.IsValid(true)))

	valid = valid && (s.OrgnlTxRef == nil || (s.OrgnlTxRef != nil && s.OrgnlTxRef.IsValid(true)))

	for j := 0; j < len(s.SplmtryData); j++ {
		valid = valid && s.SplmtryData[j].IsValid(true)
	}

	return valid
}

// IsValid checks if GroupHeader91 is valid
func (s GroupHeader91) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.MsgId.IsValid(false)
	valid = valid && s.CreDtTm.IsValid(false)
	valid = valid && (s.InstgAgt == nil || (s.InstgAgt != nil && s.InstgAgt.IsValid(true)))

	valid = valid && (s.InstdAgt == nil || (s.InstdAgt != nil && s.InstdAgt.IsValid(true)))

	return valid
}

// IsValid checks if FIToFIPaymentStatusRequestV03 is valid
func (s FIToFIPaymentStatusRequestV03) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.GrpHdr.IsValid(false)
	for j := 0; j < len(s.OrgnlGrpInf); j++ {
		valid = valid && s.OrgnlGrpInf[j].IsValid(true)
	}

	for j := 0; j < len(s.TxInf); j++ {
		valid = valid && s.TxInf[j].IsValid(true)
	}

	for j := 0; j < len(s.SplmtryData); j++ {
		valid = valid && s.SplmtryData[j].IsValid(true)
	}

	return valid
}
