// Autogenerated
package module

import (
	"aig/pkg/ansible"
)

type Package struct {
	ModuleName string
	Params     PackageParams
	Result     PackageResult
}

type PackageParams struct {

	// Name
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// State
	State string `yaml:"state,omitempty" json:"state,omitempty"`

	// Use
	Use string `yaml:"use,omitempty" json:"use,omitempty"`
}

type PackageResult struct {
	Raw string
}

func (m *Package) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewPackage() *Package {
	return &Package{
		ModuleName: "package",
	}
}
