package module

// Autogenerated file

import (
	"fmt"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("group", func() types.Module {
		return NewGroup()
	})
}

//
// Group (group) - Add or remove groups
//
func NewGroup() *Group {
	module := Group{}
	// Create dynamic param values
	paramValues := map[string]types.Value{}
	paramValues["gid"] = types.NewIntValue(&module.Params.Gid)
	paramValues["local"] = types.NewBoolValue(&module.Params.Local)
	paramValues["name"] = types.NewStringValue(&module.Params.Name)
	paramValues["non_unique"] = types.NewBoolValue(&module.Params.NonUnique)
	paramValues["state"] = types.NewStringValue(&module.Params.State)
	paramValues["system"] = types.NewBoolValue(&module.Params.System)
	module.Params.values = paramValues

	// Create dynamic result values
	resultValues := map[string]types.Value{}

	resultValues["gid"] = types.NewIntValue(&module.Result.Gid)
	resultValues["name"] = types.NewStringValue(&module.Result.Name)
	resultValues["state"] = types.NewStringValue(&module.Result.State)
	resultValues["system"] = types.NewBoolValue(&module.Result.System)
	module.Result.values = resultValues

	return &module
}

// Group (group) - Add or remove groups
//
// Manage presence of groups on a host.
//
// For Windows targets, use the M(ansible.windows.win_group) module instead.
//
//
// Source: https://github.com/ansible/ansible/blob/v2.13.1/lib/ansible/modules/group.py
type Group struct {
	Params GroupParams
	Result GroupResult
}

type GroupParams struct {

	// Gid
	// Optional I(GID) to set for the group.
	//
	// Default: <no value>
	// Required: false
	Gid int `yaml:"gid,omitempty" json:"gid,omitempty" cty:"gid"`

	// Local
	// Forces the use of "local" command alternatives on platforms that implement it.
	// This is useful in environments that use centralized authentication when you want to manipulate the local groups. (for example, it uses C(lgroupadd) instead of C(groupadd)).
	// This requires that these commands exist on the targeted host, otherwise it will be a fatal error.
	//
	// Default: no
	// Required: false
	Local bool `yaml:"local,omitempty" json:"local,omitempty" cty:"local"`

	// Name
	// Name of the group to manage.
	//
	// Default: <no value>
	// Required: true
	Name string `yaml:"name,omitempty" json:"name,omitempty" cty:"name"`

	// NonUnique
	// This option allows to change the group ID to a non-unique value. Requires C(gid).
	// Not supported on macOS or BusyBox distributions.
	//
	// Default: no
	// Required: false
	NonUnique bool `yaml:"non_unique,omitempty" json:"non_unique,omitempty" cty:"non_unique"`

	// State
	// Whether the group should be present or not on the remote host.
	//
	// Default: present
	// Required: false
	State string `yaml:"state,omitempty" json:"state,omitempty" cty:"state"`

	// System
	// If I(yes), indicates that the group created is a system group.
	//
	// Default: no
	// Required: false
	System bool `yaml:"system,omitempty" json:"system,omitempty" cty:"system"`

	values map[string]types.Value
}

func (p *GroupParams) Names() []string {
	names := []string{}
	for name := range p.values {
		names = append(names, name)
	}
	return names
}

func (p *GroupParams) Set(name string, value interface{}) error {
	v, ok := p.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (p *GroupParams) Get(name string) (interface{}, error) {
	v, ok := p.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

type GroupResult struct {
	types.CommonReturn
	Raw string

	// Gid
	// Group ID of the group.
	Gid int `yaml:"gid,omitempty" json:"gid,omitempty" cty:"gid"`

	// Name
	// Group name.
	Name string `yaml:"name,omitempty" json:"name,omitempty" cty:"name"`

	// State
	// Whether the group is present or not.
	State string `yaml:"state,omitempty" json:"state,omitempty" cty:"state"`

	// System
	// Whether the group is a system group or not.
	System bool `yaml:"system,omitempty" json:"system,omitempty" cty:"system"`

	values map[string]types.Value
}

func (r *GroupResult) Names() []string {
	names := []string{}
	for name := range r.values {
		names = append(names, name)
	}
	return names
}

func (r *GroupResult) Set(name string, value interface{}) error {
	v, ok := r.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (r *GroupResult) Get(name string) (interface{}, error) {
	v, ok := r.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

func (m *Group) GetResult() types.Result {
	return &m.Result
}

func (m *Group) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Group) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Group) GetParams() types.Params {
	return &m.Params
}

func (m *Group) GetType() string {
	return "group"
}
