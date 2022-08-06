package module

// Autogenerated file

import (
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("dpkg_selections", func() Module {
		return NewDpkgSelections()
	})
}

//
// DpkgSelections (dpkg_selections) - Dpkg package selection selections
//
func NewDpkgSelections() *DpkgSelections {
	return &DpkgSelections{}
}

// DpkgSelections (dpkg_selections) - Dpkg package selection selections
//
// Change dpkg package selection state via --get-selections and --set-selections.
type DpkgSelections struct {
	Params DpkgSelectionsParams
	Result DpkgSelectionsResult
}

type DpkgSelectionsParams struct {

	// Name
	// Name of the package.
	//
	// Default: <no value>
	// Required: true
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// Selection
	// The selection state to set the package to.
	//
	// Default: <no value>
	// Required: true
	Selection string `yaml:"selection,omitempty" json:"selection,omitempty"`
}

type DpkgSelectionsResult struct {
	types.CommonReturn
	Raw string
}

func (m *DpkgSelections) GetResult() interface{} {
	return &m.Result
}

func (m *DpkgSelections) GetResultRaw() string {
	return m.Result.Raw
}

func (m *DpkgSelections) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *DpkgSelections) GetParams() interface{} {
	return &m.Params
}

func (m *DpkgSelections) GetType() string {
	return "dpkg_selections"
}
