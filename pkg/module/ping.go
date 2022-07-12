// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
)

func init() {
	addModuleFactory("ping", func() Module {
		return NewPing()
	})
}

type Ping struct {
	ModuleName string
	Params     PingParams
	Result     PingResult
}

type PingParams struct {

	// Data
	Data string `yaml:"data,omitempty" json:"data,omitempty"`
}

type PingResult struct {
	Raw string

	// Ping
	Ping string `yaml:"ping,omitempty" json:"ping,omitempty"`
}

func (m *Ping) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func (m *Ping) GetResult() interface{} {
	return &m.Result
}

func (m *Ping) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Ping) GetParams() interface{} {
	return &m.Params
}

func (m *Ping) GetType() string {
	return m.ModuleName
}

func NewPing() *Ping {
	return &Ping{
		ModuleName: "ping",
	}
}
