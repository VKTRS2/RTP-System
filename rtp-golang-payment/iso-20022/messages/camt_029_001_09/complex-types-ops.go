// Package camt_029_001_09
// Do not Edit. This stuff it's been automatically generated.
package camt_029_001_09

// IsValid checks if PaymentTransaction103 is valid
func (s PaymentTransaction103) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.CxlStsId.IsValid(true)
	valid = valid && (s.RslvdCase == nil || (s.RslvdCase != nil && s.RslvdCase.IsValid(true)))

	valid = valid && s.OrgnlInstrId.IsValid(true)
	valid = valid && s.OrgnlEndToEndId.IsValid(true)
	valid = valid && s.UETR.IsValid(true)
	valid = valid && s.TxCxlSts.IsValid(true)
	for j := 0; j < len(s.CxlStsRsnInf); j++ {
		valid = valid && s.CxlStsRsnInf[j].IsValid(true)
	}

	valid = valid && (s.OrgnlInstdAmt == nil || (s.OrgnlInstdAmt != nil && s.OrgnlInstdAmt.IsValid(true)))

	valid = valid && (s.OrgnlReqdExctnDt == nil || (s.OrgnlReqdExctnDt != nil && s.OrgnlReqdExctnDt.IsValid(true)))

	valid = valid && s.OrgnlReqdColltnDt.IsValid(true)
	valid = valid && (s.OrgnlTxRef == nil || (s.OrgnlTxRef != nil && s.OrgnlTxRef.IsValid(true)))

	return valid
}

// IsValid checks if ResolutionOfInvestigationV09 is valid
func (s ResolutionOfInvestigationV09) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Assgnmt.IsValid(false)
	valid = valid && (s.RslvdCase == nil || (s.RslvdCase != nil && s.RslvdCase.IsValid(true)))

	valid = valid && s.Sts.IsValid(false)
	for j := 0; j < len(s.CxlDtls); j++ {
		valid = valid && s.CxlDtls[j].IsValid(true)
	}

	valid = valid && (s.ModDtls == nil || (s.ModDtls != nil && s.ModDtls.IsValid(true)))

	valid = valid && (s.ClmNonRctDtls == nil || (s.ClmNonRctDtls != nil && s.ClmNonRctDtls.IsValid(true)))

	valid = valid && (s.StmtDtls == nil || (s.StmtDtls != nil && s.StmtDtls.IsValid(true)))

	valid = valid && (s.CrrctnTx == nil || (s.CrrctnTx != nil && s.CrrctnTx.IsValid(true)))

	valid = valid && (s.RsltnRltdInf == nil || (s.RsltnRltdInf != nil && s.RsltnRltdInf.IsValid(true)))

	for j := 0; j < len(s.SplmtryData); j++ {
		valid = valid && s.SplmtryData[j].IsValid(true)
	}

	return valid
}

// IsValid checks if Compensation2 is valid
func (s Compensation2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Amt.IsValid(false)
	valid = valid && s.DbtrAgt.IsValid(false)
	valid = valid && s.CdtrAgt.IsValid(false)
	valid = valid && s.Rsn.IsValid(false)

	return valid
}

// IsValid checks if NumberOfTransactionsPerStatus1 is valid
func (s NumberOfTransactionsPerStatus1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.DtldNbOfTxs.IsValid(false)
	valid = valid && s.DtldSts.IsValid(false)
	valid = valid && s.DtldCtrlSum.IsValid(true)

	return valid
}

// IsValid checks if ClaimNonReceipt2Choice is valid
func (s ClaimNonReceipt2Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Accptd == nil || (s.Accptd != nil && s.Accptd.IsValid(true)))

	valid = valid && (s.Rjctd == nil || (s.Rjctd != nil && s.Rjctd.IsValid(true)))

	return valid
}

// IsValid checks if CorrectiveGroupInformation1 is valid
func (s CorrectiveGroupInformation1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.MsgId.IsValid(false)
	valid = valid && s.MsgNmId.IsValid(false)
	valid = valid && s.CreDtTm.IsValid(true)

	return valid
}

// IsValid checks if CompensationReason1Choice is valid
func (s CompensationReason1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if CancellationStatusReason4 is valid
func (s CancellationStatusReason4) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Orgtr == nil || (s.Orgtr != nil && s.Orgtr.IsValid(true)))

	valid = valid && (s.Rsn == nil || (s.Rsn != nil && s.Rsn.IsValid(true)))

	for j := 0; j < len(s.AddtlInf); j++ {
		valid = valid && s.AddtlInf[j].IsValid(true)
	}

	return valid
}

