package module

// Autogenerated file

import (
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("package", func() Module {
		return NewPackage()
	})
}

//
// Package (package) - Generic OS package manager
//
func NewPackage() *Package {
	return &Package{}
}

// Package (package) - Generic OS package manager
//
// This modules manages packages on a target without specifying a package manager module (like M(ansible.builtin.yum), M(ansible.builtin.apt), ...). It is convenient to use in an heterogeneous environment of machines without having to create a specific task for each package manager. C(package) calls behind the module for the package manager used by the operating system discovered by the module M(ansible.builtin.setup).  If C(setup) was not yet run, C(package) will run it.
// This module acts as a proxy to the underlying package manager module. While all arguments will be passed to the underlying module, not all modules support the same arguments. This documentation only covers the minimum intersection of module arguments that all packaging modules support.
// For Windows targets, use the M(ansible.windows.win_package) module instead.
type Package struct {
	Params PackageParams
	Result PackageResult
}

type PackageParams struct {

	// Name
	// Package name, or package specifier with version.
	// Syntax varies with package manager. For example C(name-1.0) or C(name=1.0).
	// Package names also vary with package manager; this module will not "translate" them per distro. For example C(libyaml-dev), C(libyaml-devel).
	//
	// Default: <no value>
	// Required: true
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// State
	// Whether to install (C(present)), or remove (C(absent)) a package.
	// You can use other states like C(latest) ONLY if they are supported by the underlying package module(s) executed.
	//
	// Default: <no value>
	// Required: true
	State string `yaml:"state,omitempty" json:"state,omitempty"`

	// Use
	// The required package manager module to use (C(yum), C(apt), and so on). The default 'auto' will use existing facts or try to autodetect it.
	// You should only use this field if the automatic selection is not working for some reason.
	//
	// Default: auto
	// Required: false
	Use string `yaml:"use,omitempty" json:"use,omitempty"`
}

type PackageResult struct {
	types.CommonReturn
	Raw string
}

func (m *Package) GetResult() interface{} {
	return &m.Result
}

func (m *Package) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Package) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Package) GetParams() interface{} {
	return &m.Params
}

func (m *Package) GetType() string {
	return "package"
}
