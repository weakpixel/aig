// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
)

type Expect struct {
	ModuleName string
	Params     ExpectParams
	Result     ExpectResult
}

type ExpectParams struct {

	// Chdir
	Chdir string `yaml:"chdir,omitempty" json:"chdir,omitempty"`

	// Command
	Command string `yaml:"command,omitempty" json:"command,omitempty"`

	// Creates
	Creates string `yaml:"creates,omitempty" json:"creates,omitempty"`

	// Echo
	Echo bool `yaml:"echo,omitempty" json:"echo,omitempty"`

	// Removes
	Removes string `yaml:"removes,omitempty" json:"removes,omitempty"`

	// Responses
	Responses map[string]interface{} `yaml:"responses,omitempty" json:"responses,omitempty"`

	// Timeout
	Timeout int `yaml:"timeout,omitempty" json:"timeout,omitempty"`
}

type ExpectResult struct {
	Raw string
}

func (m *Expect) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewExpect() *Expect {
	return &Expect{
		ModuleName: "expect",
	}
}
