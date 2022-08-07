package module

// Autogenerated file

import (
	"fmt"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("reboot", func() types.Module {
		return NewReboot()
	})
}

//
// Reboot (reboot) - Reboot a machine
//
func NewReboot() *Reboot {
	module := Reboot{}
	// Create dynamic param values
	paramValues := map[string]types.Value{}
	paramValues["boot_time_command"] = types.NewStringValue(&module.Params.BootTimeCommand)
	paramValues["connect_timeout"] = types.NewIntValue(&module.Params.ConnectTimeout)
	paramValues["msg"] = types.NewStringValue(&module.Params.Msg)
	paramValues["post_reboot_delay"] = types.NewIntValue(&module.Params.PostRebootDelay)
	paramValues["pre_reboot_delay"] = types.NewIntValue(&module.Params.PreRebootDelay)
	paramValues["reboot_command"] = types.NewStringValue(&module.Params.RebootCommand)
	paramValues["reboot_timeout"] = types.NewIntValue(&module.Params.RebootTimeout)
	paramValues["search_paths"] = types.NewStringArrayValue(&module.Params.SearchPaths)
	paramValues["test_command"] = types.NewStringValue(&module.Params.TestCommand)
	module.Params.values = paramValues

	// Create dynamic result values
	resultValues := map[string]types.Value{}

	resultValues["elapsed"] = types.NewIntValue(&module.Result.Elapsed)
	resultValues["rebooted"] = types.NewBoolValue(&module.Result.Rebooted)
	module.Result.values = resultValues

	return &module
}

// Reboot (reboot) - Reboot a machine
//
// Reboot a machine, wait for it to go down, come back up, and respond to commands.
//
// For Windows targets, use the M(ansible.windows.win_reboot) module instead.
//
//
// Source: https://github.com/ansible/ansible/blob/v2.13.1/lib/ansible/modules/reboot.py
type Reboot struct {
	Params RebootParams
	Result RebootResult
}

type RebootParams struct {

	// BootTimeCommand
	// Command to run that returns a unique string indicating the last time the system was booted.
	// Setting this to a command that has different output each time it is run will cause the task to fail.
	//
	// Default: cat /proc/sys/kernel/random/boot_id
	// Required: false
	BootTimeCommand string `yaml:"boot_time_command,omitempty" json:"boot_time_command,omitempty"`

	// ConnectTimeout
	// Maximum seconds to wait for a successful connection to the managed hosts before trying again.
	// If unspecified, the default setting for the underlying connection plugin is used.
	//
	// Default: <no value>
	// Required: false
	ConnectTimeout int `yaml:"connect_timeout,omitempty" json:"connect_timeout,omitempty"`

	// Msg
	// Message to display to users before reboot.
	//
	// Default: Reboot initiated by Ansible
	// Required: false
	Msg string `yaml:"msg,omitempty" json:"msg,omitempty"`

	// PostRebootDelay
	// Seconds to wait after the reboot command was successful before attempting to validate the system rebooted successfully.
	// This is useful if you want wait for something to settle despite your connection already working.
	//
	// Default: 0
	// Required: false
	PostRebootDelay int `yaml:"post_reboot_delay,omitempty" json:"post_reboot_delay,omitempty"`

	// PreRebootDelay
	// Seconds to wait before reboot. Passed as a parameter to the reboot command.
	// On Linux, macOS and OpenBSD, this is converted to minutes and rounded down. If less than 60, it will be set to 0.
	// On Solaris and FreeBSD, this will be seconds.
	//
	// Default: 0
	// Required: false
	PreRebootDelay int `yaml:"pre_reboot_delay,omitempty" json:"pre_reboot_delay,omitempty"`

	// RebootCommand
	// Command to run that reboots the system, including any parameters passed to the command.
	// Can be an absolute path to the command or just the command name. If an absolute path to the command is not given, C(search_paths) on the target system will be searched to find the absolute path.
	// This will cause C(pre_reboot_delay), C(post_reboot_delay), and C(msg) to be ignored.
	//
	// Default: [determined based on target OS]
	// Required: false
	RebootCommand string `yaml:"reboot_command,omitempty" json:"reboot_command,omitempty"`

	// RebootTimeout
	// Maximum seconds to wait for machine to reboot and respond to a test command.
	// This timeout is evaluated separately for both reboot verification and test command success so the maximum execution time for the module is twice this amount.
	//
	// Default: 600
	// Required: false
	RebootTimeout int `yaml:"reboot_timeout,omitempty" json:"reboot_timeout,omitempty"`

	// SearchPaths
	// Paths to search on the remote machine for the C(shutdown) command.
	// I(Only) these paths will be searched for the C(shutdown) command. C(PATH) is ignored in the remote node when searching for the C(shutdown) command.
	//
	// Default: [/sbin /bin /usr/sbin /usr/bin /usr/local/sbin]
	// Required: false
	SearchPaths []string `yaml:"search_paths,omitempty" json:"search_paths,omitempty"`

	// TestCommand
	// Command to run on the rebooted host and expect success from to determine the machine is ready for further tasks.
	//
	// Default: whoami
	// Required: false
	TestCommand string `yaml:"test_command,omitempty" json:"test_command,omitempty"`

	values map[string]types.Value
}

func (p *RebootParams) Names() []string {
	names := []string{}
	for name := range p.values {
		names = append(names, name)
	}
	return names
}

func (p *RebootParams) Set(name string, value interface{}) error {
	v, ok := p.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (p *RebootParams) Get(name string) (interface{}, error) {
	v, ok := p.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

type RebootResult struct {
	types.CommonReturn
	Raw string

	// Elapsed
	// The number of seconds that elapsed waiting for the system to be rebooted.
	Elapsed int `yaml:"elapsed,omitempty" json:"elapsed,omitempty"`

	// Rebooted
	// true if the machine was rebooted
	Rebooted bool `yaml:"rebooted,omitempty" json:"rebooted,omitempty"`

	values map[string]types.Value
}

func (r *RebootResult) Names() []string {
	names := []string{}
	for name := range r.values {
		names = append(names, name)
	}
	return names
}

func (r *RebootResult) Set(name string, value interface{}) error {
	v, ok := r.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (r *RebootResult) Get(name string) (interface{}, error) {
	v, ok := r.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

func (m *Reboot) GetResult() types.Result {
	return &m.Result
}

func (m *Reboot) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Reboot) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Reboot) GetParams() types.Params {
	return &m.Params
}

func (m *Reboot) GetType() string {
	return "reboot"
}
