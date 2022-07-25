package model_test

import (
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
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

func NewModel() (model.GoModel, error) {

	schemas := []string{
		"~/iso-20022/schemas/pain.013.001.07.xsd", "~/iso-20022/schemas/pain.014.001.07.xsd",
	}

	p := parser.NewParser(parser.Config{Registry: registry.Config{SimpleTypesInCommonPackage: true}})
	for _, xsdFileName := range schemas {

		msgName := strings.TrimSuffix(filepath.Base(xsdFileName), ".xsd")
		log.Info().Msgf("Msg: %s", msgName)

		fn, ok := util.ResolvePath(xsdFileName)
		if !ok {
			return model.GoModel{}, fmt.Errorf("could not resolve %s", xsdFileName)
		}

		b, err := ioutil.ReadFile(fn)
		if err != nil {
			return model.GoModel{}, err
		}

		msg := registry.ISO20022Message{Name: msgName}
		err = p.Parse(msg, b)
		if err != nil {
			return model.GoModel{}, err
		}
	}

	gm, err := model.NewModel(&model.DefaultModelCfg, []registry.ISO20022Message{{Name: "pain.013.001.07"}, {Name: "pain.014.001.07"}}, p.TypeRegistry)
	if err != nil {
		return model.GoModel{}, err
	}

	return gm, nil
}

var theModel model.GoModel

func TestMain(m *testing.M) {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	var err error
	theModel, err = NewModel()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGoModel(t *testing.T) {
	t.Log("Dumping GoModel.................")
	theModel.ShowInfo()
}

func TestSimpleVisitGoModel(t *testing.T) {

	t.Log("Simple Visiting GoModel.................")
	sv := &model.SimpleVisitor{}
	err := theModel.VisitDocument("pain_013_001_07", sv)
	require.NoError(t, err)

	sv.ShowInfo("")
	t.Log("number of leaves: ", sv.NumberOfLeaves)
}

func TestTreeVisitGoModel(t *testing.T) {
	t.Log("Tree Visiting GoModel.................")
	tv := &model.TreeVisitor{}
	err := theModel.VisitDocument("pain_013_001_07", tv)
	require.NoError(t, err)

	_ = tv.ShowInfo("")
}
