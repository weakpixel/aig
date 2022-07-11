
// Autogenerated
package module

import (
	"aig/pkg/ansible"
	
)

type KnownHosts struct {
	ModuleName string
	Options KnownHostsOptions
	Result KnownHostsResult
}

type KnownHostsOptions struct {
	
		// HashHost 
		HashHost bool `yaml:"hash_host,omitempty" json:"hash_host,omitempty"`
	
		// Key 
		Key string `yaml:"key,omitempty" json:"key,omitempty"`
	
		// Name 
		Name string `yaml:"name,omitempty" json:"name,omitempty"`
	
		// Path 
		Path string `yaml:"path,omitempty" json:"path,omitempty"`
	
		// State 
		State string `yaml:"state,omitempty" json:"state,omitempty"`
	
}

type KnownHostsResult struct {
	Raw string 
	
}

func (m *KnownHosts) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Options, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewKnownHosts() *KnownHosts {
	return &KnownHosts {
		ModuleName: "known_hosts",
	}
}
	