// IsValid checks if ClaimNonReceipt2 is valid
func (s ClaimNonReceipt2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.DtPrcd.IsValid(false)
	valid = valid && (s.OrgnlNxtAgt == nil || (s.OrgnlNxtAgt != nil && s.OrgnlNxtAgt.IsValid(true)))

	return valid
}

// IsValid checks if GenericIdentification3 is valid
func (s GenericIdentification3) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(false)
	valid = valid && s.Issr.IsValid(true)

	return valid
}

// IsValid checks if CorrectivePaymentInitiation4 is valid
func (s CorrectivePaymentInitiation4) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.GrpHdr == nil || (s.GrpHdr != nil && s.GrpHdr.IsValid(true)))

	valid = valid && s.PmtInfId.IsValid(true)
	valid = valid && s.InstrId.IsValid(true)
	valid = valid && s.EndToEndId.IsValid(true)
	valid = valid && s.UETR.IsValid(true)
	valid = valid && s.InstdAmt.IsValid(false)
	valid = valid && (s.ReqdExctnDt == nil || (s.ReqdExctnDt != nil && s.ReqdExctnDt.IsValid(true)))

	valid = valid && s.ReqdColltnDt.IsValid(true)

	return valid
}

// IsValid checks if CorrectiveInterbankTransaction2 is valid
func (s CorrectiveInterbankTransaction2) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.GrpHdr == nil || (s.GrpHdr != nil && s.GrpHdr.IsValid(true)))

	valid = valid && s.InstrId.IsValid(true)
	valid = valid && s.EndToEndId.IsValid(true)
	valid = valid && s.TxId.IsValid(true)
	valid = valid && s.UETR.IsValid(true)
	valid = valid && s.IntrBkSttlmAmt.IsValid(false)
	valid = valid && s.IntrBkSttlmDt.IsValid(false)

	return valid
}

// IsValid checks if ChargesRecord3 is valid
func (s ChargesRecord3) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Amt.IsValid(false)
	valid = valid && s.CdtDbtInd.IsValid(true)
	valid = valid && s.ChrgInclInd.IsValid(true)
	valid = valid && (s.Tp == nil || (s.Tp != nil && s.Tp.IsValid(true)))

	valid = valid && s.Rate.IsValid(true)
	valid = valid && s.Br.IsValid(true)
	valid = valid && (s.Agt == nil || (s.Agt != nil && s.Agt.IsValid(true)))

	valid = valid && (s.Tax == nil || (s.Tax != nil && s.Tax.IsValid(true)))

	return valid
}

// IsValid checks if ResolutionData1 is valid
func (s ResolutionData1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.EndToEndId.IsValid(true)
	valid = valid && s.TxId.IsValid(true)
	valid = valid && s.UETR.IsValid(true)
	valid = valid && (s.IntrBkSttlmAmt == nil || (s.IntrBkSttlmAmt != nil && s.IntrBkSttlmAmt.IsValid(true)))

	valid = valid && s.IntrBkSttlmDt.IsValid(true)
	valid = valid && s.ClrChanl.IsValid(true)
	valid = valid && (s.Compstn == nil || (s.Compstn != nil && s.Compstn.IsValid(true)))

	for j := 0; j < len(s.Chrgs); j++ {
		valid = valid && s.Chrgs[j].IsValid(true)
	}

	return valid
}

// IsValid checks if OriginalGroupHeader14 is valid
func (s OriginalGroupHeader14) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.OrgnlGrpCxlId.IsValid(true)
	valid = valid && (s.RslvdCase == nil || (s.RslvdCase != nil && s.RslvdCase.IsValid(true)))

	valid = valid && s.OrgnlMsgId.IsValid(false)
	valid = valid && s.OrgnlMsgNmId.IsValid(false)
	valid = valid && s.OrgnlCreDtTm.IsValid(true)
	valid = valid && s.OrgnlNbOfTxs.IsValid(true)
	valid = valid && s.OrgnlCtrlSum.IsValid(true)
	valid = valid && s.GrpCxlSts.IsValid(true)
	for j := 0; j < len(s.CxlStsRsnInf); j++ {
		valid = valid && s.CxlStsRsnInf[j].IsValid(true)
	}

	for j := 0; j < len(s.NbOfTxsPerCxlSts); j++ {
		valid = valid && s.NbOfTxsPerCxlSts[j].IsValid(true)
	}

	return valid
}

