// Autogenerated
package module

import (
	"aig/pkg/ansible"
)

type Debug struct {
	ModuleName string
	Options    DebugOptions
	Result     DebugResult
}

type DebugOptions struct {

	// Msg
	Msg string `yaml:"msg,omitempty" json:"msg,omitempty"`

	// Var
	Var string `yaml:"var,omitempty" json:"var,omitempty"`

	// Verbosity
	Verbosity int `yaml:"verbosity,omitempty" json:"verbosity,omitempty"`
}

type DebugResult struct {
	Raw string
}

func (m *Debug) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Options, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewDebug() *Debug {
	return &Debug{
		ModuleName: "debug",
	}
}
