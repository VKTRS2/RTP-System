package examples_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/pain_013_001_07"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
	"github.com/stretchr/testify/require"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestDocumentPain_013_001_07(t *testing.T) {

	b, err := pain_013_001_07_Document.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_Pain_013_001_07_1, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_Pain_013_001_07_1)

	d2, err := pain_013_001_07.NewDocumentFromXML(b)
	require.NoError(t, err)

	b2, err := d2.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_Pain_013_001_07_2, b2, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_Pain_013_001_07_2)

}

func TestDocumentPain_013_001_07_SetOps(t *testing.T) {
	d := pain_013_001_07.NewDocument()

	err := d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_MsgId, common.MustToMax35Text("pain013-DS01-20220322"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_CreDtTm, common.MustToISODateTime(time.Now()))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_NbOfTxs, common.MustToMax15NumericText("1"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm, common.MustToMax140Text(NameOfInitiatingParty))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_Id, common.MustToMax35Text(InitiatingPartyOrgId))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_Issr, common.MustToMax35Text(InitiatingPartyOrgIssr))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_PmtInfId, common.MustToMax35Text(PaymentInstructionId))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_PmtMtd, common.MustToPaymentMethod7Code(common.PaymentMethod7CodeTRF))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_SvcLvl_Cd, common.MustToExternalServiceLevel1Code("SRTP"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_LclInstrm_Prtry, common.MustToMax35Text("NOTPROVIDED"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_CtgyPurp_Cd, common.MustToExternalCategoryPurpose1Code("OTHR"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_ReqdExctnDt_Dt, common.MustToISODate(time.Now()))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_XpryDt_Dt, common.MustToISODate(time.Now()))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Nm, common.MustToMax140Text(DebitorName))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrLine, common.MustToMax70Text(DebitorAddress))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Id, common.MustToMax35Text("123456789"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd, common.MustToExternalPersonIdentification1Code("POID"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_IBAN, common.MustToIBAN2007Identifier(DebitorIBAN))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_BICFI, common.MustToBICFIDec2014Identifier(DebitorAgentBIC))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId_EndToEndId, common.MustToMax35Text(InvoiceNumber))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_SvcLvl_Cd, common.MustToExternalServiceLevel1Code("SRTP"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_LclInstrm_Prtry, common.MustToMax35Text("NOTPROVIDED"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_CtgyPurp_Cd, common.MustToExternalCategoryPurpose1Code("OTHR"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_AmtModAllwd, false)
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_EarlyPmtAllwd, false)
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_GrntedPmtReqd, false)
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Ccy, common.MustToActiveOrHistoricCurrencyCode("EUR"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Value, xsdt.MustToDecimal(535.35))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChrgBr, common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSLEV))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_BICFI, common.MustToBICFIDec2014Identifier(CreditorAgentBIC))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Nm, common.MustToMax140Text(CreditorName))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrLine, common.MustToMax70Text(CreditorAddress))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_Id, common.MustToMax35Text(CreditorOrgId))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_SchmeNm_Cd, common.MustToExternalOrganisationIdentification1Code("BOID"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_IBAN, common.MustToIBAN2007Identifier(CreditorIBAN))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Ustrd, common.MustToMax140Text(RemittanceInfo))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Ustrd, common.MustToMax140Text(RemittanceInfoAT87))
	require.NoError(t, err)

	b, err := d.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_Pain_013_001_07_1, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_Pain_013_001_07_1)

	b1, err := pain_013_001_07_Document.ToXML()
	require.NoError(t, err)

	require.Equal(t, len(b), len(b1), "size of arrays are different")
	numOfDiffs := 0
	for i := 0; i < len(b); i++ {
		if b[i] != b1[i] {
			numOfDiffs++
		}
	}

	// The difference might also due to the use of time.Now() function
	require.Equal(t, numOfDiffs, 0, "number of diffs in project")
	//t.Log(d.Get_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId_EndToEndId())

}

func TestDocumentPain_013_001_07_SetOp(t *testing.T) {
	d := pain_013_001_07.NewDocument()

	err := d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_MsgId, common.MustToMax35Text("pain013-DS01-20220322"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Id, common.MustToMax35Text("first-id"))
	require.NoError(t, err)
	err = d.Set("CdtrPmtActvtnReq.PmtInf.Dbtr.Id.PrvtId.Othr[+].Id", common.MustToMax35Text("second-id"))
	require.NoError(t, err)

	b, err := d.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_Pain_013_001_07_1, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_Pain_013_001_07_1)

	v, err := d.Get(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_MsgId)
	require.NoError(t, err)
	t.Logf("%s --> %v of %T", pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_MsgId, v, v)

	v, err = d.Get(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Id)
	require.NoError(t, err)
	t.Logf("%s --> %v of %T", pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Id, v, v)

	// properties not set.
	v, err = d.Get(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm)
	require.NoError(t, err)
	t.Logf("%s --> %v of %T", pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm, v, v)

	v, err = d.Get(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm)
	require.NoError(t, err)
	t.Logf("%s --> %v of %T", pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp, v, v)

	b1, err := d.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_Pain_013_001_07_2, b1, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_Pain_013_001_07_2)

	require.Equal(t, len(b), len(b1), "size of arrays are different")
	numOfDiffs := 0
	for i := 0; i < len(b); i++ {
		if b[i] != b1[i] {
			numOfDiffs++
		}
	}

	// The difference might also due to the use of time.Now() function
	require.Equal(t, numOfDiffs, 0, "number of diffs in project")
}

func TestDocumentPain_013_001_07_GetOp(t *testing.T) {
	d := pain_013_001_07.NewDocument()

	err := d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_MsgId, common.MustToMax35Text("pain013-DS01-20220322"))
	require.NoError(t, err)

	v, err := d.Get(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Id)
	require.NoError(t, err)
	t.Log(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Id, v)

	b, err := d.ToXML()
	require.NoError(t, err)
	t.Log(b)
}
