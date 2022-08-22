package module

// Autogenerated file

import (
	"fmt"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("apt", func() types.Module {
		return NewApt()
	})
}

//
// Apt (apt) - Manages apt-packages
//
func NewApt() *Apt {
	module := Apt{}
	// Create dynamic param values
	paramValues := map[string]types.Value{}
	paramValues["allow_change_held_packages"] = types.NewBoolValue(&module.Params.AllowChangeHeldPackages)
	paramValues["allow_downgrade"] = types.NewBoolValue(&module.Params.AllowDowngrade)
	paramValues["allow_unauthenticated"] = types.NewBoolValue(&module.Params.AllowUnauthenticated)
	paramValues["autoclean"] = types.NewBoolValue(&module.Params.Autoclean)
	paramValues["autoremove"] = types.NewBoolValue(&module.Params.Autoremove)
	paramValues["cache_valid_time"] = types.NewIntValue(&module.Params.CacheValidTime)
	paramValues["clean"] = types.NewBoolValue(&module.Params.Clean)
	paramValues["deb"] = types.NewStringValue(&module.Params.Deb)
	paramValues["default_release"] = types.NewStringValue(&module.Params.DefaultRelease)
	paramValues["dpkg_options"] = types.NewStringValue(&module.Params.DpkgOptions)
	paramValues["fail_on_autoremove"] = types.NewBoolValue(&module.Params.FailOnAutoremove)
	paramValues["force"] = types.NewBoolValue(&module.Params.Force)
	paramValues["force_apt_get"] = types.NewBoolValue(&module.Params.ForceAptGet)
	paramValues["install_recommends"] = types.NewBoolValue(&module.Params.InstallRecommends)
	paramValues["lock_timeout"] = types.NewIntValue(&module.Params.LockTimeout)
	paramValues["name"] = types.NewStringListValue(&module.Params.Name)
	paramValues["only_upgrade"] = types.NewBoolValue(&module.Params.OnlyUpgrade)
	paramValues["policy_rc_d"] = types.NewIntValue(&module.Params.PolicyRcD)
	paramValues["purge"] = types.NewBoolValue(&module.Params.Purge)
	paramValues["state"] = types.NewStringValue(&module.Params.State)
	paramValues["update_cache"] = types.NewBoolValue(&module.Params.UpdateCache)
	paramValues["update_cache_retries"] = types.NewIntValue(&module.Params.UpdateCacheRetries)
	paramValues["update_cache_retry_max_delay"] = types.NewIntValue(&module.Params.UpdateCacheRetryMaxDelay)
	paramValues["upgrade"] = types.NewStringValue(&module.Params.Upgrade)
	module.Params.values = paramValues

	// Create dynamic result values
	resultValues := map[string]types.Value{}

	resultValues["cache_update_time"] = types.NewIntValue(&module.Result.CacheUpdateTime)
	resultValues["cache_updated"] = types.NewBoolValue(&module.Result.CacheUpdated)
	resultValues["stderr"] = types.NewStringValue(&module.Result.Stderr)
	resultValues["stdout"] = types.NewStringValue(&module.Result.Stdout)
	module.Result.values = resultValues

	return &module
}

// Apt (apt) - Manages apt-packages
//
// Manages I(apt) packages (such as for Debian/Ubuntu).
//
//
// Source: https://github.com/ansible/ansible/blob/v2.13.1/lib/ansible/modules/apt.py
type Apt struct {
	Params AptParams
	Result AptResult
}

