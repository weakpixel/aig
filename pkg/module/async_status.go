// Autogenerated
package module

import (
	"aig/pkg/ansible"
)

type AsyncStatus struct {
	ModuleName string
	Params     AsyncStatusParams
	Result     AsyncStatusResult
}

type AsyncStatusParams struct {

	// Jid
	Jid string `yaml:"jid,omitempty" json:"jid,omitempty"`

	// Mode
	Mode string `yaml:"mode,omitempty" json:"mode,omitempty"`
}

type AsyncStatusResult struct {
	Raw string

	// AnsibleJobId
	AnsibleJobId string `yaml:"ansible_job_id,omitempty" json:"ansible_job_id,omitempty"`

	// Erased
	Erased string `yaml:"erased,omitempty" json:"erased,omitempty"`

	// Finished
	Finished int `yaml:"finished,omitempty" json:"finished,omitempty"`

	// Started
	Started int `yaml:"started,omitempty" json:"started,omitempty"`

	// Stderr
	Stderr string `yaml:"stderr,omitempty" json:"stderr,omitempty"`

	// Stdout
	Stdout string `yaml:"stdout,omitempty" json:"stdout,omitempty"`
}

func (m *AsyncStatus) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewAsyncStatus() *AsyncStatus {
	return &AsyncStatus{
		ModuleName: "async_status",
	}
}
