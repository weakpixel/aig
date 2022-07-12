// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("stat", func() Module {
		return NewStat()
	})
}

type Stat struct {
	ModuleName string
	Params     StatParams
	Result     StatResult
}

type StatParams struct {

	// ChecksumAlgorithm
	ChecksumAlgorithm string `yaml:"checksum_algorithm,omitempty" json:"checksum_algorithm,omitempty"`

	// Follow
	Follow bool `yaml:"follow,omitempty" json:"follow,omitempty"`

	// GetAttributes
	GetAttributes bool `yaml:"get_attributes,omitempty" json:"get_attributes,omitempty"`

	// GetChecksum
	GetChecksum bool `yaml:"get_checksum,omitempty" json:"get_checksum,omitempty"`

	// GetMime
	GetMime bool `yaml:"get_mime,omitempty" json:"get_mime,omitempty"`

	// Path
	Path string `yaml:"path,omitempty" json:"path,omitempty"`
}

type StatResult struct {
	types.CommonReturn
	Raw string

	// Stat
	Stat interface{} `yaml:"stat,omitempty" json:"stat,omitempty"`
}

func (m *Stat) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func (m *Stat) GetResult() interface{} {
	return &m.Result
}

func (m *Stat) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Stat) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Stat) GetParams() interface{} {
	return &m.Params
}

func (m *Stat) GetType() string {
	return m.ModuleName
}

func NewStat() *Stat {
	return &Stat{
		ModuleName: "stat",
	}
}
