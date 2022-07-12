// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
)

type Hostname struct {
	ModuleName string
	Params     HostnameParams
	Result     HostnameResult
}

type HostnameParams struct {

	// Name
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// Use
	Use string `yaml:"use,omitempty" json:"use,omitempty"`
}

type HostnameResult struct {
	Raw string
}

func (m *Hostname) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewHostname() *Hostname {
	return &Hostname{
		ModuleName: "hostname",
	}
}
