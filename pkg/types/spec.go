package types

import "encoding/json"

func Parse(raw []byte) (*Spec, error) {
	spec := &Spec{
		Modules: []*ModuleSpec{},
	}
	err := json.Unmarshal(raw, spec)
	if err != nil {
		return spec, err
	}
	return spec, err
}

type Spec struct {
	Modules []*ModuleSpec `yaml:"modules" json:"modules"`
}

type ModuleSpec struct {
	NormalizedName   string      `yaml:"normalizedName" json:"normalizedName"`
	ModuleName       string      `yaml:"module" json:"module"`
	ShortDescription string      `yaml:"short_description" json:"short_description"`
	Description      []string    `yaml:"description" json:"description"`
	Params           ParamSpecs  `yaml:"options" json:"options"`
	Returns          ReturnSpecs `yaml:"returns" json:"returns"`
	Path             string      `yaml:"-" json:"-"`
	Documentation    string      `yaml:"doc" json:"doc"`
	Return           string      `yaml:"return" json:"return"`
	SourceLink       string      `yaml:"-" json:"-"`
}

type ParamSpecs map[string]*ParamSpec

func (p *ParamSpecs) Names() []string {
	names := []string{}
	for name, _ := range *p {
		names = append(names, name)
	}
	return names
}

func (p *ParamSpecs) Defaults() map[string]interface{} {
	vals := map[string]interface{}{}
	for name, p := range *p {
		if p.Default == nil || p.Default == "<no value>" || p.Default == "" {
			continue
		}

		vals[name] = p.Default
	}

	return vals
}

type ParamSpec struct {
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

type ReturnSpecs map[string]*ReturnSpec

func (p *ReturnSpecs) Names() []string {
	names := []string{}
	for name, _ := range *p {
		names = append(names, name)
	}
	return names
}

type ReturnSpec struct {
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
func (r *ReturnSpec) Desc() []string {
	switch t := r.Description.(type) {
	case string:
		return []string{t}
	case []string:
		return t
	}
	return []string{}
}
