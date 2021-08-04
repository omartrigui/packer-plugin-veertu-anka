// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package ankaregistry

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion   *string           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	RegistryName        *string           `mapstructure:"registry_name" cty:"registry_name" hcl:"registry_name"`
	RegistryURL         *string           `mapstructure:"registry_path" cty:"registry_path" hcl:"registry_path"`
	NodeCertPath        *string           `mapstructure:"cert" cty:"cert" hcl:"cert"`
	NodeKeyPath         *string           `mapstructure:"key" cty:"key" hcl:"key"`
	CaRootPath          *string           `mapstructure:"cacert" cty:"cacert" hcl:"cacert"`
	IsInsecure          *bool             `mapstructure:"insecure" cty:"insecure" hcl:"insecure"`
	Tag                 *string           `mapstructure:"tag" cty:"tag" hcl:"tag"`
	Description         *string           `mapstructure:"description" cty:"description" hcl:"description"`
	RemoteVM            *string           `mapstructure:"remote_vm" cty:"remote_vm" hcl:"remote_vm"`
	Local               *bool             `mapstructure:"local" cty:"local" hcl:"local"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"registry_name":              &hcldec.AttrSpec{Name: "registry_name", Type: cty.String, Required: false},
		"registry_path":              &hcldec.AttrSpec{Name: "registry_path", Type: cty.String, Required: false},
		"cert":                       &hcldec.AttrSpec{Name: "cert", Type: cty.String, Required: false},
		"key":                        &hcldec.AttrSpec{Name: "key", Type: cty.String, Required: false},
		"cacert":                     &hcldec.AttrSpec{Name: "cacert", Type: cty.String, Required: false},
		"insecure":                   &hcldec.AttrSpec{Name: "insecure", Type: cty.Bool, Required: false},
		"tag":                        &hcldec.AttrSpec{Name: "tag", Type: cty.String, Required: false},
		"description":                &hcldec.AttrSpec{Name: "description", Type: cty.String, Required: false},
		"remote_vm":                  &hcldec.AttrSpec{Name: "remote_vm", Type: cty.String, Required: false},
		"local":                      &hcldec.AttrSpec{Name: "local", Type: cty.Bool, Required: false},
	}
	return s
}