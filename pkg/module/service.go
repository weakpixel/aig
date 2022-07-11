// Autogenerated
package module

import (
	"aig/pkg/ansible"
)

type Service struct {
	ModuleName string
	Options    ServiceOptions
	Result     ServiceResult
}

type ServiceOptions struct {

	// Arguments
	Arguments string `yaml:"arguments,omitempty" json:"arguments,omitempty"`

	// Enabled
	Enabled bool `yaml:"enabled,omitempty" json:"enabled,omitempty"`

	// Name
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// Pattern
	Pattern string `yaml:"pattern,omitempty" json:"pattern,omitempty"`

	// Runlevel
	Runlevel string `yaml:"runlevel,omitempty" json:"runlevel,omitempty"`

	// Sleep
	Sleep int `yaml:"sleep,omitempty" json:"sleep,omitempty"`

	// State
	State string `yaml:"state,omitempty" json:"state,omitempty"`

	// Use
	Use string `yaml:"use,omitempty" json:"use,omitempty"`
}

type ServiceResult struct {
	Raw string
}

func (m *Service) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Options, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewService() *Service {
	return &Service{
		ModuleName: "service",
	}
}
