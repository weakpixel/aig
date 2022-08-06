package module

// Autogenerated file

import (
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("meta", func() Module {
		return NewMeta()
	})
}

//
// Meta (meta) - Execute Ansible 'actions'
//
func NewMeta() *Meta {
	return &Meta{}
}

// Meta (meta) - Execute Ansible 'actions'
//
// Meta tasks are a special kind of task which can influence Ansible internal execution or state.
// Meta tasks can be used anywhere within your playbook.
// This module is also supported for Windows targets.
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
}

type MetaResult struct {
	types.CommonReturn
	Raw string
}

func (m *Meta) GetResult() interface{} {
	return &m.Result
}

func (m *Meta) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Meta) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Meta) GetParams() interface{} {
	return &m.Params
}

func (m *Meta) GetType() string {
	return "meta"
}
