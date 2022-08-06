package types

import "encoding/json"

func Parse(raw []byte) (*Spec, error) {
	spec := &Spec{
		Modules: []*Module{},
	}
	err := json.Unmarshal(raw, spec)
	if err != nil {
		return spec, err
	}
	return spec, err
}

type Spec struct {
	Modules []*Module `yaml:"modules" json:"modules"`
}

type Module struct {
	NormalizedName   string             `yaml:"normalizedName" json:"normalizedName"`
	ModuleName       string             `yaml:"module" json:"module"`
	ShortDescription string             `yaml:"short_description" json:"short_description"`
	Description      []string           `yaml:"description" json:"description"`
	Params           map[string]*Param  `yaml:"options" json:"options"`
	Returns          map[string]*Return `yaml:"returns" json:"returns"`
	Path             string             `yaml:"-" json:"-"`
	Documentation    string             `yaml:"doc" json:"doc"`
	Return           string             `yaml:"return" json:"return"`
	SourceLink       string             `yaml:"-" json:"-"`
}

type Param struct {
	NormalizedName string      `yaml:"normalizedName" json:"normalizedName"`
	StructTag      string      `yaml:"structTag" json:"structTag"`
	GoType         string      `yaml:"goType" json:"goType"`
	GoElements     string      `yaml:"goElements" json:"goElements"`
	Description    []string    `yaml:"description" json:"description"`
	Type           string      `default:"str" yaml:"type" json:"type"`
	Required       bool        `default:"no" yaml:"required" json:"required"`
	Default        interface{} `yaml:"default" json:"default"`
	Elements       string      `yaml:"elements" json:"elements"`
	Aliases        []string    `yaml:"aliases" json:"aliases"`
}

type Return struct {
	NormalizedName string      `yaml:"normalizedName" json:"normalizedName"`
	Description    interface{} `yaml:"description" json:"description"`
	Returned       string      `yaml:"returned" json:"returned"`
	// complex type is not supported...
	Type   string      `default:"str" yaml:"type" json:"type"`
	Sample interface{} `yaml:"sample" json:"sample"`

	GoType    string `yaml:"goType" json:"goType"`
	StructTag string `yaml:"structTag" json:"structTag"`
}

// Desc special handling since description can be either a string or a string slice.
func (r *Return) Desc() []string {
	switch t := r.Description.(type) {
	case string:
		return []string{t}
	case []string:
		return t
	}
	return []string{}
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
