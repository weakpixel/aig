// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("tempfile", func() Module {
		return NewTempfile()
	})
}

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
	types.CommonReturn
	Raw string

	// Path
	Path string `yaml:"path,omitempty" json:"path,omitempty"`
}

func (m *Tempfile) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func (m *Tempfile) GetResult() interface{} {
	return &m.Result
}

func (m *Tempfile) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Tempfile) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Tempfile) GetParams() interface{} {
	return &m.Params
}

func (m *Tempfile) GetType() string {
	return m.ModuleName
}

func NewTempfile() *Tempfile {
	return &Tempfile{
		ModuleName: "tempfile",
	}
}
