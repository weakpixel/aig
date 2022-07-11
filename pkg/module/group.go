// Autogenerated
package module

import (
	"aig/pkg/ansible"
)

type Group struct {
	ModuleName string
	Params     GroupParams
	Result     GroupResult
}

type GroupParams struct {

	// Gid
	Gid int `yaml:"gid,omitempty" json:"gid,omitempty"`

	// Local
	Local bool `yaml:"local,omitempty" json:"local,omitempty"`

	// Name
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// NonUnique
	NonUnique bool `yaml:"non_unique,omitempty" json:"non_unique,omitempty"`

	// State
	State string `yaml:"state,omitempty" json:"state,omitempty"`

	// System
	System bool `yaml:"system,omitempty" json:"system,omitempty"`
}

type GroupResult struct {
	Raw string

	// Gid
	Gid int `yaml:"gid,omitempty" json:"gid,omitempty"`

	// Name
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// State
	State string `yaml:"state,omitempty" json:"state,omitempty"`

	// System
	System bool `yaml:"system,omitempty" json:"system,omitempty"`
}

func (m *Group) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewGroup() *Group {
	return &Group{
		ModuleName: "group",
	}
}
