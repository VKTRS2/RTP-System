package examples_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/camt_055_001_08"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
	"github.com/stretchr/testify/require"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

const ()

func TestDocumentCamt_055_001_08(t *testing.T) {

	b, err := camt_055_001_08_Document.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_Camt_055_001_08_1, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_Camt_055_001_08_1)

}

func TestDocumentCamt_055_001_08_SetOps(t *testing.T) {
	d := camt_055_001_08.NewDocument()

	err := d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Assgnmt_Id, common.MustToMax35Text(AssignmentId))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Assgnmt_Assgnr_Pty_Nm, common.MustToMax140Text(AssignerName))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Assgnmt_Assgnr_Pty_Id_OrgId_Othr_Id, common.MustToMax35Text(AssignerOrgId))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Assgnmt_Assgnr_Pty_Id_OrgId_Othr_SchmeNm_Cd, common.MustToExternalOrganisationIdentification1Code("BOID"))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Assgnmt_Assgne_Agt_FinInstnId_BICFI, common.MustToBICFIDec2014Identifier(AssigneeAgentBic))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Assgnmt_CreDtTm, common.MustToISODateTime(time.Now()))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlGrpInfAndCxl_GrpCxlId, common.MustToMax35Text("camt055-DS10-20220330"))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlGrpInfAndCxl_OrgnlMsgId, common.MustToMax35Text("pacs013-DS01-20220322"))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlGrpInfAndCxl_OrgnlMsgNmId, common.MustToMax35Text("pain.013.001.07"))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlGrpInfAndCxl_OrgnlCreDtTm, common.MustToISODateTime(time.Now()))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlGrpInfAndCxl_NbOfTxs, common.MustToMax15NumericText("1"))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlGrpInfAndCxl_CtrlSum, xsdt.MustToDecimal(535.25))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_OrgnlPmtInfId, common.MustToMax35Text(OriginalPaymentInformationId))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlId, common.MustToMax35Text("CancellationId1234"))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlEndToEndId, common.MustToMax35Text(InvoiceNumber))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlRsnInf_Orgtr_Nm, common.MustToMax140Text(AssignerName))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlRsnInf_Orgtr_Id_OrgId_Othr_Id, common.MustToMax35Text(AssignerOrgId))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlRsnInf_Orgtr_Id_OrgId_Othr_SchmeNm_Cd, common.MustToExternalOrganisationIdentification1Code("BOID"))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlRsnInf_Rsn_Cd, common.MustToExternalCancellationReason1Code("AM09"))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_CxlRsnInf_AddtlInf, common.MustToMax105Text("error in previous invoice's amount"))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_Amt_InstdAmt_Ccy, common.MustToActiveOrHistoricCurrencyCode("EUR"))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_Amt_InstdAmt_Value, xsdt.MustToDecimal(535.25))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_ReqdExctnDt_Dt, common.MustToISODate(time.Now()))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_PmtTpInf_SvcLvl_Cd, common.MustToExternalServiceLevel1Code("SEPA"))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_PmtTpInf_LclInstrm_Prtry, common.MustToMax35Text("NOTPROVIDED"))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_RmtInf_Ustrd, common.MustToMax140Text(RemittanceInfo))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_Cdtr_Pty_Nm, common.MustToMax140Text(CreditorName))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_Cdtr_Pty_Id_OrgId_Othr_Id, common.MustToMax35Text(CreditorOrgId))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_Cdtr_Pty_Id_OrgId_Othr_SchmeNm_Cd, common.MustToExternalOrganisationIdentification1Code("BOID"))
	require.NoError(t, err)
	err = d.Set(camt_055_001_08.Path_CstmrPmtCxlReq_Undrlyg_OrgnlPmtInfAndCxl_TxInf_OrgnlTxRef_CdtrAcct_Id_IBAN, common.MustToIBAN2007Identifier(CreditorIBAN))
	require.NoError(t, err)

	b, err := d.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_Camt_055_001_08_1, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_Camt_055_001_08_1)

	b1, err := camt_055_001_08_Document.ToXML()
	require.NoError(t, err)

	require.Equal(t, len(b), len(b1), "size of arrays are different")
	numOfDiffs := 0
	for i := 0; i < len(b); i++ {
		if b[i] != b1[i] {
			numOfDiffs++
		}
	}

	require.Equal(t, numOfDiffs, 0, "number of diffs in project")

}
