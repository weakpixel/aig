// Autogenerated
package module

import (
	"aig/pkg/ansible"
)

type RpmKey struct {
	ModuleName string
	Options    RpmKeyOptions
	Result     RpmKeyResult
}

type RpmKeyOptions struct {

	// Fingerprint
	Fingerprint string `yaml:"fingerprint,omitempty" json:"fingerprint,omitempty"`

	// Key
	Key string `yaml:"key,omitempty" json:"key,omitempty"`

	// State
	State string `yaml:"state,omitempty" json:"state,omitempty"`

	// ValidateCerts
	ValidateCerts bool `yaml:"validate_certs,omitempty" json:"validate_certs,omitempty"`
}

type RpmKeyResult struct {
	Raw string
}

func (m *RpmKey) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Options, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewRpmKey() *RpmKey {
	return &RpmKey{
		ModuleName: "rpm_key",
	}
}
