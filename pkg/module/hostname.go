package module

// Autogenerated file

import (
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("hostname", func() Module {
		return NewHostname()
	})
}

//
// Hostname (hostname) - Manage hostname
//
func NewHostname() *Hostname {
	return &Hostname{}
}

// Hostname (hostname) - Manage hostname
//
// Set system's hostname. Supports most OSs/Distributions including those using C(systemd).
// Windows, HP-UX, and AIX are not currently supported.
type Hostname struct {
	Params HostnameParams
	Result HostnameResult
}

type HostnameParams struct {

	// Name
	// Name of the host.
	// If the value is a fully qualified domain name that does not resolve from the given host, this will cause the module to hang for a few seconds while waiting for the name resolution attempt to timeout.
	//
	// Default: <no value>
	// Required: true
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// Use
	// Which strategy to use to update the hostname.
	// If not set we try to autodetect, but this can be problematic, particularly with containers as they can present misleading information.
	// Note that 'systemd' should be specified for RHEL/EL/CentOS 7+. Older distributions should use 'redhat'.
	//
	// Default: <no value>
	// Required: false
	Use string `yaml:"use,omitempty" json:"use,omitempty"`
}

type HostnameResult struct {
	types.CommonReturn
	Raw string
}

func (m *Hostname) GetResult() interface{} {
	return &m.Result
}

func (m *Hostname) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Hostname) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Hostname) GetParams() interface{} {
	return &m.Params
}

func (m *Hostname) GetType() string {
	return "hostname"
}
