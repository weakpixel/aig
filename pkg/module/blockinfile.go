// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("blockinfile", func() Module {
		return NewBlockinfile()
	})
}

type Blockinfile struct {
	ModuleName string
	Params     BlockinfileParams
	Result     BlockinfileResult
}

type BlockinfileParams struct {

	// Backup
	Backup bool `yaml:"backup,omitempty" json:"backup,omitempty"`

	// Block
	Block string `yaml:"block,omitempty" json:"block,omitempty"`

	// Create
	Create bool `yaml:"create,omitempty" json:"create,omitempty"`

	// Insertafter
	Insertafter string `yaml:"insertafter,omitempty" json:"insertafter,omitempty"`

	// Insertbefore
	Insertbefore string `yaml:"insertbefore,omitempty" json:"insertbefore,omitempty"`

	// Marker
	Marker string `yaml:"marker,omitempty" json:"marker,omitempty"`

	// MarkerBegin
	MarkerBegin string `yaml:"marker_begin,omitempty" json:"marker_begin,omitempty"`

	// MarkerEnd
	MarkerEnd string `yaml:"marker_end,omitempty" json:"marker_end,omitempty"`

	// Path
	Path string `yaml:"path,omitempty" json:"path,omitempty"`

	// State
	State string `yaml:"state,omitempty" json:"state,omitempty"`
}

type BlockinfileResult struct {
	types.CommonReturn
	Raw string
}

func (m *Blockinfile) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func (m *Blockinfile) GetResult() interface{} {
	return &m.Result
}

func (m *Blockinfile) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Blockinfile) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Blockinfile) GetParams() interface{} {
	return &m.Params
}

func (m *Blockinfile) GetType() string {
	return m.ModuleName
}

func NewBlockinfile() *Blockinfile {
	return &Blockinfile{
		ModuleName: "blockinfile",
	}
}
