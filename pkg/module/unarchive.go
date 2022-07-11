
// Autogenerated
package module

import (
	"aig/pkg/ansible"
	
)

type Unarchive struct {
	ModuleName string
	Options UnarchiveOptions
	Result UnarchiveResult
}

type UnarchiveOptions struct {
	
		// Copy 
		Copy bool `yaml:"copy,omitempty" json:"copy,omitempty"`
	
		// Creates 
		Creates string `yaml:"creates,omitempty" json:"creates,omitempty"`
	
		// Dest 
		Dest string `yaml:"dest,omitempty" json:"dest,omitempty"`
	
		// Exclude 
		Exclude []string `yaml:"exclude,omitempty" json:"exclude,omitempty"`
	
		// ExtraOpts 
		ExtraOpts []string `yaml:"extra_opts,omitempty" json:"extra_opts,omitempty"`
	
		// Include 
		Include []string `yaml:"include,omitempty" json:"include,omitempty"`
	
		// IoBufferSize 
		IoBufferSize int `yaml:"io_buffer_size,omitempty" json:"io_buffer_size,omitempty"`
	
		// KeepNewer 
		KeepNewer bool `yaml:"keep_newer,omitempty" json:"keep_newer,omitempty"`
	
		// ListFiles 
		ListFiles bool `yaml:"list_files,omitempty" json:"list_files,omitempty"`
	
		// RemoteSrc 
		RemoteSrc bool `yaml:"remote_src,omitempty" json:"remote_src,omitempty"`
	
		// Src 
		Src string `yaml:"src,omitempty" json:"src,omitempty"`
	
		// ValidateCerts 
		ValidateCerts bool `yaml:"validate_certs,omitempty" json:"validate_certs,omitempty"`
	
}

type UnarchiveResult struct {
	Raw string 
	
		// Dest 
		Dest string `yaml:"dest,omitempty" json:"dest,omitempty"`
	
		// Files 
		Files []map[string]interface{} `yaml:"files,omitempty" json:"files,omitempty"`
	
		// Gid 
		Gid int `yaml:"gid,omitempty" json:"gid,omitempty"`
	
		// Group 
		Group string `yaml:"group,omitempty" json:"group,omitempty"`
	
		// Handler 
		Handler string `yaml:"handler,omitempty" json:"handler,omitempty"`
	
		// Mode 
		Mode string `yaml:"mode,omitempty" json:"mode,omitempty"`
	
		// Owner 
		Owner string `yaml:"owner,omitempty" json:"owner,omitempty"`
	
		// Size 
		Size int `yaml:"size,omitempty" json:"size,omitempty"`
	
		// Src 
		Src string `yaml:"src,omitempty" json:"src,omitempty"`
	
		// State 
		State string `yaml:"state,omitempty" json:"state,omitempty"`
	
		// Uid 
		Uid int `yaml:"uid,omitempty" json:"uid,omitempty"`
	
}

func (m *Unarchive) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Options, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewUnarchive() *Unarchive {
	return &Unarchive {
		ModuleName: "unarchive",
	}
}
	