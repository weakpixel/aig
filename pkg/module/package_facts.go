package module

// Autogenerated file

import (
	"fmt"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("package_facts", func() types.Module {
		return NewPackageFacts()
	})
}

//
// PackageFacts (package_facts) - Package information as facts
//
func NewPackageFacts() *PackageFacts {
	module := PackageFacts{}
	// Create dynamic param values
	paramValues := map[string]types.Value{}
	paramValues["manager"] = types.NewStringArrayValue(&module.Params.Manager)
	paramValues["strategy"] = types.NewStringValue(&module.Params.Strategy)
	module.Params.values = paramValues

	// Create dynamic result values
	resultValues := map[string]types.Value{}

	module.Result.values = resultValues

	return &module
}

// PackageFacts (package_facts) - Package information as facts
//
// Return information about installed packages as facts.
//
//
// Source: https://github.com/ansible/ansible/blob/v2.13.1/lib/ansible/modules/package_facts.py
type PackageFacts struct {
	Params PackageFactsParams
	Result PackageFactsResult
}

type PackageFactsParams struct {

	// Manager
	// The package manager used by the system so we can query the package information.
	// Since 2.8 this is a list and can support multiple package managers per system.
	// The 'portage' and 'pkg' options were added in version 2.8.
	// The 'apk' option was added in version 2.11.
	// The 'pkg_info' option was added in version 2.13.
	//
	// Default: [auto]
	// Required: false
	Manager []string `yaml:"manager,omitempty" json:"manager,omitempty"`

	// Strategy
	// This option controls how the module queries the package managers on the system. C(first) means it will return only information for the first supported package manager available. C(all) will return information for all supported and available package managers on the system.
	//
	// Default: first
	// Required: false
	Strategy string `yaml:"strategy,omitempty" json:"strategy,omitempty"`

	values map[string]types.Value
}

func (p *PackageFactsParams) Names() []string {
	names := []string{}
	for name := range p.values {
		names = append(names, name)
	}
	return []string{}
}

func (p *PackageFactsParams) Set(name string, value interface{}) error {
	v, ok := p.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (p *PackageFactsParams) Get(name string) (interface{}, error) {
	v, ok := p.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

type PackageFactsResult struct {
	types.CommonReturn
	Raw string

	// AnsibleFacts
	// Facts to add to ansible_facts.
	AnsibleFacts interface{} `yaml:"ansible_facts,omitempty" json:"ansible_facts,omitempty"`

	values map[string]types.Value
}

func (r *PackageFactsResult) Names() []string {
	names := []string{}
	for name := range r.values {
		names = append(names, name)
	}
	return []string{}
}

func (r *PackageFactsResult) Set(name string, value interface{}) error {
	v, ok := r.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (r *PackageFactsResult) Get(name string) (interface{}, error) {
	v, ok := r.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

func (m *PackageFacts) GetResult() types.Result {
	return &m.Result
}

func (m *PackageFacts) GetResultRaw() string {
	return m.Result.Raw
}

func (m *PackageFacts) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *PackageFacts) GetParams() types.Params {
	return &m.Params
}

func (m *PackageFacts) GetType() string {
	return "package_facts"
}
