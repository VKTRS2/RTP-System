package model

import "path"

type ModelConfig struct {
	BasePackageImport string `yaml:"base-package-import" mapstructure:"base-package-import" json:"base-package-import"`
	XsDtPackageImport string `yaml:"xsdt-package-import" mapstructure:"xsdt-package-import" json:"xsdt-package-import"`
}

var DefaultModelCfg = ModelConfig{
	BasePackageImport: "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages",
	XsDtPackageImport: "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt",
}

func (c *ModelConfig) PostProcess() error {
	if c.XsDtPackageImport == "" {
		c.XsDtPackageImport = "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
	}

	return nil
}

func (c *ModelConfig) PackageImport(pkgName string) string {
	return path.Join(c.BasePackageImport, pkgName)
}
