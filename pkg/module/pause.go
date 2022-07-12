// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
)

func init() {
	addModuleFactory("pause", func() Module {
		return NewPause()
	})
}

type Pause struct {
	ModuleName string
	Params     PauseParams
	Result     PauseResult
}

type PauseParams struct {

	// Echo
	Echo bool `yaml:"echo,omitempty" json:"echo,omitempty"`

	// Minutes
	Minutes string `yaml:"minutes,omitempty" json:"minutes,omitempty"`

	// Prompt
	Prompt string `yaml:"prompt,omitempty" json:"prompt,omitempty"`

	// Seconds
	Seconds string `yaml:"seconds,omitempty" json:"seconds,omitempty"`
}

type PauseResult struct {
	Raw string

	// Delta
	Delta string `yaml:"delta,omitempty" json:"delta,omitempty"`

	// Echo
	Echo bool `yaml:"echo,omitempty" json:"echo,omitempty"`

	// Start
	Start string `yaml:"start,omitempty" json:"start,omitempty"`

	// Stdout
	Stdout string `yaml:"stdout,omitempty" json:"stdout,omitempty"`

	// Stop
	Stop string `yaml:"stop,omitempty" json:"stop,omitempty"`

	// UserInput
	UserInput string `yaml:"user_input,omitempty" json:"user_input,omitempty"`
}

func (m *Pause) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func (m *Pause) GetResult() interface{} {
	return &m.Result
}

func (m *Pause) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Pause) GetParams() interface{} {
	return &m.Params
}

func (m *Pause) GetType() string {
	return m.ModuleName
}

func NewPause() *Pause {
	return &Pause{
		ModuleName: "pause",
	}
}
