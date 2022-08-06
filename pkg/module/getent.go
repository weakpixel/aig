package module

// Autogenerated file

import (
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("getent", func() Module {
		return NewGetent()
	})
}

//
// Getent (getent) - A wrapper to the unix getent utility
//
func NewGetent() *Getent {
	return &Getent{}
}

// Getent (getent) - A wrapper to the unix getent utility
//
// Runs getent against one of it's various databases and returns information into the host's facts, in a getent_<database> prefixed variable.
type Getent struct {
	Params GetentParams
	Result GetentResult
}

type GetentParams struct {

	// Database
	// The name of a getent database supported by the target system (passwd, group, hosts, etc).
	//
	// Default: <no value>
	// Required: true
	Database string `yaml:"database,omitempty" json:"database,omitempty"`

	// FailKey
	// If a supplied key is missing this will make the task fail if C(yes).
	//
	// Default: yes
	// Required: false
	FailKey bool `yaml:"fail_key,omitempty" json:"fail_key,omitempty"`

	// Key
	// Key from which to return values from the specified database, otherwise the full contents are returned.
	//
	// Default:
	// Required: false
	Key string `yaml:"key,omitempty" json:"key,omitempty"`

	// Service
	// Override all databases with the specified service
	// The underlying system must support the service flag which is not always available.
	//
	// Default: <no value>
	// Required: false
	Service string `yaml:"service,omitempty" json:"service,omitempty"`

	// Split
	// Character used to split the database values into lists/arrays such as ':' or '	', otherwise  it will try to pick one depending on the database.
	//
	// Default: <no value>
	// Required: false
	Split string `yaml:"split,omitempty" json:"split,omitempty"`
}

type GetentResult struct {
	types.CommonReturn
	Raw string

	// AnsibleFacts
	// Facts to add to ansible_facts.
	AnsibleFacts map[string]interface{} `yaml:"ansible_facts,omitempty" json:"ansible_facts,omitempty"`
}

func (m *Getent) GetResult() interface{} {
	return &m.Result
}

func (m *Getent) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Getent) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Getent) GetParams() interface{} {
	return &m.Params
}

func (m *Getent) GetType() string {
	return "getent"
}
