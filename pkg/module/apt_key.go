package module

// Autogenerated file

import (
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("apt_key", func() Module {
		return NewAptKey()
	})
}

//
// AptKey (apt_key) - Add or remove an apt key
//
func NewAptKey() *AptKey {
	return &AptKey{}
}

// AptKey (apt_key) - Add or remove an apt key
//
// Add or remove an I(apt) key, optionally downloading it.
//
// Source: https://github.com/ansible/ansible/blob/v2.13.1/lib/ansible/modules/apt_key.py
type AptKey struct {
	Params AptKeyParams
	Result AptKeyResult
}

type AptKeyParams struct {

	// Data
	// The keyfile contents to add to the keyring.
	//
	// Default: <no value>
	// Required: false
	Data string `yaml:"data,omitempty" json:"data,omitempty"`

	// File
	// The path to a keyfile on the remote server to add to the keyring.
	//
	// Default: <no value>
	// Required: false
	File string `yaml:"file,omitempty" json:"file,omitempty"`

	// Id
	// The identifier of the key.
	// Including this allows check mode to correctly report the changed state.
	// If specifying a subkey's id be aware that apt-key does not understand how to remove keys via a subkey id.  Specify the primary key's id instead.
	// This parameter is required when C(state) is set to C(absent).
	//
	// Default: <no value>
	// Required: false
	Id string `yaml:"id,omitempty" json:"id,omitempty"`

	// Keyring
	// The full path to specific keyring file in C(/etc/apt/trusted.gpg.d/).
	//
	// Default: <no value>
	// Required: false
	Keyring string `yaml:"keyring,omitempty" json:"keyring,omitempty"`

	// Keyserver
	// The keyserver to retrieve key from.
	//
	// Default: <no value>
	// Required: false
	Keyserver string `yaml:"keyserver,omitempty" json:"keyserver,omitempty"`

	// State
	// Ensures that the key is present (added) or absent (revoked).
	//
	// Default: present
	// Required: false
	State string `yaml:"state,omitempty" json:"state,omitempty"`

	// Url
	// The URL to retrieve key from.
	//
	// Default: <no value>
	// Required: false
	Url string `yaml:"url,omitempty" json:"url,omitempty"`

	// ValidateCerts
	// If C(no), SSL certificates for the target url will not be validated. This should only be used on personally controlled sites using self-signed certificates.
	//
	// Default: yes
	// Required: false
	ValidateCerts bool `yaml:"validate_certs,omitempty" json:"validate_certs,omitempty"`
}

type AptKeyResult struct {
	types.CommonReturn
	Raw string

	// After
	// List of apt key ids or fingerprints after any modification
	After []map[string]interface{} `yaml:"after,omitempty" json:"after,omitempty"`

	// Before
	// List of apt key ids or fingprints before any modifications
	Before []map[string]interface{} `yaml:"before,omitempty" json:"before,omitempty"`

	// Fp
	// Fingerprint of the key to import
	Fp string `yaml:"fp,omitempty" json:"fp,omitempty"`

	// Id
	// key id from source
	Id string `yaml:"id,omitempty" json:"id,omitempty"`

	// KeyId
	// calculated key id, it should be same as 'id', but can be different
	KeyId string `yaml:"key_id,omitempty" json:"key_id,omitempty"`

	// ShortId
	// caclulated short key id
	ShortId string `yaml:"short_id,omitempty" json:"short_id,omitempty"`
}

func (m *AptKey) GetResult() interface{} {
	return &m.Result
}

func (m *AptKey) GetResultRaw() string {
	return m.Result.Raw
}

func (m *AptKey) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *AptKey) GetParams() interface{} {
	return &m.Params
}

func (m *AptKey) GetType() string {
	return "apt_key"
}
