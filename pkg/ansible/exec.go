package ansible

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"text/template"
	"time"

	"github.com/weakpixel/aig/pkg/module/src"
	"github.com/weakpixel/aig/pkg/types"
)

type data struct {
	Shebang       string
	ZipData       string
	Year          int
	Month         int
	Day           int
	Hour          int
	Minute        int
	Second        int
	ModuleFqn     string
	Params        string
	AnsibleModule string
}

func (d *data) Write(w io.Writer) error {
	t, err := template.New("ansible_exec").Parse(ansibleTemplate)
	if err != nil {
		return err
	}
	return t.Execute(w, d)
}

type args struct {
	ModuleArgs interface{} `json:"ANSIBLE_MODULE_ARGS"`
}

func argsToString(a args) string {
	val, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(val)
}

func loadModulesChunk() (string, error) {
	return src.ModuleSources, nil
}

func newPackage(pythonBin string, module types.Module) (*data, error) {
	raw, err := loadModulesChunk()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	d := data{
		Shebang:   "#!" + pythonBin,
		ZipData:   string(raw),
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

func lookupPython() (string, error) {
	bin, err := exec.LookPath("python")
	if bin == "" || err != nil {
		bin, err = exec.LookPath("python3")
		if err != nil {
			return "", err
		}
	}
	return bin, nil
}

func writeBinTmp(pkg *data) (string, error) {
	f, err := os.CreateTemp("", pkg.AnsibleModule)
	if err != nil {
		return "", err
	}
	defer func() {
		f.Close()
	}()
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

func ExecuteLocal(module types.Module) (string, error) {
	bin, err := lookupPython()
	if err != nil {
		return "", err
	}
	pkg, err := newPackage(bin, module)
	if err != nil {
		return "", err
	}

	pkgBin, err := writeBinTmp(pkg)
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(pkgBin)

	var buf bytes.Buffer
	cmd := exec.Command(pkgBin)
	cmd.Stderr = os.Stdout
	cmd.Stdout = &buf
	err = cmd.Run()

	rawResult := buf.String()
	decodeErr := json.Unmarshal([]byte(rawResult), module.GetResult())
	if decodeErr != nil {
		// if decoding failed than we can assume that the execution failed.
		// return original error
		return rawResult, fmt.Errorf("execution failed: %s  Decoding error: %s", err, decodeErr)
	}
	return rawResult, nil

}

func ExecuteRemote(r *Remote, module types.Module) (string, error) {
	pkg, err := newRemotePackage(r.PythonBin, module)
	if err != nil {
		return "", err
	}
	pkgBin, err := writeRemoteBinTmp(pkg)
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(pkgBin)
	moduleFile := path.Join("/tmp", fmt.Sprintf("ansible-module-%s.py", module.GetType()))
	err = r.Conn.Upload(pkgBin, moduleFile)
	if err != nil {
		return "", err
	}
	return r.Conn.ExecOutput(fmt.Sprintf("chmod +x %s && %s; res=$?; rm -f %s; exit $res", moduleFile, moduleFile, moduleFile))
}
