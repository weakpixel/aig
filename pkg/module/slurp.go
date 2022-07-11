// Autogenerated
package module

import (
	"aig/pkg/ansible"
)

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

func NewSlurp() *Slurp {
	return &Slurp{
		ModuleName: "slurp",
	}
}
