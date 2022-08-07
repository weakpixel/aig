package module

// Autogenerated file

import (
	"fmt"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("meta", func() types.Module {
		return NewMeta()
	})
}

//
// Meta (meta) - Execute Ansible 'actions'
//
func NewMeta() *Meta {
	module := Meta{}
	// Create dynamic param values
	paramValues := map[string]types.Value{}
	paramValues["free_form"] = types.NewStringValue(&module.Params.FreeForm)
	module.Params.values = paramValues

	// Create dynamic result values
	resultValues := map[string]types.Value{}

	module.Result.values = resultValues

	return &module
}

// Meta (meta) - Execute Ansible 'actions'
//
// Meta tasks are a special kind of task which can influence Ansible internal execution or state.
//
// Meta tasks can be used anywhere within your playbook.
//
// This module is also supported for Windows targets.
//
//
// Source: https://github.com/ansible/ansible/blob/v2.13.1/lib/ansible/modules/meta.py
type Meta struct {
	Params MetaParams
	Result MetaResult
}

type MetaParams struct {

	// FreeForm
	// This module takes a free form command, as a string. There is not an actual option named "free form".  See the examples!
	// C(flush_handlers) makes Ansible run any handler tasks which have thus far been notified. Ansible inserts these tasks internally at certain points to implicitly trigger handler runs (after pre/post tasks, the final role execution, and the main tasks section of your plays).
	// C(refresh_inventory) (added in Ansible 2.0) forces the reload of the inventory, which in the case of dynamic inventory scripts means they will be re-executed. If the dynamic inventory script is using a cache, Ansible cannot know this and has no way of refreshing it (you can disable the cache or, if available for your specific inventory datasource (e.g. aws), you can use the an inventory plugin instead of an inventory script). This is mainly useful when additional hosts are created and users wish to use them instead of using the M(ansible.builtin.add_host) module.
	// C(noop) (added in Ansible 2.0) This literally does 'nothing'. It is mainly used internally and not recommended for general use.
	// C(clear_facts) (added in Ansible 2.1) causes the gathered facts for the hosts specified in the play's list of hosts to be cleared, including the fact cache.
	// C(clear_host_errors) (added in Ansible 2.1) clears the failed state (if any) from hosts specified in the play's list of hosts.
	// C(end_play) (added in Ansible 2.2) causes the play to end without failing the host(s). Note that this affects all hosts.
	// C(reset_connection) (added in Ansible 2.3) interrupts a persistent connection (i.e. ssh + control persist)
	// C(end_host) (added in Ansible 2.8) is a per-host variation of C(end_play). Causes the play to end for the current host without failing it.
	// C(end_batch) (added in Ansible 2.12) causes the current batch (see C(serial)) to end without failing the host(s). Note that with C(serial=0) or undefined this behaves the same as C(end_play).
	//
	// Default: <no value>
	// Required: true
	FreeForm string `yaml:"free_form,omitempty" json:"free_form,omitempty"`

	values map[string]types.Value
}

func (p *MetaParams) Names() []string {
	names := []string{}
	for name := range p.values {
		names = append(names, name)
	}
	return names
}

func (p *MetaParams) Set(name string, value interface{}) error {
	v, ok := p.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (p *MetaParams) Get(name string) (interface{}, error) {
	v, ok := p.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

type MetaResult struct {
	types.CommonReturn
	Raw string

	values map[string]types.Value
}

func (r *MetaResult) Names() []string {
	names := []string{}
	for name := range r.values {
		names = append(names, name)
	}
	return names
}

func (r *MetaResult) Set(name string, value interface{}) error {
	v, ok := r.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (r *MetaResult) Get(name string) (interface{}, error) {
	v, ok := r.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

func (m *Meta) GetResult() types.Result {
	return &m.Result
}

func (m *Meta) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Meta) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Meta) GetParams() types.Params {
	return &m.Params
}

func (m *Meta) GetType() string {
	return "meta"
}
