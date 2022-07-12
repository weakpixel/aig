// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
)

type Command struct {
	ModuleName string
	Params     CommandParams
	Result     CommandResult
}

type CommandParams struct {

	// Argv
	Argv []string `yaml:"argv,omitempty" json:"argv,omitempty"`

	// Chdir
	Chdir string `yaml:"chdir,omitempty" json:"chdir,omitempty"`

	// Cmd
	Cmd string `yaml:"cmd,omitempty" json:"cmd,omitempty"`

	// Creates
	Creates string `yaml:"creates,omitempty" json:"creates,omitempty"`

	// FreeForm
	FreeForm string `yaml:"free_form,omitempty" json:"free_form,omitempty"`

	// Removes
	Removes string `yaml:"removes,omitempty" json:"removes,omitempty"`

	// Stdin
	Stdin string `yaml:"stdin,omitempty" json:"stdin,omitempty"`

	// StdinAddNewline
	StdinAddNewline bool `yaml:"stdin_add_newline,omitempty" json:"stdin_add_newline,omitempty"`

	// StripEmptyEnds
	StripEmptyEnds bool `yaml:"strip_empty_ends,omitempty" json:"strip_empty_ends,omitempty"`

	// Warn
	Warn bool `yaml:"warn,omitempty" json:"warn,omitempty"`
}

type CommandResult struct {
	Raw string

	// Cmd
	Cmd []map[string]interface{} `yaml:"cmd,omitempty" json:"cmd,omitempty"`

	// Delta
	Delta string `yaml:"delta,omitempty" json:"delta,omitempty"`

	// End
	End string `yaml:"end,omitempty" json:"end,omitempty"`

	// Msg
	Msg bool `yaml:"msg,omitempty" json:"msg,omitempty"`

	// Rc
	Rc int `yaml:"rc,omitempty" json:"rc,omitempty"`

	// Start
	Start string `yaml:"start,omitempty" json:"start,omitempty"`

	// Stderr
	Stderr string `yaml:"stderr,omitempty" json:"stderr,omitempty"`

	// StderrLines
	StderrLines []map[string]interface{} `yaml:"stderr_lines,omitempty" json:"stderr_lines,omitempty"`

	// Stdout
	Stdout string `yaml:"stdout,omitempty" json:"stdout,omitempty"`

	// StdoutLines
	StdoutLines []map[string]interface{} `yaml:"stdout_lines,omitempty" json:"stdout_lines,omitempty"`
}

func (m *Command) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func NewCommand() *Command {
	return &Command{
		ModuleName: "command",
	}
}
