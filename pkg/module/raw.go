
// Autogenerated
package module

import (
	"aig/pkg/ansible"
	
)

type Raw struct {
	ModuleName string
	Options RawOptions
	Result RawResult
}

type RawOptions struct {
	
		// Executable 
		Executable string `yaml:"executable,omitempty" json:"executable,omitempty"`
	
		// FreeForm 
		FreeForm string `yaml:"free_form,omitempty" json:"free_form,omitempty"`
	
}

type RawResult struct {
	Raw string 
	
}

func (m *Raw) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Options, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewRaw() *Raw {
	return &Raw {
		ModuleName: "raw",
	}
}
	