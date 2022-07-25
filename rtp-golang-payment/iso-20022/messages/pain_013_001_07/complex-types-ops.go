// Package pain_013_001_07
// Do not Edit. This stuff it's been automatically generated.
package pain_013_001_07

// IsValid checks if Cheque11 is valid
func (s Cheque11) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.ChqTp.IsValid(true)
	valid = valid && s.ChqNb.IsValid(true)
	valid = valid && (s.ChqFr == nil || (s.ChqFr != nil && s.ChqFr.IsValid(true)))

	valid = valid && (s.DlvryMtd == nil || (s.DlvryMtd != nil && s.DlvryMtd.IsValid(true)))

	valid = valid && (s.DlvrTo == nil || (s.DlvrTo != nil && s.DlvrTo.IsValid(true)))

	valid = valid && s.InstrPrty.IsValid(true)
	valid = valid && s.ChqMtrtyDt.IsValid(true)
	valid = valid && s.FrmsCd.IsValid(true)
	for j := 0; j < len(s.MemoFld); j++ {
		valid = valid && s.MemoFld[j].IsValid(true)
	}

	valid = valid && s.RgnlClrZone.IsValid(true)
	valid = valid && s.PrtLctn.IsValid(true)
	for j := 0; j < len(s.Sgntr); j++ {
		valid = valid && s.Sgntr[j].IsValid(true)
	}

	return valid
}

// IsValid checks if InstructionForCreditorAgent1 is valid
func (s InstructionForCreditorAgent1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.InstrInf.IsValid(true)

	return valid
}

// IsValid checks if ChequeDeliveryMethod1Choice is valid
func (s ChequeDeliveryMethod1Choice) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Cd.IsValid(true)
	valid = valid && s.Prtry.IsValid(true)

	return valid
}

// IsValid checks if PaymentInstruction31 is valid
func (s PaymentInstruction31) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.PmtInfId.IsValid(true)
	valid = valid && s.PmtMtd.IsValid(false)
	valid = valid && (s.PmtTpInf == nil || (s.PmtTpInf != nil && s.PmtTpInf.IsValid(true)))

	valid = valid && s.ReqdExctnDt.IsValid(false)
	valid = valid && (s.XpryDt == nil || (s.XpryDt != nil && s.XpryDt.IsValid(true)))

	valid = valid && (s.PmtCond == nil || (s.PmtCond != nil && s.PmtCond.IsValid(true)))

	valid = valid && s.Dbtr.IsValid(false)
	valid = valid && (s.DbtrAcct == nil || (s.DbtrAcct != nil && s.DbtrAcct.IsValid(true)))

	valid = valid && s.DbtrAgt.IsValid(false)
	valid = valid && (s.UltmtDbtr == nil || (s.UltmtDbtr != nil && s.UltmtDbtr.IsValid(true)))

	valid = valid && s.ChrgBr.IsValid(true)
	if len(s.CdtTrfTx) == 0 {
		valid = false
	}
	for j := 0; j < len(s.CdtTrfTx); j++ {
		valid = valid && s.CdtTrfTx[j].IsValid(false)
	}

	return valid
}

// IsValid checks if StructuredRegulatoryReporting3 is valid
func (s StructuredRegulatoryReporting3) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Tp.IsValid(true)
	valid = valid && s.Dt.IsValid(true)
	valid = valid && s.Ctry.IsValid(true)
	valid = valid && s.Cd.IsValid(true)
	valid = valid && (s.Amt == nil || (s.Amt != nil && s.Amt.IsValid(true)))

	for j := 0; j < len(s.Inf); j++ {
		valid = valid && s.Inf[j].IsValid(true)
	}

	return valid
}

// IsValid checks if CreditorPaymentActivationRequestV07 is valid
func (s CreditorPaymentActivationRequestV07) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.GrpHdr.IsValid(false)
	if len(s.PmtInf) == 0 {
		valid = false
	}
	for j := 0; j < len(s.PmtInf); j++ {
		valid = valid && s.PmtInf[j].IsValid(false)
	}

	for j := 0; j < len(s.SplmtryData); j++ {
		valid = valid && s.SplmtryData[j].IsValid(true)
	}

	return valid
}

// IsValid checks if GroupHeader78 is valid
func (s GroupHeader78) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.MsgId.IsValid(false)
	valid = valid && s.CreDtTm.IsValid(false)
	valid = valid && s.NbOfTxs.IsValid(false)
	valid = valid && s.CtrlSum.IsValid(true)
	valid = valid && s.InitgPty.IsValid(false)

	return valid
}

// IsValid checks if RemittanceLocationData1 is valid
func (s RemittanceLocationData1) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Mtd.IsValid(false)
	valid = valid && s.ElctrncAdr.IsValid(true)
	valid = valid && (s.PstlAdr == nil || (s.PstlAdr != nil && s.PstlAdr.IsValid(true)))

	return valid
}

// IsValid checks if RemittanceLocation7 is valid
func (s RemittanceLocation7) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.RmtId.IsValid(true)
	for j := 0; j < len(s.RmtLctnDtls); j++ {
		valid = valid && s.RmtLctnDtls[j].IsValid(true)
	}

	return valid
}

