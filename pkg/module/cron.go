package module

// Autogenerated file

import (
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("cron", func() Module {
		return NewCron()
	})
}

//
// Cron (cron) - Manage cron.d and crontab entries
//
func NewCron() *Cron {
	return &Cron{}
}

// Cron (cron) - Manage cron.d and crontab entries
//
// Use this module to manage crontab and environment variables entries. This module allows you to create environment variables and named crontab entries, update, or delete them.
// When crontab jobs are managed: the module includes one line with the description of the crontab entry C("#Ansible: <name>") corresponding to the "name" passed to the module, which is used by future ansible/module calls to find/check the state. The "name" parameter should be unique, and changing the "name" value will result in a new cron task being created (or a different one being removed).
// When environment variables are managed, no comment line is added, but, when the module needs to find/check the state, it uses the "name" parameter to find the environment variable definition line.
// When using symbols such as %, they must be properly escaped.
//
// Source: https://github.com/ansible/ansible/blob/v2.13.1/lib/ansible/modules/cron.py
type Cron struct {
	Params CronParams
	Result CronResult
}

type CronParams struct {

	// Backup
	// If set, create a backup of the crontab before it is modified. The location of the backup is returned in the C(backup_file) variable by this module.
	//
	// Default: no
	// Required: false
	Backup bool `yaml:"backup,omitempty" json:"backup,omitempty"`

	// CronFile
	// If specified, uses this file instead of an individual user's crontab. The assumption is that this file is exclusively managed by the module, do not use if the file contains multiple entries, NEVER use for /etc/crontab.
	// If this is a relative path, it is interpreted with respect to I(/etc/cron.d).
	// Many linux distros expect (and some require) the filename portion to consist solely of upper- and lower-case letters, digits, underscores, and hyphens.
	// Using this parameter requires you to specify the I(user) as well, unless I(state) is not I(present).
	// Either this parameter or I(name) is required
	//
	// Default: <no value>
	// Required: false
	CronFile string `yaml:"cron_file,omitempty" json:"cron_file,omitempty"`

	// Day
	// Day of the month the job should run (C(1-31), C(*), C(*/2), and so on).
	//
	// Default: *
	// Required: false
	Day string `yaml:"day,omitempty" json:"day,omitempty"`

	// Disabled
	// If the job should be disabled (commented out) in the crontab.
	// Only has effect if I(state=present).
	//
	// Default: no
	// Required: false
	Disabled bool `yaml:"disabled,omitempty" json:"disabled,omitempty"`

	// Env
	// If set, manages a crontab's environment variable.
	// New variables are added on top of crontab.
	// I(name) and I(value) parameters are the name and the value of environment variable.
	//
	// Default: false
	// Required: false
	Env bool `yaml:"env,omitempty" json:"env,omitempty"`

	// Hour
	// Hour when the job should run (C(0-23), C(*), C(*/2), and so on).
	//
	// Default: *
	// Required: false
	Hour string `yaml:"hour,omitempty" json:"hour,omitempty"`

	// Insertafter
	// Used with I(state=present) and I(env).
	// If specified, the environment variable will be inserted after the declaration of specified environment variable.
	//
	// Default: <no value>
	// Required: false
	Insertafter string `yaml:"insertafter,omitempty" json:"insertafter,omitempty"`

	// Insertbefore
	// Used with I(state=present) and I(env).
	// If specified, the environment variable will be inserted before the declaration of specified environment variable.
	//
	// Default: <no value>
	// Required: false
	Insertbefore string `yaml:"insertbefore,omitempty" json:"insertbefore,omitempty"`

	// Job
	// The command to execute or, if env is set, the value of environment variable.
	// The command should not contain line breaks.
	// Required if I(state=present).
	//
	// Default: <no value>
	// Required: false
	Job string `yaml:"job,omitempty" json:"job,omitempty"`

	// Minute
	// Minute when the job should run (C(0-59), C(*), C(*/2), and so on).
	//
	// Default: *
	// Required: false
	Minute string `yaml:"minute,omitempty" json:"minute,omitempty"`

	// Month
	// Month of the year the job should run (C(1-12), C(*), C(*/2), and so on).
	//
	// Default: *
	// Required: false
	Month string `yaml:"month,omitempty" json:"month,omitempty"`

	// Name
	// Description of a crontab entry or, if env is set, the name of environment variable.
	// This parameter is always required as of ansible-core 2.12.
	//
	// Default: <no value>
	// Required: true
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// SpecialTime
	// Special time specification nickname.
	//
	// Default: <no value>
	// Required: false
	SpecialTime string `yaml:"special_time,omitempty" json:"special_time,omitempty"`

	// State
	// Whether to ensure the job or environment variable is present or absent.
	//
	// Default: present
	// Required: false
	State string `yaml:"state,omitempty" json:"state,omitempty"`

	// User
	// The specific user whose crontab should be modified.
	// When unset, this parameter defaults to the current user.
	//
	// Default: <no value>
	// Required: false
	User string `yaml:"user,omitempty" json:"user,omitempty"`

	// Weekday
	// Day of the week that the job should run (C(0-6) for Sunday-Saturday, C(*), and so on).
	//
	// Default: *
	// Required: false
	Weekday string `yaml:"weekday,omitempty" json:"weekday,omitempty"`
}

type CronResult struct {
	types.CommonReturn
	Raw string
}

func (m *Cron) GetResult() interface{} {
	return &m.Result
}

func (m *Cron) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Cron) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Cron) GetParams() interface{} {
	return &m.Params
}

func (m *Cron) GetType() string {
	return "cron"
}
