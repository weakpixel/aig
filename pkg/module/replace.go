// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
)

type Replace struct {
	ModuleName string
	Params     ReplaceParams
	Result     ReplaceResult
}

type ReplaceParams struct {

	// After
	After string `yaml:"after,omitempty" json:"after,omitempty"`

	// Backup
	Backup bool `yaml:"backup,omitempty" json:"backup,omitempty"`

	// Before
	Before string `yaml:"before,omitempty" json:"before,omitempty"`

	// Encoding
	Encoding string `yaml:"encoding,omitempty" json:"encoding,omitempty"`

	// Others
	Others string `yaml:"others,omitempty" json:"others,omitempty"`

	// Path
	Path string `yaml:"path,omitempty" json:"path,omitempty"`

	// Regexp
	Regexp string `yaml:"regexp,omitempty" json:"regexp,omitempty"`

	// Replace
	Replace string `yaml:"replace,omitempty" json:"replace,omitempty"`
}

type ReplaceResult struct {
	Raw string
}

func (m *Replace) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewReplace() *Replace {
	return &Replace{
		ModuleName: "replace",
	}
}
