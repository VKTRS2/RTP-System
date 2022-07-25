package main

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/docmapper"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/camt_055_001_08"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
)

func Gen_AFC_DS10_Camt055_001_08(fn string) error {

	// In the normal case use the NewMapppingClass and avoid the WithFuncMap(nil) call that sets the builtins...
	dm, err := docmapper.NewMapperClass("camt.055.001.08", "AFC-DS10-camt.055.001.08")
	if err != nil {
		return err
	}

	dm.Rules = []docmapper.MappingRule{
		{
			Name:        "Assgnmt_Id",
			SourceValue: `"{$.payee-company-name}-DS10-" + objId()`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Assgnmt_Id,
		},
		{
			Name:        "Assgnmt_Assgnr_Pty_Nm",
			SourceValue: `{$.payee-company-name} {$.payee-name}`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Assgnmt_Assgnr_Pty_Nm,
		},
		{
			Name:        "Assgnmt_Assgnr_Pty_Id_OrgId_Othr_Id",
			SourceValue: `{$.payee-id}`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Assgnmt_Assgnr_Pty_Id_OrgId_Othr_Id,
		},
		{
			Name:        "Assgnmt_Assgnr_Pty_Id_OrgId_Othr_SchmeNm_Cd",
			SourceValue: `BOID`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Assgnmt_Assgnr_Pty_Id_OrgId_Othr_SchmeNm_Cd,
		},
		{
			Name:        "Assgnmt_Assgne_Agt_FinInstnId_BICFI",
			SourceValue: `{$.payer-psp-id}`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Assgnmt_Assgne_Agt_FinInstnId_BICFI,
		},
		{
			Name:        "Assgnmt_CreDtTm",
			SourceValue: `now()`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Assgnmt_CreDtTm,
		},
		{
			Name:        "Undrlyg_OrgnlGrpInfAndCxl_GrpCxlId",
			SourceValue: `{$.payee-company-name}-DS10-" + objId()`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlGrpInfAndCxl_GrpCxlId,
		},
		{
			Name:        "Undrlyg_OrgnlGrpInfAndCxl_OrgnlMsgId",
			SourceValue: `"{$.payee-company-name}-DS01-" + objId()`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlGrpInfAndCxl_OrgnlMsgId,
		},
		{
			Name:        "Undrlyg_OrgnlGrpInfAndCxl_OrgnlMsgNmId",
			SourceValue: `pain.013.001.07`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlGrpInfAndCxl_OrgnlMsgNmId,
		},
		{
			Name:        "Undrlyg_OrgnlGrpInfAndCxl_OrgnlCreDtTm",
			SourceValue: `now()`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlGrpInfAndCxl_OrgnlCreDtTm,
		},
		{
			Name:        "Undrlyg_OrgnlGrpInfAndCxl_NbOfTxs",
			SourceValue: `1`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlGrpInfAndCxl_NbOfTxs,
		},
		{
			Name:        "Undrlyg_OrgnlGrpInfAndCxl_CtrlSum",
			SourceValue: `{$.amount}`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlGrpInfAndCxl_CtrlSum,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_OrgnlPmtInfId",
			SourceValue: `"F/{$.payee-e2e-rtp-ref}/P/" + rtpId()`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_OrgnlPmtInfId,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlId",
			SourceValue: `{$.payee-company-name}-DS10-" + objId()`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlId,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlInstrId",
			SourceValue: `{$.payee-e2e-rtp-ref}`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlInstrId,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlEndToEndId",
			SourceValue: `{$.payee-e2e-rtp-ref}`,

			TargetPath: camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlEndToEndId,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlRsnInf_Orgtr_Nm",
			SourceValue: `{$.payee-company-name} {$.payee-name}`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlRsnInf_Orgtr_Nm,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlRsnInf_Orgtr_Id_OrgId_Othr_Id",
			SourceValue: `{$.payee-id}`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlRsnInf_Orgtr_Id_OrgId_Othr_Id,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Cd",
			SourceValue: `BOID`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Cd,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlRsnInf_Rsn_Cd",
			SourceValue: `{$.reason-code}`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlRsnInf_Rsn_Cd,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlRsnInf_AddtlInf",
			SourceValue: `{$.additional-info}`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlRsnInf_AddtlInf,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_Amt_InstdAmt_Ccy",
			SourceValue: `EUR`,

			TargetPath: camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_Amt_InstdAmt_Ccy,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_Amt_InstdAmt_Value",
			SourceValue: `{$.amount}`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_Amt_InstdAmt_Value,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_ReqdExctnDt_Dt",
			SourceValue: `{$.pmt-req-exec-date}`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_ReqdExctnDt_Dt,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_PmtTpInf_SvcLvl_Cd",
			SourceValue: `SEPA`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_PmtTpInf_SvcLvl_Cd,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_PmtTpInf_LclInstrm_Prtry",
			SourceValue: `{$.pmt-instrument}`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_PmtTpInf_LclInstrm_Prtry,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_RmtInf_Ustrd",
			SourceValue: `AT41/{$.payee-e2e-rtp-ref}/AT05/pagamento fattura`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_RmtInf_Ustrd,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_Cdtr_Pty_Nm",
			SourceValue: `{$.payee-company-name} {$.payee_name}`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_Cdtr_Pty_Nm,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_Cdtr_Pty_Id_OrgId_Othr_Id",
			SourceValue: `{$.payee_id}`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_Cdtr_Pty_Id_OrgId_Othr_Id,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_Cdtr_Pty_Id_OrgId_Othr_SchmeNm_Cd",
			SourceValue: `BOID`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_Cdtr_Pty_Id_OrgId_Othr_SchmeNm_Cd,
		},
		{
			Name:        "Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_CdtrAcct_Id_IBAN",
			SourceValue: `{$.payee-iban}`,
			TargetPath:  camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_CdtrAcct_Id_IBAN,
		},
	}

	b, err := yaml.Marshal(dm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fn, b, fs.ModePerm)
	return err
}
