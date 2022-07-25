// Package pain_014_001_07
// Do not Edit. This stuff it's been automatically generated.
package pain_014_001_07

// IsValid checks if GroupHeader87 is valid
func (s GroupHeader87) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.MsgId.IsValid(false)
	valid = valid && s.CreDtTm.IsValid(false)
	valid = valid && s.InitgPty.IsValid(false)
	valid = valid && (s.DbtrAgt == nil || (s.DbtrAgt != nil && s.DbtrAgt.IsValid(true)))

	valid = valid && (s.CdtrAgt == nil || (s.CdtrAgt != nil && s.CdtrAgt.IsValid(true)))

	return valid
}

// IsValid checks if StatusReasonInformation12 is valid
func (s StatusReasonInformation12) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Orgtr == nil || (s.Orgtr != nil && s.Orgtr.IsValid(true)))

	valid = valid && (s.Rsn == nil || (s.Rsn != nil && s.Rsn.IsValid(true)))

	for j := 0; j < len(s.AddtlInf); j++ {
		valid = valid && s.AddtlInf[j].IsValid(true)
	}

	return valid
}

// IsValid checks if NumberOfTransactionsPerStatus5 is valid
func (s NumberOfTransactionsPerStatus5) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.DtldNbOfTxs.IsValid(false)
	valid = valid && s.DtldSts.IsValid(false)
	valid = valid && s.DtldCtrlSum.IsValid(true)

	return valid
}

// IsValid checks if OriginalPaymentInstruction31 is valid
func (s OriginalPaymentInstruction31) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.OrgnlPmtInfId.IsValid(false)
	valid = valid && s.OrgnlNbOfTxs.IsValid(true)
	valid = valid && s.OrgnlCtrlSum.IsValid(true)
	valid = valid && s.PmtInfSts.IsValid(true)
	for j := 0; j < len(s.StsRsnInf); j++ {
		valid = valid && s.StsRsnInf[j].IsValid(true)
	}

	for j := 0; j < len(s.NbOfTxsPerSts); j++ {
		valid = valid && s.NbOfTxsPerSts[j].IsValid(true)
	}

	for j := 0; j < len(s.TxInfAndSts); j++ {
		valid = valid && s.TxInfAndSts[j].IsValid(true)
	}

	return valid
}

// IsValid checks if OriginalGroupInformation30 is valid
func (s OriginalGroupInformation30) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.OrgnlMsgId.IsValid(false)
	valid = valid && s.OrgnlMsgNmId.IsValid(false)
	valid = valid && s.OrgnlCreDtTm.IsValid(true)
	valid = valid && s.OrgnlNbOfTxs.IsValid(true)
	valid = valid && s.OrgnlCtrlSum.IsValid(true)
	valid = valid && s.GrpSts.IsValid(true)
	for j := 0; j < len(s.StsRsnInf); j++ {
		valid = valid && s.StsRsnInf[j].IsValid(true)
	}

	for j := 0; j < len(s.NbOfTxsPerSts); j++ {
		valid = valid && s.NbOfTxsPerSts[j].IsValid(true)
	}

	return valid
}

// IsValid checks if StatusReason6Choice is valid
func (s StatusReason6Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if PaymentConditionStatus1 is valid
func (s PaymentConditionStatus1) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.AccptdAmt == nil || (s.AccptdAmt != nil && s.AccptdAmt.IsValid(true)))

	valid = valid && s.GrntedPmt.IsValid(false)
	valid = valid && s.EarlyPmt.IsValid(false)

	return valid
}

// IsValid checks if OriginalTransactionReference29 is valid
func (s OriginalTransactionReference29) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Amt == nil || (s.Amt != nil && s.Amt.IsValid(true)))

	valid = valid && (s.ReqdExctnDt == nil || (s.ReqdExctnDt != nil && s.ReqdExctnDt.IsValid(true)))

	valid = valid && (s.XpryDt == nil || (s.XpryDt != nil && s.XpryDt.IsValid(true)))

	valid = valid && (s.PmtCond == nil || (s.PmtCond != nil && s.PmtCond.IsValid(true)))

	valid = valid && (s.PmtTpInf == nil || (s.PmtTpInf != nil && s.PmtTpInf.IsValid(true)))

	valid = valid && s.PmtMtd.IsValid(true)
	valid = valid && (s.RmtInf == nil || (s.RmtInf != nil && s.RmtInf.IsValid(true)))

	for j := 0; j < len(s.NclsdFile); j++ {
		valid = valid && s.NclsdFile[j].IsValid(true)
	}

	valid = valid && (s.UltmtDbtr == nil || (s.UltmtDbtr != nil && s.UltmtDbtr.IsValid(true)))

	valid = valid && (s.Dbtr == nil || (s.Dbtr != nil && s.Dbtr.IsValid(true)))

	valid = valid && (s.DbtrAcct == nil || (s.DbtrAcct != nil && s.DbtrAcct.IsValid(true)))

	valid = valid && (s.DbtrAgt == nil || (s.DbtrAgt != nil && s.DbtrAgt.IsValid(true)))

	valid = valid && s.CdtrAgt.IsValid(false)
	valid = valid && s.Cdtr.IsValid(false)
	valid = valid && (s.CdtrAcct == nil || (s.CdtrAcct != nil && s.CdtrAcct.IsValid(true)))

	valid = valid && (s.UltmtCdtr == nil || (s.UltmtCdtr != nil && s.UltmtCdtr.IsValid(true)))

	return valid
}

// IsValid checks if PaymentTransaction104 is valid
func (s PaymentTransaction104) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.StsId.IsValid(true)
	valid = valid && s.OrgnlInstrId.IsValid(true)
	valid = valid && s.OrgnlEndToEndId.IsValid(true)
	valid = valid && s.OrgnlUETR.IsValid(true)
	valid = valid && s.TxSts.IsValid(true)
	for j := 0; j < len(s.StsRsnInf); j++ {
		valid = valid && s.StsRsnInf[j].IsValid(true)
	}

	valid = valid && (s.PmtCondSts == nil || (s.PmtCondSts != nil && s.PmtCondSts.IsValid(true)))

	for j := 0; j < len(s.ChrgsInf); j++ {
		valid = valid && s.ChrgsInf[j].IsValid(true)
	}

	valid = valid && s.DbtrDcsnDtTm.IsValid(true)
	valid = valid && s.AccptncDtTm.IsValid(true)
	valid = valid && s.AcctSvcrRef.IsValid(true)
	valid = valid && s.ClrSysRef.IsValid(true)
	valid = valid && (s.OrgnlTxRef == nil || (s.OrgnlTxRef != nil && s.OrgnlTxRef.IsValid(true)))

	for j := 0; j < len(s.NclsdFile); j++ {
		valid = valid && s.NclsdFile[j].IsValid(true)
	}

	for j := 0; j < len(s.SplmtryData); j++ {
		valid = valid && s.SplmtryData[j].IsValid(true)
	}

	return valid
}

// IsValid checks if CreditorPaymentActivationRequestStatusReportV07 is valid
func (s CreditorPaymentActivationRequestStatusReportV07) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.GrpHdr.IsValid(false)
	valid = valid && s.OrgnlGrpInfAndSts.IsValid(false)
	for j := 0; j < len(s.OrgnlPmtInfAndSts); j++ {
		valid = valid && s.OrgnlPmtInfAndSts[j].IsValid(true)
	}

	for j := 0; j < len(s.SplmtryData); j++ {
		valid = valid && s.SplmtryData[j].IsValid(true)
	}

	return valid
}
