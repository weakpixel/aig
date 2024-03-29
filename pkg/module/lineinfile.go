package module

// Autogenerated file

import (
	"fmt"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("lineinfile", func() types.Module {
		return NewLineinfile()
	})
}

//
// Lineinfile (lineinfile) - Manage lines in text files
//
func NewLineinfile() *Lineinfile {
	module := Lineinfile{}
	// Create dynamic param values
	paramValues := map[string]types.Value{}
	paramValues["backrefs"] = types.NewBoolValue(&module.Params.Backrefs)
	paramValues["backup"] = types.NewBoolValue(&module.Params.Backup)
	paramValues["create"] = types.NewBoolValue(&module.Params.Create)
	paramValues["firstmatch"] = types.NewBoolValue(&module.Params.Firstmatch)
	paramValues["insertafter"] = types.NewStringValue(&module.Params.Insertafter)
	paramValues["insertbefore"] = types.NewStringValue(&module.Params.Insertbefore)
	paramValues["line"] = types.NewStringValue(&module.Params.Line)
	paramValues["others"] = types.NewStringValue(&module.Params.Others)
	paramValues["path"] = types.NewStringValue(&module.Params.Path)
	paramValues["regexp"] = types.NewStringValue(&module.Params.Regexp)
	paramValues["search_string"] = types.NewStringValue(&module.Params.SearchString)
	paramValues["state"] = types.NewStringValue(&module.Params.State)
	module.Params.values = paramValues

	// Create dynamic result values
	resultValues := map[string]types.Value{}

	module.Result.values = resultValues

	return &module
}

// Lineinfile (lineinfile) - Manage lines in text files
//
// This module ensures a particular line is in a file, or replace an existing line using a back-referenced regular expression.
//
// This is primarily useful when you want to change a single line in a file only.
//
// See the M(ansible.builtin.replace) module if you want to change multiple, similar lines or check M(ansible.builtin.blockinfile) if you want to insert/update/remove a block of lines in a file. For other cases, see the M(ansible.builtin.copy) or M(ansible.builtin.template) modules.
//
//
// Source: https://github.com/ansible/ansible/blob/v2.13.1/lib/ansible/modules/lineinfile.py
type Lineinfile struct {
	Params LineinfileParams
	Result LineinfileResult
}

