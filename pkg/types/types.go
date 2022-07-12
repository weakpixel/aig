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