type AptParams struct {

	// AllowChangeHeldPackages
	// Allows changing the version of a package which is on the apt hold list
	//
	// Default: no
	// Required: false
	AllowChangeHeldPackages bool `yaml:"allow_change_held_packages,omitempty" json:"allow_change_held_packages,omitempty" cty:"allow_change_held_packages"`

	// AllowDowngrade
	// Corresponds to the C(--allow-downgrades) option for I(apt).
	// This option enables the named package and version to replace an already installed higher version of that package.
	// Note that setting I(allow_downgrade=true) can make this module behave in a non-idempotent way.
	// (The task could end up with a set of packages that does not match the complete list of specified packages to install).
	//
	// Default: no
	// Required: false
	AllowDowngrade bool `yaml:"allow_downgrade,omitempty" json:"allow_downgrade,omitempty" cty:"allow_downgrade"`

	// AllowUnauthenticated
	// Ignore if packages cannot be authenticated. This is useful for bootstrapping environments that manage their own apt-key setup.
	// C(allow_unauthenticated) is only supported with state: I(install)/I(present)
	//
	// Default: no
	// Required: false
	AllowUnauthenticated bool `yaml:"allow_unauthenticated,omitempty" json:"allow_unauthenticated,omitempty" cty:"allow_unauthenticated"`

	// Autoclean
	// If C(yes), cleans the local repository of retrieved package files that can no longer be downloaded.
	//
	// Default: no
	// Required: false
	Autoclean bool `yaml:"autoclean,omitempty" json:"autoclean,omitempty" cty:"autoclean"`

	// Autoremove
	// If C(yes), remove unused dependency packages for all module states except I(build-dep). It can also be used as the only option.
	// Previous to version 2.4, autoclean was also an alias for autoremove, now it is its own separate command. See documentation for further information.
	//
	// Default: no
	// Required: false
	Autoremove bool `yaml:"autoremove,omitempty" json:"autoremove,omitempty" cty:"autoremove"`

	// CacheValidTime
	// Update the apt cache if it is older than the I(cache_valid_time). This option is set in seconds.
	// As of Ansible 2.4, if explicitly set, this sets I(update_cache=yes).
	//
	// Default: 0
	// Required: false
	CacheValidTime int `yaml:"cache_valid_time,omitempty" json:"cache_valid_time,omitempty" cty:"cache_valid_time"`

	// Clean
	// Run the equivalent of C(apt-get clean) to clear out the local repository of retrieved package files. It removes everything but the lock file from /var/cache/apt/archives/ and /var/cache/apt/archives/partial/.
	// Can be run as part of the package installation (clean runs before install) or as a separate step.
	//
	// Default: no
	// Required: false
	Clean bool `yaml:"clean,omitempty" json:"clean,omitempty" cty:"clean"`

	// Deb
	// Path to a .deb package on the remote machine.
	// If :// in the path, ansible will attempt to download deb before installing. (Version added 2.1)
	// Requires the C(xz-utils) package to extract the control file of the deb package to install.
	//
	// Default: <no value>
	// Required: false
	Deb string `yaml:"deb,omitempty" json:"deb,omitempty" cty:"deb"`

	// DefaultRelease
	// Corresponds to the C(-t) option for I(apt) and sets pin priorities
	//
	// Default: <no value>
	// Required: false
	DefaultRelease string `yaml:"default_release,omitempty" json:"default_release,omitempty" cty:"default_release"`

	// DpkgOptions
	// Add dpkg options to apt command. Defaults to '-o "Dpkg::Options::=--force-confdef" -o "Dpkg::Options::=--force-confold"'
	// Options should be supplied as comma separated list
	//
	// Default: force-confdef,force-confold
	// Required: false
	DpkgOptions string `yaml:"dpkg_options,omitempty" json:"dpkg_options,omitempty" cty:"dpkg_options"`

	// FailOnAutoremove
	// Corresponds to the C(--no-remove) option for C(apt).
	// If C(yes), it is ensured that no packages will be removed or the task will fail.
	// C(fail_on_autoremove) is only supported with state except C(absent)
	//
	// Default: no
	// Required: false
	FailOnAutoremove bool `yaml:"fail_on_autoremove,omitempty" json:"fail_on_autoremove,omitempty" cty:"fail_on_autoremove"`

	// Force
	// Corresponds to the C(--force-yes) to I(apt-get) and implies C(allow_unauthenticated: yes) and C(allow_downgrade: yes)
	// This option will disable checking both the packages' signatures and the certificates of the web servers they are downloaded from.
	// This option *is not* the equivalent of passing the C(-f) flag to I(apt-get) on the command line
	// **This is a destructive operation with the potential to destroy your system, and it should almost never be used.** Please also see C(man apt-get) for more information.
	//
	// Default: no
	// Required: false
	Force bool `yaml:"force,omitempty" json:"force,omitempty" cty:"force"`

	// ForceAptGet
	// Force usage of apt-get instead of aptitude
	//
	// Default: no
	// Required: false
	ForceAptGet bool `yaml:"force_apt_get,omitempty" json:"force_apt_get,omitempty" cty:"force_apt_get"`

	// InstallRecommends
	// Corresponds to the C(--no-install-recommends) option for I(apt). C(yes) installs recommended packages.  C(no) does not install recommended packages. By default, Ansible will use the same defaults as the operating system. Suggested packages are never installed.
	//
	// Default: <no value>
	// Required: false
	InstallRecommends bool `yaml:"install_recommends,omitempty" json:"install_recommends,omitempty" cty:"install_recommends"`

	// LockTimeout
	// How many seconds will this action wait to acquire a lock on the apt db.
	// Sometimes there is a transitory lock and this will retry at least until timeout is hit.
	//
	// Default: 60
	// Required: false
	LockTimeout int `yaml:"lock_timeout,omitempty" json:"lock_timeout,omitempty" cty:"lock_timeout"`

	// Name
	// A list of package names, like C(foo), or package specifier with version, like C(foo=1.0) or C(foo>=1.0). Name wildcards (fnmatch) like C(apt*) and version wildcards like C(foo=1.0*) are also supported.
	//
	// Default: <no value>
	// Required: false
	Name []string `yaml:"name,omitempty" json:"name,omitempty" cty:"name"`

	// OnlyUpgrade
	// Only upgrade a package if it is already installed.
	//
	// Default: no
	// Required: false
	OnlyUpgrade bool `yaml:"only_upgrade,omitempty" json:"only_upgrade,omitempty" cty:"only_upgrade"`

	// PolicyRcD
	// Force the exit code of /usr/sbin/policy-rc.d.
	// For example, if I(policy_rc_d=101) the installed package will not trigger a service start.
	// If /usr/sbin/policy-rc.d already exists, it is backed up and restored after the package installation.
	// If C(null), the /usr/sbin/policy-rc.d isn't created/changed.
	//
	// Default: <no value>
	// Required: false
	PolicyRcD int `yaml:"policy_rc_d,omitempty" json:"policy_rc_d,omitempty" cty:"policy_rc_d"`

	// Purge
	// Will force purging of configuration files if the module state is set to I(absent).
	//
	// Default: no
	// Required: false
	Purge bool `yaml:"purge,omitempty" json:"purge,omitempty" cty:"purge"`

	// State
	// Indicates the desired package state. C(latest) ensures that the latest version is installed. C(build-dep) ensures the package build dependencies are installed. C(fixed) attempt to correct a system with broken dependencies in place.
	//
	// Default: present
	// Required: false
	State string `yaml:"state,omitempty" json:"state,omitempty" cty:"state"`

	// UpdateCache
	// Run the equivalent of C(apt-get update) before the operation. Can be run as part of the package installation or as a separate step.
	// Default is not to update the cache.
	//
	// Default: <no value>
	// Required: false
	UpdateCache bool `yaml:"update_cache,omitempty" json:"update_cache,omitempty" cty:"update_cache"`

	// UpdateCacheRetries
	// Amount of retries if the cache update fails. Also see I(update_cache_retry_max_delay).
	//
	// Default: 5
	// Required: false
	UpdateCacheRetries int `yaml:"update_cache_retries,omitempty" json:"update_cache_retries,omitempty" cty:"update_cache_retries"`

	// UpdateCacheRetryMaxDelay
	// Use an exponential backoff delay for each retry (see I(update_cache_retries)) up to this max delay in seconds.
	//
	// Default: 12
	// Required: false
	UpdateCacheRetryMaxDelay int `yaml:"update_cache_retry_max_delay,omitempty" json:"update_cache_retry_max_delay,omitempty" cty:"update_cache_retry_max_delay"`

	// Upgrade
	// If yes or safe, performs an aptitude safe-upgrade.
	// If full, performs an aptitude full-upgrade.
	// If dist, performs an apt-get dist-upgrade.
	// Note: This does not upgrade a specific package, use state=latest for that.
	// Note: Since 2.4, apt-get is used as a fall-back if aptitude is not present.
	//
	// Default: no
	// Required: false
	Upgrade string `yaml:"upgrade,omitempty" json:"upgrade,omitempty" cty:"upgrade"`

	values map[string]types.Value
}