type LineinfileParams struct {

	// Backrefs
	// Used with C(state=present).
	// If set, C(line) can contain backreferences (both positional and named) that will get populated if the C(regexp) matches.
	// This parameter changes the operation of the module slightly; C(insertbefore) and C(insertafter) will be ignored, and if the C(regexp) does not match anywhere in the file, the file will be left unchanged.
	// If the C(regexp) does match, the last matching line will be replaced by the expanded line parameter.
	// Mutually exclusive with C(search_string).
	//
	// Default: no
	// Required: false
	Backrefs bool `yaml:"backrefs,omitempty" json:"backrefs,omitempty" cty:"backrefs"`

	// Backup
	// Create a backup file including the timestamp information so you can get the original file back if you somehow clobbered it incorrectly.
	//
	// Default: no
	// Required: false
	Backup bool `yaml:"backup,omitempty" json:"backup,omitempty" cty:"backup"`

	// Create
	// Used with C(state=present).
	// If specified, the file will be created if it does not already exist.
	// By default it will fail if the file is missing.
	//
	// Default: no
	// Required: false
	Create bool `yaml:"create,omitempty" json:"create,omitempty" cty:"create"`

	// Firstmatch
	// Used with C(insertafter) or C(insertbefore).
	// If set, C(insertafter) and C(insertbefore) will work with the first line that matches the given regular expression.
	//
	// Default: no
	// Required: false
	Firstmatch bool `yaml:"firstmatch,omitempty" json:"firstmatch,omitempty" cty:"firstmatch"`

	// Insertafter
	// Used with C(state=present).
	// If specified, the line will be inserted after the last match of specified regular expression.
	// If the first match is required, use(firstmatch=yes).
	// A special value is available; C(EOF) for inserting the line at the end of the file.
	// If specified regular expression has no matches, EOF will be used instead.
	// If C(insertbefore) is set, default value C(EOF) will be ignored.
	// If regular expressions are passed to both C(regexp) and C(insertafter), C(insertafter) is only honored if no match for C(regexp) is found.
	// May not be used with C(backrefs) or C(insertbefore).
	//
	// Default: EOF
	// Required: false
	Insertafter string `yaml:"insertafter,omitempty" json:"insertafter,omitempty" cty:"insertafter"`

	// Insertbefore
	// Used with C(state=present).
	// If specified, the line will be inserted before the last match of specified regular expression.
	// If the first match is required, use C(firstmatch=yes).
	// A value is available; C(BOF) for inserting the line at the beginning of the file.
	// If specified regular expression has no matches, the line will be inserted at the end of the file.
	// If regular expressions are passed to both C(regexp) and C(insertbefore), C(insertbefore) is only honored if no match for C(regexp) is found.
	// May not be used with C(backrefs) or C(insertafter).
	//
	// Default: <no value>
	// Required: false
	Insertbefore string `yaml:"insertbefore,omitempty" json:"insertbefore,omitempty" cty:"insertbefore"`

	// Line
	// The line to insert/replace into the file.
	// Required for C(state=present).
	// If C(backrefs) is set, may contain backreferences that will get expanded with the C(regexp) capture groups if the regexp matches.
	//
	// Default: <no value>
	// Required: false
	Line string `yaml:"line,omitempty" json:"line,omitempty" cty:"line"`

	// Others
	// All arguments accepted by the M(ansible.builtin.file) module also work here.
	//
	// Default: <no value>
	// Required: false
	Others string `yaml:"others,omitempty" json:"others,omitempty" cty:"others"`

	// Path
	// The file to modify.
	// Before Ansible 2.3 this option was only usable as I(dest), I(destfile) and I(name).
	//
	// Default: <no value>
	// Required: true
	Path string `yaml:"path,omitempty" json:"path,omitempty" cty:"path"`

	// Regexp
	// The regular expression to look for in every line of the file.
	// For C(state=present), the pattern to replace if found. Only the last line found will be replaced.
	// For C(state=absent), the pattern of the line(s) to remove.
	// If the regular expression is not matched, the line will be added to the file in keeping with C(insertbefore) or C(insertafter) settings.
	// When modifying a line the regexp should typically match both the initial state of the line as well as its state after replacement by C(line) to ensure idempotence.
	// Uses Python regular expressions. See U(https://docs.python.org/3/library/re.html).
	//
	// Default: <no value>
	// Required: false
	Regexp string `yaml:"regexp,omitempty" json:"regexp,omitempty" cty:"regexp"`

	// SearchString
	// The literal string to look for in every line of the file. This does not have to match the entire line.
	// For C(state=present), the line to replace if the string is found in the file. Only the last line found will be replaced.
	// For C(state=absent), the line(s) to remove if the string is in the line.
	// If the literal expression is not matched, the line will be added to the file in keeping with C(insertbefore) or C(insertafter) settings.
	// Mutually exclusive with C(backrefs) and C(regexp).
	//
	// Default: <no value>
	// Required: false
	SearchString string `yaml:"search_string,omitempty" json:"search_string,omitempty" cty:"search_string"`

	// State
	// Whether the line should be there or not.
	//
	// Default: present
	// Required: false
	State string `yaml:"state,omitempty" json:"state,omitempty" cty:"state"`

	values map[string]types.Value
}

func (p *LineinfileParams) Names() []string {
	names := []string{}
	for name := range p.values {
		names = append(names, name)
	}
	return names
}

func (p *LineinfileParams) Set(name string, value interface{}) error {
	v, ok := p.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (p *LineinfileParams) Get(name string) (interface{}, error) {
	v, ok := p.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

type LineinfileResult struct {
	types.CommonReturn
	Raw string

	values map[string]types.Value
}

func (r *LineinfileResult) Names() []string {
	names := []string{}
	for name := range r.values {
		names = append(names, name)
	}
	return names
}

func (r *LineinfileResult) Set(name string, value interface{}) error {
	v, ok := r.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (r *LineinfileResult) Get(name string) (interface{}, error) {
	v, ok := r.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

func (m *Lineinfile) GetResult() types.Result {
	return &m.Result
}

func (m *Lineinfile) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Lineinfile) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Lineinfile) GetParams() types.Params {
	return &m.Params
}

func (m *Lineinfile) GetType() string {
	return "lineinfile"
}
