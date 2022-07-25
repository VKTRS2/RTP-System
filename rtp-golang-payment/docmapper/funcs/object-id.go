package funcs

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
)

func NewObjectId() string {
	return util.NewObjectId().String()
}
