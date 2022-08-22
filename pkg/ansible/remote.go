package ansible

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
	"time"

	"github.com/k0sproject/rig"
	"github.com/k0sproject/rig/log"
	"github.com/weakpixel/aig/pkg/types"
)

// StdLog is a simplistic logger for rig
type muteLog struct {
	log.Logger
}

// Debugf prints a debug level log message
func (l *muteLog) Debugf(t string, args ...interface{}) {
}

// Infof prints an info level log message
func (l *muteLog) Infof(t string, args ...interface{}) {
}

// Errorf prints an error level log message
func (l *muteLog) Errorf(t string, args ...interface{}) {
}

func init() {
	log.Log = &muteLog{}
}

func NewRemote(ssh *rig.SSH) (*Remote, error) {
	r := &Remote{
		Conn: rig.Connection{
			SSH: ssh,
		},
	}
	err := r.Conn.Connect()
	if err != nil {
		return nil, err
	}

	pythonBin, err := r.Conn.ExecOutput("which python3 || which python2 || which python")
	if err != nil {
		return nil, err
	}
	if pythonBin == "" {
		return nil, fmt.Errorf("cannot detct python")
	}
	r.PythonBin = pythonBin

	r.uploadModules()

	return r, nil

}

func (r *Remote) uploadModules() error {
	raw, err := loadModulesChunk()
	if err != nil {
		return err
	}

	f, err := os.CreateTemp("", "ansible_modules.zip")
	if err != nil {
		return err
	}

	rawDecoded, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(f.Name(), rawDecoded, 0600)
	if err != nil {
		return err
	}
	err = r.Conn.Upload(f.Name(), "/tmp/ansible_modules.zip")
	if err != nil {
		return err
	}
	return nil
	// r.Conn.Exec()
}

type Remote struct {
	Conn      rig.Connection
	PythonBin string
}

func newRemotePackage(pythonBin string, module types.Module) (*data, error) {
	now := time.Now()
	d := data{
		Shebang:   "#!" + pythonBin,
		Year:      now.Year(),
		Month:     int(now.Month()),
		Day:       now.Day(),
		Hour:      now.Hour(),
		Minute:    now.Minute(),
		Second:    now.Second(),
		ModuleFqn: "ansible.modules." + module.GetType(),
		Params: argsToString(args{
			ModuleArgs: module.GetParams(),
		}),
		AnsibleModule: "ansible" + module.GetType(),
	}
	return &d, nil
}

func writeRemoteBinTmp(pkg *data) (string, error) {
	f, err := os.CreateTemp("", pkg.AnsibleModule)
	if err != nil {
		return "", err
	}
	defer func() {
		f.Close()
	}()

	t, err := template.New("ansible_exec").Parse(ansibleTemplateRemote)
	if err != nil {
		return "", err
	}
	err = t.Execute(f, pkg)
	if err != nil {
		return "", err
	}
	err = pkg.Write(f)
	if err != nil {
		return "", err
	}
	err = os.Chmod(f.Name(), 0700)
	if err != nil {
		return "", err
	}
	return f.Name(), nil
}
