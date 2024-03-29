package module

// Autogenerated file

import (
	"fmt"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("template", func() types.Module {
		return NewTemplate()
	})
}

//
// Template (template) - Template a file out to a target host
//
func NewTemplate() *Template {
	module := Template{}
	// Create dynamic param values
	paramValues := map[string]types.Value{}
	paramValues["follow"] = types.NewBoolValue(&module.Params.Follow)
	module.Params.values = paramValues

	// Create dynamic result values
	resultValues := map[string]types.Value{}

	module.Result.values = resultValues

	return &module
}

// Template (template) - Template a file out to a target host
//
//
// Source: https://github.com/ansible/ansible/blob/v2.13.1/lib/ansible/modules/template.py
type Template struct {
	Params TemplateParams
	Result TemplateResult
}

type TemplateParams struct {

	// Follow
	// Determine whether symbolic links should be followed.
	// When set to C(yes) symbolic links will be followed, if they exist.
	// When set to C(no) symbolic links will not be followed.
	// Previous to Ansible 2.4, this was hardcoded as C(yes).
	//
	// Default: no
	// Required: false
	Follow bool `yaml:"follow,omitempty" json:"follow,omitempty" cty:"follow"`

	values map[string]types.Value
}

func (p *TemplateParams) Names() []string {
	names := []string{}
	for name := range p.values {
		names = append(names, name)
	}
	return names
}

func (p *TemplateParams) Set(name string, value interface{}) error {
	v, ok := p.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (p *TemplateParams) Get(name string) (interface{}, error) {
	v, ok := p.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

type TemplateResult struct {
	types.CommonReturn
	Raw string

	values map[string]types.Value
}

func (r *TemplateResult) Names() []string {
	names := []string{}
	for name := range r.values {
		names = append(names, name)
	}
	return names
}

func (r *TemplateResult) Set(name string, value interface{}) error {
	v, ok := r.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (r *TemplateResult) Get(name string) (interface{}, error) {
	v, ok := r.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

func (m *Template) GetResult() types.Result {
	return &m.Result
}

func (m *Template) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Template) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Template) GetParams() types.Params {
	return &m.Params
}

func (m *Template) GetType() string {
	return "template"
}