// IsValid checks if TaxInformation8 is valid
func (s TaxInformation8) IsValid(optional bool) bool {

	valid := true
	valid = valid && (s.Cdtr == nil || (s.Cdtr != nil && s.Cdtr.IsValid(true)))

	valid = valid && (s.Dbtr == nil || (s.Dbtr != nil && s.Dbtr.IsValid(true)))

	valid = valid && s.AdmstnZone.IsValid(true)
	valid = valid && s.RefNb.IsValid(true)
	valid = valid && s.Mtd.IsValid(true)
	valid = valid && (s.TtlTaxblBaseAmt == nil || (s.TtlTaxblBaseAmt != nil && s.TtlTaxblBaseAmt.IsValid(true)))

	valid = valid && (s.TtlTaxAmt == nil || (s.TtlTaxAmt != nil && s.TtlTaxAmt.IsValid(true)))

	valid = valid && s.Dt.IsValid(true)
	valid = valid && s.SeqNb.IsValid(true)
	for j := 0; j < len(s.Rcrd); j++ {
		valid = valid && s.Rcrd[j].IsValid(true)
	}

	return valid
}

// IsValid checks if PaymentIdentification6 is valid
func (s PaymentIdentification6) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.InstrId.IsValid(true)
	valid = valid && s.EndToEndId.IsValid(false)
	valid = valid && s.UETR.IsValid(true)

	return valid
}

// IsValid checks if NameAndAddress16 is valid
func (s NameAndAddress16) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Nm.IsValid(false)
	valid = valid && s.Adr.IsValid(false)

	return valid
}

// IsValid checks if RegulatoryReporting3 is valid
func (s RegulatoryReporting3) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.DbtCdtRptgInd.IsValid(true)
	valid = valid && (s.Authrty == nil || (s.Authrty != nil && s.Authrty.IsValid(true)))

	for j := 0; j < len(s.Dtls); j++ {
		valid = valid && s.Dtls[j].IsValid(true)
	}

	return valid
}

// IsValid checks if RegulatoryAuthority2 is valid
func (s RegulatoryAuthority2) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.Nm.IsValid(true)
	valid = valid && s.Ctry.IsValid(true)

	return valid
}

// IsValid checks if CreditTransferTransaction35 is valid
func (s CreditTransferTransaction35) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.PmtId.IsValid(false)
	valid = valid && (s.PmtTpInf == nil || (s.PmtTpInf != nil && s.PmtTpInf.IsValid(true)))

	valid = valid && (s.PmtCond == nil || (s.PmtCond != nil && s.PmtCond.IsValid(true)))

	valid = valid && s.Amt.IsValid(false)
	valid = valid && s.ChrgBr.IsValid(false)
	valid = valid && (s.ChqInstr == nil || (s.ChqInstr != nil && s.ChqInstr.IsValid(true)))

	valid = valid && (s.UltmtDbtr == nil || (s.UltmtDbtr != nil && s.UltmtDbtr.IsValid(true)))

	valid = valid && (s.IntrmyAgt1 == nil || (s.IntrmyAgt1 != nil && s.IntrmyAgt1.IsValid(true)))

	valid = valid && (s.IntrmyAgt2 == nil || (s.IntrmyAgt2 != nil && s.IntrmyAgt2.IsValid(true)))

	valid = valid && (s.IntrmyAgt3 == nil || (s.IntrmyAgt3 != nil && s.IntrmyAgt3.IsValid(true)))

	valid = valid && s.CdtrAgt.IsValid(false)
	valid = valid && s.Cdtr.IsValid(false)
	valid = valid && (s.CdtrAcct == nil || (s.CdtrAcct != nil && s.CdtrAcct.IsValid(true)))

	valid = valid && (s.UltmtCdtr == nil || (s.UltmtCdtr != nil && s.UltmtCdtr.IsValid(true)))

	for j := 0; j < len(s.InstrForCdtrAgt); j++ {
		valid = valid && s.InstrForCdtrAgt[j].IsValid(true)
	}

	valid = valid && (s.Purp == nil || (s.Purp != nil && s.Purp.IsValid(true)))

	for j := 0; j < len(s.RgltryRptg); j++ {
		valid = valid && s.RgltryRptg[j].IsValid(true)
	}

	valid = valid && (s.Tax == nil || (s.Tax != nil && s.Tax.IsValid(true)))

	for j := 0; j < len(s.RltdRmtInf); j++ {
		valid = valid && s.RltdRmtInf[j].IsValid(true)
	}

	valid = valid && (s.RmtInf == nil || (s.RmtInf != nil && s.RmtInf.IsValid(true)))

	for j := 0; j < len(s.NclsdFile); j++ {
		valid = valid && s.NclsdFile[j].IsValid(true)
	}

	for j := 0; j < len(s.SplmtryData); j++ {
		valid = valid && s.SplmtryData[j].IsValid(true)
	}

	return valid
}
