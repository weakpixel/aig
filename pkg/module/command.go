package module

// Autogenerated file

import (
	"fmt"

	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("command", func() types.Module {
		return NewCommand()
	})
}

//
// Command (command) - Execute commands on targets
//
func NewCommand() *Command {
	module := Command{}
	// Create dynamic param values
	paramValues := map[string]types.Value{}
	paramValues["argv"] = types.NewStringListValue(&module.Params.Argv)
	paramValues["chdir"] = types.NewStringValue(&module.Params.Chdir)
	paramValues["cmd"] = types.NewStringValue(&module.Params.Cmd)
	paramValues["creates"] = types.NewStringValue(&module.Params.Creates)
	paramValues["free_form"] = types.NewStringValue(&module.Params.FreeForm)
	paramValues["removes"] = types.NewStringValue(&module.Params.Removes)
	paramValues["stdin"] = types.NewStringValue(&module.Params.Stdin)
	paramValues["stdin_add_newline"] = types.NewBoolValue(&module.Params.StdinAddNewline)
	paramValues["strip_empty_ends"] = types.NewBoolValue(&module.Params.StripEmptyEnds)
	paramValues["warn"] = types.NewBoolValue(&module.Params.Warn)
	module.Params.values = paramValues

	// Create dynamic result values
	resultValues := map[string]types.Value{}

	resultValues["cmd"] = types.NewStringListValue(&module.Result.Cmd)
	resultValues["delta"] = types.NewStringValue(&module.Result.Delta)
	resultValues["end"] = types.NewStringValue(&module.Result.End)
	resultValues["msg"] = types.NewBoolValue(&module.Result.Msg)
	resultValues["rc"] = types.NewIntValue(&module.Result.Rc)
	resultValues["start"] = types.NewStringValue(&module.Result.Start)
	resultValues["stderr"] = types.NewStringValue(&module.Result.Stderr)
	resultValues["stderr_lines"] = types.NewStringListValue(&module.Result.StderrLines)
	resultValues["stdout"] = types.NewStringValue(&module.Result.Stdout)
	resultValues["stdout_lines"] = types.NewStringListValue(&module.Result.StdoutLines)
	module.Result.values = resultValues

	return &module
}

// Command (command) - Execute commands on targets
//
// The C(command) module takes the command name followed by a list of space-delimited arguments.
//
// The given command will be executed on all selected nodes.
//
// The command(s) will not be processed through the shell, so variables like C($HOSTNAME) and operations like C("*"), C("<"), C(">"), C("|"), C(";") and C("&") will not work. Use the M(ansible.builtin.shell) module if you need these features.
//
// To create C(command) tasks that are easier to read than the ones using space-delimited arguments, pass parameters using the C(args) L(task keyword,../reference_appendices/playbooks_keywords.html#task) or use C(cmd) parameter.
//
// Either a free form command or C(cmd) parameter is required, see the examples.
//
// For Windows targets, use the M(ansible.windows.win_command) module instead.
//
//
// Source: https://github.com/ansible/ansible/blob/v2.13.1/lib/ansible/modules/command.py
type Command struct {
	Params CommandParams
	Result CommandResult
}

