// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
)

type Tempfile struct {
	ModuleName string
	Params     TempfileParams
	Result     TempfileResult
}

type TempfileParams struct {

	// Path
	Path string `yaml:"path,omitempty" json:"path,omitempty"`

	// Prefix
	Prefix string `yaml:"prefix,omitempty" json:"prefix,omitempty"`

	// State
	State string `yaml:"state,omitempty" json:"state,omitempty"`

	// Suffix
	Suffix string `yaml:"suffix,omitempty" json:"suffix,omitempty"`
}

type TempfileResult struct {
	Raw string

	// Path
	Path string `yaml:"path,omitempty" json:"path,omitempty"`
}

func (m *Tempfile) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewTempfile() *Tempfile {
	return &Tempfile{
		ModuleName: "tempfile",
	}
}
