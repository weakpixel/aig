// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
)

type Debconf struct {
	ModuleName string
	Params     DebconfParams
	Result     DebconfResult
}

type DebconfParams struct {

	// Name
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// Question
	Question string `yaml:"question,omitempty" json:"question,omitempty"`

	// Unseen
	Unseen bool `yaml:"unseen,omitempty" json:"unseen,omitempty"`

	// Value
	Value string `yaml:"value,omitempty" json:"value,omitempty"`

	// Vtype
	Vtype string `yaml:"vtype,omitempty" json:"vtype,omitempty"`
}

type DebconfResult struct {
	Raw string
}

func (m *Debconf) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewDebconf() *Debconf {
	return &Debconf{
		ModuleName: "debconf",
	}
}
