// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
)

func init() {
	addModuleFactory("debconf", func() Module {
		return NewDebconf()
	})
}

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

func (m *Debconf) GetResult() interface{} {
	return &m.Result
}

func (m *Debconf) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Debconf) GetParams() interface{} {
	return &m.Params
}

func (m *Debconf) GetType() string {
	return m.ModuleName
}

func NewDebconf() *Debconf {
	return &Debconf{
		ModuleName: "debconf",
	}
}
