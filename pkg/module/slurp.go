// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
)

func init() {
	addModuleFactory("slurp", func() Module {
		return NewSlurp()
	})
}

type Slurp struct {
	ModuleName string
	Params     SlurpParams
	Result     SlurpResult
}

type SlurpParams struct {

	// Src
	Src string `yaml:"src,omitempty" json:"src,omitempty"`
}

type SlurpResult struct {
	Raw string

	// Content
	Content string `yaml:"content,omitempty" json:"content,omitempty"`

	// Encoding
	Encoding string `yaml:"encoding,omitempty" json:"encoding,omitempty"`

	// Source
	Source string `yaml:"source,omitempty" json:"source,omitempty"`
}

func (m *Slurp) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func (m *Slurp) GetResult() interface{} {
	return &m.Result
}

func (m *Slurp) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Slurp) GetParams() interface{} {
	return &m.Params
}

func (m *Slurp) GetType() string {
	return m.ModuleName
}

func NewSlurp() *Slurp {
	return &Slurp{
		ModuleName: "slurp",
	}
}
