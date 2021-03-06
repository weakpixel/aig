// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("file", func() Module {
		return NewFile()
	})
}

type File struct {
	ModuleName string
	Params     FileParams
	Result     FileResult
}

type FileParams struct {

	// AccessTime
	AccessTime string `yaml:"access_time,omitempty" json:"access_time,omitempty"`

	// AccessTimeFormat
	AccessTimeFormat string `yaml:"access_time_format,omitempty" json:"access_time_format,omitempty"`

	// Follow
	Follow bool `yaml:"follow,omitempty" json:"follow,omitempty"`

	// Force
	Force bool `yaml:"force,omitempty" json:"force,omitempty"`

	// ModificationTime
	ModificationTime string `yaml:"modification_time,omitempty" json:"modification_time,omitempty"`

	// ModificationTimeFormat
	ModificationTimeFormat string `yaml:"modification_time_format,omitempty" json:"modification_time_format,omitempty"`

	// Path
	Path string `yaml:"path,omitempty" json:"path,omitempty"`

	// Recurse
	Recurse bool `yaml:"recurse,omitempty" json:"recurse,omitempty"`

	// Src
	Src string `yaml:"src,omitempty" json:"src,omitempty"`

	// State
	State string `yaml:"state,omitempty" json:"state,omitempty"`
}

type FileResult struct {
	types.CommonReturn
	Raw string

	// Dest
	Dest string `yaml:"dest,omitempty" json:"dest,omitempty"`

	// Path
	Path string `yaml:"path,omitempty" json:"path,omitempty"`
}

func (m *File) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func (m *File) GetResult() interface{} {
	return &m.Result
}

func (m *File) GetResultRaw() string {
	return m.Result.Raw
}

func (m *File) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *File) GetParams() interface{} {
	return &m.Params
}

func (m *File) GetType() string {
	return m.ModuleName
}

func NewFile() *File {
	return &File{
		ModuleName: "file",
	}
}
