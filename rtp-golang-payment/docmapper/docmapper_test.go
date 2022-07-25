package docmapper_test

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/docmapper"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/pain_013_001_07"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

const (
	ISO20022MessageFile = "iso20022-message.xml"
)

var sample = []byte(`
   {
      "payee-company-name":"PPAY",
      "payee-iban":"IT90U0760103200001042598423",
      "payee-name":"Postepay S.p.A",
      "payee-id":"06874351007",
      "payee-e2e-rtp-ref":"0000008121000020",
      "payee-psp-bic":"BPPIITRRXXX",
      "payer-id":"0000032551150363",
      "amount":"13542",
      "payer-psp-id":"BPPIITRRXXX",
      "rtp-expiry-date":"2021-01-26",
      "pmt-req-exec-date":"2021-01-26",
      "pmt-instrument":"NOT PROVIDED",
      "rtp-timestamp":"2022-04-06 17:27:06"
      ,"rmt-inf": "pagamento fattura"
      ,"rmt-inf-87": "pagamento fattura 2"
   }
`)

func TestMap2Pain013Doc(t *testing.T) {

	numClasses, err := docmapper.InitEmbeddedRegistry(docClassFS, docClassFSRootPath)
	require.NoError(t, err)

	t.Log("num loaded classes", numClasses)

	m := make(docmapper.MappableDocMap)
	err = json.Unmarshal(sample, &m)
	require.NoError(t, err)

	dm := docmapper.Registry["AFC-DS01-pain.013.001.07"]
	d := pain_013_001_07.NewDocument()
	err = dm.Map(m, &d, true)
	require.NoError(t, err)

	t.Log("is document valid? --> ", d.IsValid(false))

	b1, err := d.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(ISO20022MessageFile, b1, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(ISO20022MessageFile)
}

func TestPain013Doc2Map(t *testing.T) {

	// In the normal case use the NewMapppingClass and avoid the WithFuncMap(nil) call that sets the builtins...
	dm, err := docmapper.NewMapperClass("pain.013.001.07", "AFC-DS01-pain.013.001.07")
	require.NoError(t, err)

	dm.Rules = []docmapper.MappingRule{
		{
			Name:        "msg-id",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_MsgId),
			MapFunc:     "trimSpace",
			TargetPath:  "msg-id",
		},
		{
			Name:        "from-name",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm),
			MapFunc:     "trimSpace",
			TargetPath:  "header.from.name",
		},
		{
			Name:        "creation-date",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_CreDtTm),
			MapFunc:     "trimSpace",
			TargetPath:  "header.creation-date",
		},
		{
			Name:        "number-of-txs",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_NbOfTxs),
			MapFunc:     "trimSpace",
			TargetPath:  "header.number-of-txs",
		},
		{
			Name:        "pmt-info-amount",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Ccy),
			MapFunc:     "trimSpace",
			TargetPath:  "message.pmtinfo.amount.currency",
		},
		{
			Name:        "pmt-info-amount",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Value),
			MapFunc:     "trimSpace",
			TargetPath:  "message.pmtinfo.amount.value",
		},
	}

	d := pain_013_001_07.NewDocument()
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_MsgId, common.MustToMax35Text("pain013-DS01-20220322"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_CreDtTm, common.MustToISODateTime(time.Now()))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_NbOfTxs, common.MustToMax15NumericText("7"))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm, common.MustToMax140Text("ACME Corp."))
	require.NoError(t, err)
	err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Value, "535.90")
	require.NoError(t, err)

	m := docmapper.MappableDocMap(make(map[string]interface{}))

	err = dm.Map(&d, m, true)
	require.NoError(t, err)

	b1, err := yaml.Marshal(dm)
	require.NoError(t, err)
	t.Log(string(b1))

	b, err := json.MarshalIndent(m, "", " ")
	require.NoError(t, err)

	t.Log(string(b))
}

func TestDocMapperFile(t *testing.T) {
	// In the normal case use the NewMapppingClass and avoid the WithFuncMap(nil) call that sets the builtins...
	dm, err := docmapper.NewMapperClass("pain.013.001.07", "AFC-DS01-pain.013.001.07")
	require.NoError(t, err)

	dm.Rules = []docmapper.MappingRule{
		{
			Name:        "msg-id",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_MsgId),
			MapFunc:     "trimSpace",
			TargetPath:  "msg-id",
		},
		{
			Name:        "from-name",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm),
			MapFunc:     "trimSpace",
			TargetPath:  "header.from.name",
		},
		{
			Name:        "creation-date",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_CreDtTm),
			MapFunc:     "trimSpace",
			TargetPath:  "header.creation-date",
		},
		{
			Name:        "number-of-txs",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_NbOfTxs),
			MapFunc:     "trimSpace",
			TargetPath:  "header.number-of-txs",
		},
		{
			Name:        "pmt-info-amount",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Ccy),
			MapFunc:     "trimSpace",
			TargetPath:  "message.pmtinfo.amount.currency",
		},
		{
			Name:        "pmt-info-amount",
			SourceValue: fmt.Sprintf("{$.%s}", pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Value),
			MapFunc:     "trimSpace",
			TargetPath:  "message.pmtinfo.amount.value",
		},
	}

	b, err := yaml.Marshal(dm)
	require.NoError(t, err)

	t.Log(string(b))
}

// The folder contains a number of .yml files each one for a different class
//go:embed defs/*yml
var docClassFS embed.FS
var docClassFSRootPath = "defs"

func TestEmbeddedRegistry(t *testing.T) {

	numClasses, err := docmapper.InitEmbeddedRegistry(docClassFS, docClassFSRootPath)
	require.NoError(t, err)

	t.Log("num loaded classes", numClasses)

}
