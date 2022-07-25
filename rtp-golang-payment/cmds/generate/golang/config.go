package golang

import (
	"os"
	"path/filepath"
)

//type ModelConfig struct {
//	CommonPackageImport string `yaml:"common-package-import" mapstructure:"common-package-import" json:"common-package-import"`
//	XsDtPackageImport   string `yaml:"xsdt-package-import" mapstructure:"xsdt-package-import" json:"xsdt-package-import"`
//}
//
//var DefaultModelCfg = ModelConfig{
//	CommonPackageImport: "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common",
//	XsDtPackageImport:   "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt",
//}
//
//func (c *ModelConfig) PostProcess() error {
//	if c.XsDtPackageImport == "" {
//		c.XsDtPackageImport = "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
//	}
//
//	return nil
//}

type Config struct {
	OutFolder  string `yaml:"out-folder" mapstructure:"out-folder" json:"out-folder"`
	FormatCode bool   `yaml:"format-code" mapstructure:"format-code" json:"format-code"`
}

func (cfg *Config) SetupOutFolder(folderName string) (string, error) {

	contentPath := filepath.Join(cfg.OutFolder, folderName)
	if _, err := os.Stat(contentPath); os.IsNotExist(err) {
		if err = os.MkdirAll(contentPath, 0755); err != nil {
			return "", err
		}
	}

	return contentPath, nil
}
