// Autogenerated
package module

import (
	"aig/pkg/ansible"
)

type Find struct {
	ModuleName string
	Params     FindParams
	Result     FindResult
}

type FindParams struct {

	// Age
	Age string `yaml:"age,omitempty" json:"age,omitempty"`

	// AgeStamp
	AgeStamp string `yaml:"age_stamp,omitempty" json:"age_stamp,omitempty"`

	// Contains
	Contains string `yaml:"contains,omitempty" json:"contains,omitempty"`

	// Depth
	Depth int `yaml:"depth,omitempty" json:"depth,omitempty"`

	// Excludes
	Excludes []string `yaml:"excludes,omitempty" json:"excludes,omitempty"`

	// FileType
	FileType string `yaml:"file_type,omitempty" json:"file_type,omitempty"`

	// Follow
	Follow bool `yaml:"follow,omitempty" json:"follow,omitempty"`

	// GetChecksum
	GetChecksum bool `yaml:"get_checksum,omitempty" json:"get_checksum,omitempty"`

	// Hidden
	Hidden bool `yaml:"hidden,omitempty" json:"hidden,omitempty"`

	// Paths
	Paths []string `yaml:"paths,omitempty" json:"paths,omitempty"`

	// Patterns
	Patterns []string `yaml:"patterns,omitempty" json:"patterns,omitempty"`

	// ReadWholeFile
	ReadWholeFile bool `yaml:"read_whole_file,omitempty" json:"read_whole_file,omitempty"`

	// Recurse
	Recurse bool `yaml:"recurse,omitempty" json:"recurse,omitempty"`

	// Size
	Size string `yaml:"size,omitempty" json:"size,omitempty"`

	// UseRegex
	UseRegex bool `yaml:"use_regex,omitempty" json:"use_regex,omitempty"`
}

type FindResult struct {
	Raw string

	// Examined
	Examined int `yaml:"examined,omitempty" json:"examined,omitempty"`

	// Files
	Files []map[string]interface{} `yaml:"files,omitempty" json:"files,omitempty"`

	// Matched
	Matched int `yaml:"matched,omitempty" json:"matched,omitempty"`

	// SkippedPaths
	SkippedPaths map[string]interface{} `yaml:"skipped_paths,omitempty" json:"skipped_paths,omitempty"`
}

func (m *Find) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewFind() *Find {
	return &Find{
		ModuleName: "find",
	}
}
