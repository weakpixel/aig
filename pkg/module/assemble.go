// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("assemble", func() Module {
		return NewAssemble()
	})
}

type Assemble struct {
	ModuleName string
	Params     AssembleParams
	Result     AssembleResult
}

type AssembleParams struct {

	// Backup
	Backup bool `yaml:"backup,omitempty" json:"backup,omitempty"`

	// Delimiter
	Delimiter string `yaml:"delimiter,omitempty" json:"delimiter,omitempty"`

	// Dest
	Dest string `yaml:"dest,omitempty" json:"dest,omitempty"`

	// IgnoreHidden
	IgnoreHidden bool `yaml:"ignore_hidden,omitempty" json:"ignore_hidden,omitempty"`

	// Regexp
	Regexp string `yaml:"regexp,omitempty" json:"regexp,omitempty"`

	// RemoteSrc
	RemoteSrc bool `yaml:"remote_src,omitempty" json:"remote_src,omitempty"`

	// Src
	Src string `yaml:"src,omitempty" json:"src,omitempty"`

	// Validate
	Validate string `yaml:"validate,omitempty" json:"validate,omitempty"`
}

type AssembleResult struct {
	types.CommonReturn
	Raw string
}

func (m *Assemble) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func (m *Assemble) GetResult() interface{} {
	return &m.Result
}

func (m *Assemble) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Assemble) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Assemble) GetParams() interface{} {
	return &m.Params
}

func (m *Assemble) GetType() string {
	return m.ModuleName
}

func NewAssemble() *Assemble {
	return &Assemble{
		ModuleName: "assemble",
	}
}