func (p *AptParams) Names() []string {
	names := []string{}
	for name := range p.values {
		names = append(names, name)
	}
	return names
}

func (p *AptParams) Set(name string, value interface{}) error {
	v, ok := p.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (p *AptParams) Get(name string) (interface{}, error) {
	v, ok := p.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

type AptResult struct {
	types.CommonReturn
	Raw string

	// CacheUpdateTime
	// time of the last cache update (0 if unknown)
	CacheUpdateTime int `yaml:"cache_update_time,omitempty" json:"cache_update_time,omitempty" cty:"cache_update_time"`

	// CacheUpdated
	// if the cache was updated or not
	CacheUpdated bool `yaml:"cache_updated,omitempty" json:"cache_updated,omitempty" cty:"cache_updated"`

	// Stderr
	// error output from apt
	Stderr string `yaml:"stderr,omitempty" json:"stderr,omitempty" cty:"stderr"`

	// Stdout
	// output from apt
	Stdout string `yaml:"stdout,omitempty" json:"stdout,omitempty" cty:"stdout"`

	values map[string]types.Value
}

func (r *AptResult) Names() []string {
	names := []string{}
	for name := range r.values {
		names = append(names, name)
	}
	return names
}

func (r *AptResult) Set(name string, value interface{}) error {
	v, ok := r.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (r *AptResult) Get(name string) (interface{}, error) {
	v, ok := r.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

func (m *Apt) GetResult() types.Result {
	return &m.Result
}

func (m *Apt) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Apt) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Apt) GetParams() types.Params {
	return &m.Params
}

func (m *Apt) GetType() string {
	return "apt"
}
