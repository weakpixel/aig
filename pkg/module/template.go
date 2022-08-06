package module

// Autogenerated file

import (
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("template", func() Module {
		return NewTemplate()
	})
}

//
// Template (template) - Template a file out to a target host
//
func NewTemplate() *Template {
	return &Template{}
}

// Template (template) - Template a file out to a target host
//
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
	Follow bool `yaml:"follow,omitempty" json:"follow,omitempty"`
}

type TemplateResult struct {
	types.CommonReturn
	Raw string
}

func (m *Template) GetResult() interface{} {
	return &m.Result
}

func (m *Template) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Template) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Template) GetParams() interface{} {
	return &m.Params
}

func (m *Template) GetType() string {
	return "template"
}