type CommandParams struct {

	// Argv
	// Passes the command as a list rather than a string.
	// Use C(argv) to avoid quoting values that would otherwise be interpreted incorrectly (for example "user name").
	// Only the string (free form) or the list (argv) form can be provided, not both.  One or the other must be provided.
	//
	// Default: <no value>
	// Required: false
	Argv []string `yaml:"argv,omitempty" json:"argv,omitempty" cty:"argv"`

	// Chdir
	// Change into this directory before running the command.
	//
	// Default: <no value>
	// Required: false
	Chdir string `yaml:"chdir,omitempty" json:"chdir,omitempty" cty:"chdir"`

	// Cmd
	// The command to run.
	//
	// Default: <no value>
	// Required: false
	Cmd string `yaml:"cmd,omitempty" json:"cmd,omitempty" cty:"cmd"`

	// Creates
	// A filename or (since 2.0) glob pattern. If a matching file already exists, this step B(will not) be run.
	// This is checked before I(removes) is checked.
	//
	// Default: <no value>
	// Required: false
	Creates string `yaml:"creates,omitempty" json:"creates,omitempty" cty:"creates"`

	// FreeForm
	// The command module takes a free form string as a command to run.
	// There is no actual parameter named 'free form'.
	//
	// Default: <no value>
	// Required: false
	FreeForm string `yaml:"free_form,omitempty" json:"free_form,omitempty" cty:"free_form"`

	// Removes
	// A filename or (since 2.0) glob pattern. If a matching file exists, this step B(will) be run.
	// This is checked after I(creates) is checked.
	//
	// Default: <no value>
	// Required: false
	Removes string `yaml:"removes,omitempty" json:"removes,omitempty" cty:"removes"`

	// Stdin
	// Set the stdin of the command directly to the specified value.
	//
	// Default: <no value>
	// Required: false
	Stdin string `yaml:"stdin,omitempty" json:"stdin,omitempty" cty:"stdin"`

	// StdinAddNewline
	// If set to C(yes), append a newline to stdin data.
	//
	// Default: yes
	// Required: false
	StdinAddNewline bool `yaml:"stdin_add_newline,omitempty" json:"stdin_add_newline,omitempty" cty:"stdin_add_newline"`

	// StripEmptyEnds
	// Strip empty lines from the end of stdout/stderr in result.
	//
	// Default: yes
	// Required: false
	StripEmptyEnds bool `yaml:"strip_empty_ends,omitempty" json:"strip_empty_ends,omitempty" cty:"strip_empty_ends"`

	// Warn
	// (deprecated) Enable or disable task warnings.
	// This feature is deprecated and will be removed in 2.14.
	// As of version 2.11, this option is now disabled by default.
	//
	// Default: no
	// Required: false
	Warn bool `yaml:"warn,omitempty" json:"warn,omitempty" cty:"warn"`

	values map[string]types.Value
}

func (p *CommandParams) Names() []string {
	names := []string{}
	for name := range p.values {
		names = append(names, name)
	}
	return names
}

func (p *CommandParams) Set(name string, value interface{}) error {
	v, ok := p.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (p *CommandParams) Get(name string) (interface{}, error) {
	v, ok := p.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

type CommandResult struct {
	types.CommonReturn
	Raw string

	// Cmd
	// The command executed by the task.
	Cmd []string `yaml:"cmd,omitempty" json:"cmd,omitempty" cty:"cmd"`

	// Delta
	// The command execution delta time.
	Delta string `yaml:"delta,omitempty" json:"delta,omitempty" cty:"delta"`

	// End
	// The command execution end time.
	End string `yaml:"end,omitempty" json:"end,omitempty" cty:"end"`

	// Msg
	// changed
	Msg string `yaml:"msg,omitempty" json:"msg,omitempty" cty:"msg"`

	// Rc
	// The command return code (0 means success).
	Rc int `yaml:"rc,omitempty" json:"rc,omitempty" cty:"rc"`

	// Start
	// The command execution start time.
	Start string `yaml:"start,omitempty" json:"start,omitempty" cty:"start"`

	// Stderr
	// The command standard error.
	Stderr string `yaml:"stderr,omitempty" json:"stderr,omitempty" cty:"stderr"`

	// StderrLines
	// The command standard error split in lines.
	StderrLines []string `yaml:"stderr_lines,omitempty" json:"stderr_lines,omitempty" cty:"stderr_lines"`

	// Stdout
	// The command standard output.
	Stdout string `yaml:"stdout,omitempty" json:"stdout,omitempty" cty:"stdout"`

	// StdoutLines
	// The command standard output split in lines.
	StdoutLines []string `yaml:"stdout_lines,omitempty" json:"stdout_lines,omitempty" cty:"stdout_lines"`

	values map[string]types.Value
}

func (r *CommandResult) Names() []string {
	names := []string{}
	for name := range r.values {
		names = append(names, name)
	}
	return names
}

func (r *CommandResult) Set(name string, value interface{}) error {
	v, ok := r.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (r *CommandResult) Get(name string) (interface{}, error) {
	v, ok := r.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

func (m *Command) GetResult() types.Result {
	return &m.Result
}

func (m *Command) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Command) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Command) GetParams() types.Params {
	return &m.Params
}

func (m *Command) GetType() string {
	return "command"
}
