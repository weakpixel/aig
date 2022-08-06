package module

// Autogenerated file

import (
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("sysvinit", func() Module {
		return NewSysvinit()
	})
}

//
// Sysvinit (sysvinit) - Manage SysV services.
//
func NewSysvinit() *Sysvinit {
	return &Sysvinit{}
}

// Sysvinit (sysvinit) - Manage SysV services.
//
// Controls services on target hosts that use the SysV init system.
//
// Source: https://github.com/ansible/ansible/blob/v2.13.1/lib/ansible/modules/sysvinit.py
type Sysvinit struct {
	Params SysvinitParams
	Result SysvinitResult
}

type SysvinitParams struct {

	// Arguments
	// Additional arguments provided on the command line that some init scripts accept.
	//
	// Default: <no value>
	// Required: false
	Arguments string `yaml:"arguments,omitempty" json:"arguments,omitempty"`

	// Daemonize
	// Have the module daemonize as the service itself might not do so properly.
	// This is useful with badly written init scripts or daemons, which commonly manifests as the task hanging as it is still holding the tty or the service dying when the task is over as the connection closes the session.
	//
	// Default: no
	// Required: false
	Daemonize bool `yaml:"daemonize,omitempty" json:"daemonize,omitempty"`

	// Enabled
	// Whether the service should start on boot. B(At least one of state and enabled are required.)
	//
	// Default: <no value>
	// Required: false
	Enabled bool `yaml:"enabled,omitempty" json:"enabled,omitempty"`

	// Name
	// Name of the service.
	//
	// Default: <no value>
	// Required: true
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// Pattern
	// A substring to look for as would be found in the output of the I(ps) command as a stand-in for a status result.
	// If the string is found, the service will be assumed to be running.
	// This option is mainly for use with init scripts that don't support the 'status' option.
	//
	// Default: <no value>
	// Required: false
	Pattern string `yaml:"pattern,omitempty" json:"pattern,omitempty"`

	// Runlevels
	// The runlevels this script should be enabled/disabled from.
	// Use this to override the defaults set by the package or init script itself.
	//
	// Default: <no value>
	// Required: false
	Runlevels []string `yaml:"runlevels,omitempty" json:"runlevels,omitempty"`

	// Sleep
	// If the service is being C(restarted) or C(reloaded) then sleep this many seconds between the stop and start command. This helps to workaround badly behaving services.
	//
	// Default: 1
	// Required: false
	Sleep int `yaml:"sleep,omitempty" json:"sleep,omitempty"`

	// State
	// C(started)/C(stopped) are idempotent actions that will not run commands unless necessary. Not all init scripts support C(restarted) nor C(reloaded) natively, so these will both trigger a stop and start as needed.
	//
	// Default: <no value>
	// Required: false
	State string `yaml:"state,omitempty" json:"state,omitempty"`
}

type SysvinitResult struct {
	types.CommonReturn
	Raw string

	// Results
	// results from actions taken
	Results interface{} `yaml:"results,omitempty" json:"results,omitempty"`
}

func (m *Sysvinit) GetResult() interface{} {
	return &m.Result
}

func (m *Sysvinit) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Sysvinit) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Sysvinit) GetParams() interface{} {
	return &m.Params
}

func (m *Sysvinit) GetType() string {
	return "sysvinit"
}
