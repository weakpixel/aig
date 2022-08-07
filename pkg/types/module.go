package types

type Module interface {
	GetResultRaw() string
	GetCommonResult() CommonReturn
	GetParams() Params
	GetResult() Result
	GetType() string
}

type Params interface {
	Names() []string
	Set(name string, value interface{}) error
	Get(name string) (interface{}, error)
}

type Result interface {
	Names() []string
	Set(name string, value interface{}) error
	Get(name string) (interface{}, error)
}

type CommonReturn struct {

	// For those modules that implement backup=no|yes when manipulating files, a path to the backup file created.
	BackupFile string `yaml:"backup_file,omitempty" json:"backup_file,omitempty"`

	// A boolean indicating if the task had to make changes.
	Changed bool `yaml:"changed,omitempty" json:"changed,omitempty"`

	// A boolean that indicates if the task was failed or not.
	Failed bool `yaml:"failed,omitempty" json:"failed,omitempty"`

	// A string with a generic message relayed to the user.
	Msg string `yaml:"msg,omitempty" json:"msg,omitempty"`

	// Some modules execute command line utilities or are geared for executing commands directly (raw, shell, command, etc), this field contains ‘return code’ of these utilities.
	Rc int `yaml:"rc,omitempty" json:"rc,omitempty"`

	// A boolean that indicates if the task was skipped or not
	Skipped bool `yaml:"skipped,omitempty" json:"skipped,omitempty"`

	// Some modules execute command line utilities or are geared for executing commands directly (raw, shell, command, etc), this field contains the error output of these utilities.
	Stderr string `yaml:"stderr,omitempty" json:"stderr,omitempty"`

	// Some modules execute command line utilities or are geared for executing commands directly (raw, shell, command, etc). This field contains the normal output of these utilities.
	Stdout string `yaml:"stdout,omitempty" json:"stdout,omitempty"`

	// Information on how the module was invoked.
	// invocation `yaml:",omitempty" json:",omitempty"`

	// If this key exists, it indicates that a loop was present for the task and that it contains a list of the normal module ‘result’ per item.
	// results `yaml:",omitempty" json:",omitempty"`
}
