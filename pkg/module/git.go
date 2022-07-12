// Autogenerated
package module

import (
	"github.com/weakpixel/aig/pkg/ansible"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("git", func() Module {
		return NewGit()
	})
}

type Git struct {
	ModuleName string
	Params     GitParams
	Result     GitResult
}

type GitParams struct {

	// AcceptHostkey
	AcceptHostkey bool `yaml:"accept_hostkey,omitempty" json:"accept_hostkey,omitempty"`

	// AcceptNewhostkey
	AcceptNewhostkey bool `yaml:"accept_newhostkey,omitempty" json:"accept_newhostkey,omitempty"`

	// Archive
	Archive string `yaml:"archive,omitempty" json:"archive,omitempty"`

	// ArchivePrefix
	ArchivePrefix string `yaml:"archive_prefix,omitempty" json:"archive_prefix,omitempty"`

	// Bare
	Bare bool `yaml:"bare,omitempty" json:"bare,omitempty"`

	// Clone
	Clone bool `yaml:"clone,omitempty" json:"clone,omitempty"`

	// Depth
	Depth int `yaml:"depth,omitempty" json:"depth,omitempty"`

	// Dest
	Dest string `yaml:"dest,omitempty" json:"dest,omitempty"`

	// Executable
	Executable string `yaml:"executable,omitempty" json:"executable,omitempty"`

	// Force
	Force bool `yaml:"force,omitempty" json:"force,omitempty"`

	// GpgWhitelist
	GpgWhitelist []string `yaml:"gpg_whitelist,omitempty" json:"gpg_whitelist,omitempty"`

	// KeyFile
	KeyFile string `yaml:"key_file,omitempty" json:"key_file,omitempty"`

	// Recursive
	Recursive bool `yaml:"recursive,omitempty" json:"recursive,omitempty"`

	// Reference
	Reference string `yaml:"reference,omitempty" json:"reference,omitempty"`

	// Refspec
	Refspec string `yaml:"refspec,omitempty" json:"refspec,omitempty"`

	// Remote
	Remote string `yaml:"remote,omitempty" json:"remote,omitempty"`

	// Repo
	Repo string `yaml:"repo,omitempty" json:"repo,omitempty"`

	// SeparateGitDir
	SeparateGitDir string `yaml:"separate_git_dir,omitempty" json:"separate_git_dir,omitempty"`

	// SingleBranch
	SingleBranch bool `yaml:"single_branch,omitempty" json:"single_branch,omitempty"`

	// SshOpts
	SshOpts string `yaml:"ssh_opts,omitempty" json:"ssh_opts,omitempty"`

	// TrackSubmodules
	TrackSubmodules bool `yaml:"track_submodules,omitempty" json:"track_submodules,omitempty"`

	// Umask
	Umask string `yaml:"umask,omitempty" json:"umask,omitempty"`

	// Update
	Update bool `yaml:"update,omitempty" json:"update,omitempty"`

	// VerifyCommit
	VerifyCommit bool `yaml:"verify_commit,omitempty" json:"verify_commit,omitempty"`

	// Version
	Version string `yaml:"version,omitempty" json:"version,omitempty"`
}

type GitResult struct {
	types.CommonReturn
	Raw string

	// After
	After string `yaml:"after,omitempty" json:"after,omitempty"`

	// Before
	Before string `yaml:"before,omitempty" json:"before,omitempty"`

	// GitDirBefore
	GitDirBefore string `yaml:"git_dir_before,omitempty" json:"git_dir_before,omitempty"`

	// GitDirNow
	GitDirNow string `yaml:"git_dir_now,omitempty" json:"git_dir_now,omitempty"`

	// RemoteUrlChanged
	RemoteUrlChanged bool `yaml:"remote_url_changed,omitempty" json:"remote_url_changed,omitempty"`

	// Warnings
	Warnings string `yaml:"warnings,omitempty" json:"warnings,omitempty"`
}

func (m *Git) Run() error {
	raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
	m.Result.Raw = raw
	return err
}

func (m *Git) GetResult() interface{} {
	return &m.Result
}

func (m *Git) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Git) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Git) GetParams() interface{} {
	return &m.Params
}

func (m *Git) GetType() string {
	return m.ModuleName
}

func NewGit() *Git {
	return &Git{
		ModuleName: "git",
	}
}
