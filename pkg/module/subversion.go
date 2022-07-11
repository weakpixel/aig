// Autogenerated
package module

import (
	"aig/pkg/ansible"
)

type Subversion struct {
	ModuleName string
	Options    SubversionOptions
	Result     SubversionResult
}

type SubversionOptions struct {

	// Checkout
	Checkout bool `yaml:"checkout,omitempty" json:"checkout,omitempty"`

	// Dest
	Dest string `yaml:"dest,omitempty" json:"dest,omitempty"`

	// Executable
	Executable string `yaml:"executable,omitempty" json:"executable,omitempty"`

	// Export
	Export bool `yaml:"export,omitempty" json:"export,omitempty"`

	// Force
	Force bool `yaml:"force,omitempty" json:"force,omitempty"`

	// InPlace
	InPlace bool `yaml:"in_place,omitempty" json:"in_place,omitempty"`

	// Password
	Password string `yaml:"password,omitempty" json:"password,omitempty"`

	// Repo
	Repo string `yaml:"repo,omitempty" json:"repo,omitempty"`

	// Revision
	Revision string `yaml:"revision,omitempty" json:"revision,omitempty"`

	// Switch
	Switch bool `yaml:"switch,omitempty" json:"switch,omitempty"`

	// Update
	Update bool `yaml:"update,omitempty" json:"update,omitempty"`

	// Username
	Username string `yaml:"username,omitempty" json:"username,omitempty"`

	// ValidateCerts
	ValidateCerts bool `yaml:"validate_certs,omitempty" json:"validate_certs,omitempty"`
}

type SubversionResult struct {
	Raw string
}

func (m *Subversion) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Options, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewSubversion() *Subversion {
	return &Subversion{
		ModuleName: "subversion",
	}
}
