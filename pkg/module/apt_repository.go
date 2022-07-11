// Autogenerated
package module

import (
	"aig/pkg/ansible"
)

type AptRepository struct {
	ModuleName string
	Params     AptRepositoryParams
	Result     AptRepositoryResult
}

type AptRepositoryParams struct {

	// Codename
	Codename string `yaml:"codename,omitempty" json:"codename,omitempty"`

	// Filename
	Filename string `yaml:"filename,omitempty" json:"filename,omitempty"`

	// InstallPythonApt
	InstallPythonApt bool `yaml:"install_python_apt,omitempty" json:"install_python_apt,omitempty"`

	// Mode
	Mode string `yaml:"mode,omitempty" json:"mode,omitempty"`

	// Repo
	Repo string `yaml:"repo,omitempty" json:"repo,omitempty"`

	// State
	State string `yaml:"state,omitempty" json:"state,omitempty"`

	// UpdateCache
	UpdateCache bool `yaml:"update_cache,omitempty" json:"update_cache,omitempty"`

	// UpdateCacheRetries
	UpdateCacheRetries int `yaml:"update_cache_retries,omitempty" json:"update_cache_retries,omitempty"`

	// UpdateCacheRetryMaxDelay
	UpdateCacheRetryMaxDelay int `yaml:"update_cache_retry_max_delay,omitempty" json:"update_cache_retry_max_delay,omitempty"`

	// ValidateCerts
	ValidateCerts bool `yaml:"validate_certs,omitempty" json:"validate_certs,omitempty"`
}

type AptRepositoryResult struct {
	Raw string
}

func (m *AptRepository) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewAptRepository() *AptRepository {
	return &AptRepository{
		ModuleName: "apt_repository",
	}
}
