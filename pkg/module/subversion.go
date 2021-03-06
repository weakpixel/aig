// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("subversion", func() Module {
		return NewSubversion()
	})
}

type Subversion struct {
	ModuleName string
	Params     SubversionParams
	Result     SubversionResult
}

type SubversionParams struct {

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
	types.CommonReturn
	Raw string
}

func (m *Subversion) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func (m *Subversion) GetResult() interface{} {
	return &m.Result
}

func (m *Subversion) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Subversion) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Subversion) GetParams() interface{} {
	return &m.Params
}

func (m *Subversion) GetType() string {
	return m.ModuleName
}

func NewSubversion() *Subversion {
	return &Subversion{
		ModuleName: "subversion",
	}
}
