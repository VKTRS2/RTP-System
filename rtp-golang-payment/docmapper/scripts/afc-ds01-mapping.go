package main

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/docmapper"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/pain_013_001_07"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
)

func Gen_AFC_DS01_Pain013_001_07(fn string) error {

	// In the normal case use the NewMapppingClass and avoid the WithFuncMap(nil) call that sets the builtins...
	dm, err := docmapper.NewMapperClass("pain.013.001.07", "AFC-DS01-pain.013.001.07")
	if err != nil {
		return err
	}

	dm.Rules = []docmapper.MappingRule{
		{
			Name:        "GrpHdr_MsgId",
			SourceValue: `"{$.payee-company-name}-DS01-" + objId()`,
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_MsgId,
		},
		{
			Name:        "GrpHdr_CreDtTm",
			SourceValue: "{$.rtp-timestamp}",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_CreDtTm,
		},
		{
			Name:        "GrpHdr_NbOfTxs",
			SourceValue: "1",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_NbOfTxs,
		},
		{
			Name:        "GrpHdr_InitgPty_Nm",
			SourceValue: "{$.payee-company-name} {$.payee-name}",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm,
		},
		{
			Name:        "GrpHdr_InitgPty_Id_OrgId_Othr_Id",
			SourceValue: "{$.payee-id}",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_Id,
		},
		{
			Name:        "GrpHdr_InitgPty_Id_OrgId_Othr_Issr",
			SourceValue: "{$.payee-company-name} {$.payee-name}",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_Issr,
		},
		{
			Name:        "PmtInf_PmtInfId",
			SourceValue: `"F/{$.payee-e2e-rtp-ref}/P/" + rtpId()`,
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_PmtInfId,
		},
		{
			Name:        "PmtInf_PmtMtd",
			SourceValue: "TRF",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_PmtMtd,
		},
		{
			Name:        "PmtInf_PmtTpInf_SvcLvl_Cd",
			SourceValue: "SRTP",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_SvcLvl_Cd,
		},
		{
			Name:        "PmtInf_PmtTpInf_LclInstrm_Prtry",
			SourceValue: "{$.pmt-instrument}",

			TargetPath: pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_LclInstrm_Prtry,
		},
		{
			Name:        "PmtInf_PmtTpInf_CtgyPurp_Cd",
			SourceValue: "OTHR",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_CtgyPurp_Cd,
		},
		{
			Name:        "PmtInf_ReqdExctnDt_Dt",
			SourceValue: "{$.pmt-req-exec-date}",

			TargetPath: pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_ReqdExctnDt_Dt,
		},
		{
			Name:        "PmtInf_XpryDt_Dt",
			SourceValue: "{$.rtp-expiry-date}",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_XpryDt_Dt,
		},
		{
			Name:        "PmtInf_Dbtr_Nm",
			SourceValue: "No-Name",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Nm,
		},
		{
			Name:        "PmtInf_Dbtr_PstlAdr_AdrLine",
			SourceValue: "No-AdrLine",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrLine,
		},
		{
			Name:        "PmtInf_Dbtr_Id_PrvtId_Othr_Id",
			SourceValue: "{$.payer-id}",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Id,
		},
		{
			Name:        "PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd",
			SourceValue: "POID",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd,
		},
		/*		{
				Name:        "PmtInf_DbtrAcct_Id_IBAN",
				SourceValue: "No-IBAN",

				TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_IBAN,
			},*/
		{
			Name:        "PmtInf_DbtrAgt_FinInstnId_BICFI",
			SourceValue: "{$.payer-psp-id}",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_BICFI,
		},
		{
			Name:        "PmtInf_CdtTrfTx_PmtId_EndToEndId",
			SourceValue: "{$.payee-e2e-rtp-ref}",

			TargetPath: pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId_EndToEndId,
		},
		{
			Name:        "PmtInf_CdtTrfTx_PmtId_InstrId",
			SourceValue: "{$.payee-e2e-rtp-ref}",

			TargetPath: pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId_InstrId,
		},
		{
			Name:        "PmtInf_CdtTrfTx_PmtTpInf_SvcLvl_Cd",
			SourceValue: "SRTP",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_SvcLvl_Cd,
		},
		{
			Name:        "PmtInf_CdtTrfTx_PmtTpInf_LclInstrm_Prtry",
			SourceValue: "{$.pmt-instrument}",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_LclInstrm_Prtry,
		},
		{
			Name:        "PmtInf_CdtTrfTx_PmtTpInf_CtgyPurp_Cd",
			SourceValue: "OTHR",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_CtgyPurp_Cd,
		},
		{
			Name:        "PmtInf_CdtTrfTx_PmtCond_AmtModAllwd",
			SourceValue: "false",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_AmtModAllwd,
		},
		{
			Name:        "PmtInf_CdtTrfTx_PmtCond_EarlyPmtAllwd",
			SourceValue: "false",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_EarlyPmtAllwd,
		},
		{
			Name:        "PmtInf_CdtTrfTx_PmtCond_GrntedPmtReqd",
			SourceValue: "false",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_GrntedPmtReqd,
		},
		{
			Name:        "PmtInf_CdtTrfTx_Amt_InstdAmt_Ccy",
			SourceValue: "EUR",

			TargetPath: pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Ccy,
		},
		{
			Name:        "PmtInf_CdtTrfTx_Amt_InstdAmt_Value",
			SourceValue: "{$.amount}",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Value,
		},
		{
			Name:        "PmtInf_CdtTrfTx_ChrgBr",
			SourceValue: "SLEV",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChrgBr,
		},
		{
			Name:        "PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_BICFI",
			SourceValue: "{$.payee-psp-bic}",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_BICFI,
		},
		{
			Name:        "PmtInf_CdtTrfTx_Cdtr_Nm",
			SourceValue: "{$.payee-company-name} {$.payee-name}",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Nm,
		},
		{
			Name:        "PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrLine",
			SourceValue: "NO-AdrLine",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrLine,
		},
		// First line in array
		{
			Name:        "PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_Id",
			SourceValue: "{$.payee-id}",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_Id,
		},
		{
			Name:        "PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_SchmeNm_Cd",
			SourceValue: "BOID",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_SchmeNm_Cd,
		},
		// Second line in array
		{
			Name:        "PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_Id",
			SourceValue: "{$.payee-company-name}",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_Id,
		},
		{
			Name:        "PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_SchmeNm_Cd",
			SourceValue: "BCID",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_SchmeNm_Cd,
		},
		{
			Name:        "PmtInf_CdtTrfTx_CdtrAcct_Id_IBAN",
			SourceValue: "{$.payee-iban}",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_IBAN,
		},
		{
			Name:        "PmtInf_CdtTrfTx_RmtInf_Ustrd",
			SourceValue: "AT41/{$.payee-e2e-rtp-ref}/AT05/pagamento fattura",
			TargetPath:  pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Ustrd,
		},
	}

	b, err := yaml.Marshal(dm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fn, b, fs.ModePerm)
	return err
}