// IsValid checks if Charges6 is valid
func (s Charges6) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.TtlChrgsAndTaxAmt == nil || (s.TtlChrgsAndTaxAmt != nil && s.TtlChrgsAndTaxAmt.IsValid(true)))

	for j := 0; j < len(s.Rcrd); j++ {
		valid = valid && s.Rcrd[j].IsValid(true)
	}

	return valid
}

// IsValid checks if StatementResolutionEntry4 is valid
func (s StatementResolutionEntry4) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.OrgnlGrpInf == nil || (s.OrgnlGrpInf != nil && s.OrgnlGrpInf.IsValid(true)))

	valid = valid && s.OrgnlStmtId.IsValid(true)
	valid = valid && s.UETR.IsValid(true)
	valid = valid && s.AcctSvcrRef.IsValid(true)
	valid = valid && (s.CrrctdAmt == nil || (s.CrrctdAmt != nil && s.CrrctdAmt.IsValid(true)))

	for j := 0; j < len(s.Chrgs); j++ {
		valid = valid && s.Chrgs[j].IsValid(true)
	}

	valid = valid && (s.Purp == nil || (s.Purp != nil && s.Purp.IsValid(true)))

	return valid
}

// IsValid checks if ClaimNonReceiptRejectReason1Choice is valid
func (s ClaimNonReceiptRejectReason1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if ModificationStatusReason2 is valid
func (s ModificationStatusReason2) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Orgtr == nil || (s.Orgtr != nil && s.Orgtr.IsValid(true)))

	valid = valid && (s.Rsn == nil || (s.Rsn != nil && s.Rsn.IsValid(true)))

	for j := 0; j < len(s.AddtlInf); j++ {
		valid = valid && s.AddtlInf[j].IsValid(true)
	}

	return valid
}

// IsValid checks if PaymentTransaction107 is valid
func (s PaymentTransaction107) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.ModStsId.IsValid(true)
	valid = valid && (s.RslvdCase == nil || (s.RslvdCase != nil && s.RslvdCase.IsValid(true)))

	valid = valid && s.OrgnlGrpInf.IsValid(false)
	valid = valid && s.OrgnlPmtInfId.IsValid(true)
	valid = valid && s.OrgnlInstrId.IsValid(true)
	valid = valid && s.OrgnlEndToEndId.IsValid(true)
	valid = valid && s.OrgnlTxId.IsValid(true)
	valid = valid && s.OrgnlClrSysRef.IsValid(true)
	valid = valid && s.OrgnlUETR.IsValid(true)
	for j := 0; j < len(s.ModStsRsnInf); j++ {
		valid = valid && s.ModStsRsnInf[j].IsValid(true)
	}

	valid = valid && (s.RsltnRltdInf == nil || (s.RsltnRltdInf != nil && s.RsltnRltdInf.IsValid(true)))

	valid = valid && (s.OrgnlIntrBkSttlmAmt == nil || (s.OrgnlIntrBkSttlmAmt != nil && s.OrgnlIntrBkSttlmAmt.IsValid(true)))

	valid = valid && s.OrgnlIntrBkSttlmDt.IsValid(true)
	valid = valid && (s.Assgnr == nil || (s.Assgnr != nil && s.Assgnr.IsValid(true)))

	valid = valid && (s.Assgne == nil || (s.Assgne != nil && s.Assgne.IsValid(true)))

	valid = valid && (s.OrgnlTxRef == nil || (s.OrgnlTxRef != nil && s.OrgnlTxRef.IsValid(true)))

	return valid
}

// IsValid checks if ModificationStatusReason1Choice is valid
func (s ModificationStatusReason1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if OriginalPaymentInstruction30 is valid
func (s OriginalPaymentInstruction30) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.OrgnlPmtInfCxlId.IsValid(true)
	valid = valid && (s.RslvdCase == nil || (s.RslvdCase != nil && s.RslvdCase.IsValid(true)))

	valid = valid && s.OrgnlPmtInfId.IsValid(false)
	valid = valid && (s.OrgnlGrpInf == nil || (s.OrgnlGrpInf != nil && s.OrgnlGrpInf.IsValid(true)))

	valid = valid && s.OrgnlNbOfTxs.IsValid(true)
	valid = valid && s.OrgnlCtrlSum.IsValid(true)
	valid = valid && s.PmtInfCxlSts.IsValid(true)
	for j := 0; j < len(s.CxlStsRsnInf); j++ {
		valid = valid && s.CxlStsRsnInf[j].IsValid(true)
	}

	for j := 0; j < len(s.NbOfTxsPerCxlSts); j++ {
		valid = valid && s.NbOfTxsPerCxlSts[j].IsValid(true)
	}

	for j := 0; j < len(s.TxInfAndSts); j++ {
		valid = valid && s.TxInfAndSts[j].IsValid(true)
	}

	return valid
}

