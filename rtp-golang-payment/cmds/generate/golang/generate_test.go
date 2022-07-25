package golang_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/cmds/generate/golang"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/cmds/generate/golang/model"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/cmds/generate/parser"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/cmds/generate/registry"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGoGenerate(t *testing.T) {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	schemas := []string{
		"~/iso-20022/schemas/pain.013.001.07.xsd", "~/iso-20022/schemas/pain.014.001.07.xsd",
	}

	p := parser.NewParser(parser.Config{Registry: registry.Config{SimpleTypesInCommonPackage: true}})
	for _, xsdFileName := range schemas {
		msgName := strings.TrimSuffix(filepath.Base(xsdFileName), ".xsd")
		t.Log("Msg: ", msgName)

		fn, ok := util.ResolvePath(xsdFileName)
		require.True(t, ok, "could not resolve %s", xsdFileName)

		b, err := ioutil.ReadFile(fn)
		require.NoError(t, err)

		msg := registry.ISO20022Message{Name: msgName}
		err = p.Parse(msg, b)
		require.NoError(t, err)
	}

	gm, err := model.NewModel(
		&model.DefaultModelCfg,
		[]registry.ISO20022Message{
			{Name: "pain.013.001.07"},
			{Name: "pain.014.001.07"},
		},
		p.TypeRegistry)
	require.NoError(t, err)

	fld, _ := util.ResolvePath("~/iso-20022/messages")
	require.NotEqual(t, fld, "", "could not resolve ~/iso-20022/messages path")
	cfg := golang.Config{
		OutFolder:  fld, // filepath.Join(fld, "messages"),
		FormatCode: true,
	}
	err = golang.Generate(&cfg, &gm)
	require.NoError(t, err)
}

func TestGoGenerateAll(t *testing.T) {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	schemaFolder, ok := util.ResolvePath("~/iso-20022/schemas/")
	require.True(t, ok, "cannot resolve schema folder")

	schemas, err := util.FindFiles(schemaFolder, util.WithFindFileType(util.FileTypeFile))
	require.NoError(t, err)
	require.NotEmpty(t, schemas, "could not find any schema")

	var msgs []registry.ISO20022Message
	p := parser.NewParser(parser.Config{Registry: registry.Config{SimpleTypesInCommonPackage: true}})
	for _, xsdFileName := range schemas {
		msgName := strings.TrimSuffix(filepath.Base(xsdFileName), ".xsd")
		t.Log("Msg: ", msgName)

		fn, ok := util.ResolvePath(xsdFileName)
		require.True(t, ok, "could not resolve %s", xsdFileName)

		b, err := ioutil.ReadFile(fn)
		require.NoError(t, err)

		msg := registry.ISO20022Message{Name: msgName}
		msgs = append(msgs, msg)

		err = p.Parse(msg, b)
		require.NoError(t, err)
	}

	gm, err := model.NewModel(&model.DefaultModelCfg, msgs, p.TypeRegistry)
	require.NoError(t, err)

	fld, _ := util.ResolvePath("~/iso-20022/messages")
	require.NotEqual(t, fld, "", "could not resolve ~/iso-20022/messages path")
	cfg := golang.Config{
		OutFolder:  fld, // filepath.Join(fld, "messages"),
		FormatCode: true,
	}
	err = golang.Generate(&cfg, &gm)
	require.NoError(t, err)
}
