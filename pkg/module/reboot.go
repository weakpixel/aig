
// Autogenerated
package module

import (
	"aig/pkg/ansible"
	
)

type Reboot struct {
	ModuleName string
	Options RebootOptions
	Result RebootResult
}

type RebootOptions struct {
	
		// BootTimeCommand 
		BootTimeCommand string `yaml:"boot_time_command,omitempty" json:"boot_time_command,omitempty"`
	
		// ConnectTimeout 
		ConnectTimeout int `yaml:"connect_timeout,omitempty" json:"connect_timeout,omitempty"`
	
		// Msg 
		Msg string `yaml:"msg,omitempty" json:"msg,omitempty"`
	
		// PostRebootDelay 
		PostRebootDelay int `yaml:"post_reboot_delay,omitempty" json:"post_reboot_delay,omitempty"`
	
		// PreRebootDelay 
		PreRebootDelay int `yaml:"pre_reboot_delay,omitempty" json:"pre_reboot_delay,omitempty"`
	
		// RebootCommand 
		RebootCommand string `yaml:"reboot_command,omitempty" json:"reboot_command,omitempty"`
	
		// RebootTimeout 
		RebootTimeout int `yaml:"reboot_timeout,omitempty" json:"reboot_timeout,omitempty"`
	
		// SearchPaths 
		SearchPaths []string `yaml:"search_paths,omitempty" json:"search_paths,omitempty"`
	
		// TestCommand 
		TestCommand string `yaml:"test_command,omitempty" json:"test_command,omitempty"`
	
}

type RebootResult struct {
	Raw string 
	
		// Elapsed 
		Elapsed int `yaml:"elapsed,omitempty" json:"elapsed,omitempty"`
	
		// Rebooted 
		Rebooted bool `yaml:"rebooted,omitempty" json:"rebooted,omitempty"`
	
}

func (m *Reboot) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Options, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewReboot() *Reboot {
	return &Reboot {
		ModuleName: "reboot",
	}
}
	