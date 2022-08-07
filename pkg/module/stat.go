package module

// Autogenerated file

import (
	"fmt"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("stat", func() types.Module {
		return NewStat()
	})
}

//
// Stat (stat) - Retrieve file or file system status
//
func NewStat() *Stat {
	module := Stat{}
	// Create dynamic param values
	paramValues := map[string]types.Value{}
	paramValues["checksum_algorithm"] = types.NewStringValue(&module.Params.ChecksumAlgorithm)
	paramValues["follow"] = types.NewBoolValue(&module.Params.Follow)
	paramValues["get_attributes"] = types.NewBoolValue(&module.Params.GetAttributes)
	paramValues["get_checksum"] = types.NewBoolValue(&module.Params.GetChecksum)
	paramValues["get_mime"] = types.NewBoolValue(&module.Params.GetMime)
	paramValues["path"] = types.NewStringValue(&module.Params.Path)
	module.Params.values = paramValues

	// Create dynamic result values
	resultValues := map[string]types.Value{}

	// NOT SUPPORTED: stat Stat interface{}
	module.Result.values = resultValues

	return &module
}

// Stat (stat) - Retrieve file or file system status
//
// Retrieves facts for a file similar to the Linux/Unix 'stat' command.
//
// For Windows targets, use the M(ansible.windows.win_stat) module instead.
//
//
// Source: https://github.com/ansible/ansible/blob/v2.13.1/lib/ansible/modules/stat.py
type Stat struct {
	Params StatParams
	Result StatResult
}

type StatParams struct {

	// ChecksumAlgorithm
	// Algorithm to determine checksum of file.
	// Will throw an error if the host is unable to use specified algorithm.
	// The remote host has to support the hashing method specified, C(md5) can be unavailable if the host is FIPS-140 compliant.
	//
	// Default: sha1
	// Required: false
	ChecksumAlgorithm string `yaml:"checksum_algorithm,omitempty" json:"checksum_algorithm,omitempty"`

	// Follow
	// Whether to follow symlinks.
	//
	// Default: no
	// Required: false
	Follow bool `yaml:"follow,omitempty" json:"follow,omitempty"`

	// GetAttributes
	// Get file attributes using lsattr tool if present.
	//
	// Default: yes
	// Required: false
	GetAttributes bool `yaml:"get_attributes,omitempty" json:"get_attributes,omitempty"`

	// GetChecksum
	// Whether to return a checksum of the file.
	//
	// Default: yes
	// Required: false
	GetChecksum bool `yaml:"get_checksum,omitempty" json:"get_checksum,omitempty"`

	// GetMime
	// Use file magic and return data about the nature of the file. this uses the 'file' utility found on most Linux/Unix systems.
	// This will add both C(mime_type) and C(charset) fields to the return, if possible.
	// In Ansible 2.3 this option changed from I(mime) to I(get_mime) and the default changed to C(yes).
	//
	// Default: yes
	// Required: false
	GetMime bool `yaml:"get_mime,omitempty" json:"get_mime,omitempty"`

	// Path
	// The full path of the file/object to get the facts of.
	//
	// Default: <no value>
	// Required: true
	Path string `yaml:"path,omitempty" json:"path,omitempty"`

	values map[string]types.Value
}

func (p *StatParams) Names() []string {
	names := []string{}
	for name := range p.values {
		names = append(names, name)
	}
	return names
}

func (p *StatParams) Set(name string, value interface{}) error {
	v, ok := p.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (p *StatParams) Get(name string) (interface{}, error) {
	v, ok := p.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

type StatResult struct {
	types.CommonReturn
	Raw string

	// Stat
	// Dictionary containing all the stat data, some platforms might add additional fields.
	Stat interface{} `yaml:"stat,omitempty" json:"stat,omitempty"`

	values map[string]types.Value
}

func (r *StatResult) Names() []string {
	names := []string{}
	for name := range r.values {
		names = append(names, name)
	}
	return names
}

func (r *StatResult) Set(name string, value interface{}) error {
	v, ok := r.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (r *StatResult) Get(name string) (interface{}, error) {
	v, ok := r.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

func (m *Stat) GetResult() types.Result {
	return &m.Result
}

func (m *Stat) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Stat) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Stat) GetParams() types.Params {
	return &m.Params
}

func (m *Stat) GetType() string {
	return "stat"
}