// IsValid checks if ChargeType3Choice is valid
func (s ChargeType3Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && (s.Prtry == nil || (s.Prtry != nil && s.Prtry.IsValid(true)))

	return valid
}

// IsValid checks if CancellationStatusReason3Choice is valid
func (s CancellationStatusReason3Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if NumberOfCancellationsPerStatus1 is valid
func (s NumberOfCancellationsPerStatus1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.DtldNbOfTxs.IsValid(false)
	valid = valid && s.DtldSts.IsValid(false)
	valid = valid && s.DtldCtrlSum.IsValid(true)

	return valid
}

// IsValid checks if CorrectiveTransaction4Choice is valid
func (s CorrectiveTransaction4Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Initn == nil || (s.Initn != nil && s.Initn.IsValid(true)))

	valid = valid && (s.IntrBk == nil || (s.IntrBk != nil && s.IntrBk.IsValid(true)))

	return valid
}

// IsValid checks if UnderlyingTransaction22 is valid
func (s UnderlyingTransaction22) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.OrgnlGrpInfAndSts == nil || (s.OrgnlGrpInfAndSts != nil && s.OrgnlGrpInfAndSts.IsValid(true)))

	for j := 0; j < len(s.OrgnlPmtInfAndSts); j++ {
		valid = valid && s.OrgnlPmtInfAndSts[j].IsValid(true)
	}

	for j := 0; j < len(s.TxInfAndSts); j++ {
		valid = valid && s.TxInfAndSts[j].IsValid(true)
	}

	return valid
}

// IsValid checks if InvestigationStatus5Choice is valid
func (s InvestigationStatus5Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Conf.IsValid(true)
	for j := 0; j < len(s.RjctdMod); j++ {
		valid = valid && s.RjctdMod[j].IsValid(true)
	}

	valid = valid && (s.DplctOf == nil || (s.DplctOf != nil && s.DplctOf.IsValid(true)))

	valid = valid && s.AssgnmtCxlConf.IsValid(true)

	return valid
}

// IsValid checks if PaymentTransaction102 is valid
func (s PaymentTransaction102) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.CxlStsId.IsValid(true)
	valid = valid && (s.RslvdCase == nil || (s.RslvdCase != nil && s.RslvdCase.IsValid(true)))

	valid = valid && (s.OrgnlGrpInf == nil || (s.OrgnlGrpInf != nil && s.OrgnlGrpInf.IsValid(true)))

	valid = valid && s.OrgnlInstrId.IsValid(true)
	valid = valid && s.OrgnlEndToEndId.IsValid(true)
	valid = valid && s.OrgnlTxId.IsValid(true)
	valid = valid && s.OrgnlClrSysRef.IsValid(true)
	valid = valid && s.OrgnlUETR.IsValid(true)
	valid = valid && s.TxCxlSts.IsValid(true)
	for j := 0; j < len(s.CxlStsRsnInf); j++ {
		valid = valid && s.CxlStsRsnInf[j].IsValid(true)
	}

	valid = valid && (s.RsltnRltdInf == nil || (s.RsltnRltdInf != nil && s.RsltnRltdInf.IsValid(true)))

	valid = valid && (s.OrgnlIntrBkSttlmAmt == nil || (s.OrgnlIntrBkSttlmAmt != nil && s.OrgnlIntrBkSttlmAmt.IsValid(true)))

	valid = valid && s.OrgnlIntrBkSttlmDt.IsValid(true)
	valid = valid && (s.Assgnr == nil || (s.Assgnr != nil && s.Assgnr.IsValid(true)))

	valid = valid && (s.Assgne == nil || (s.Assgne != nil && s.Assgne.IsValid(true)))

	valid = valid && (s.OrgnlTxRef == nil || (s.OrgnlTxRef != nil && s.OrgnlTxRef.IsValid(true)))

	return valid
}

// IsValid checks if TaxCharges2 is valid
func (s TaxCharges2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Id.IsValid(true)
	valid = valid && s.Rate.IsValid(true)
	valid = valid && (s.Amt == nil || (s.Amt != nil && s.Amt.IsValid(true)))

	return valid
}
