package ansible

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
	"time"
)

type Data struct {
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

type Args struct {
	ModuleArgs interface{} `json:"ANSIBLE_MODULE_ARGS"`
}

func ArgstoString(a Args) string {
	val, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	res := string(val)
	// res = strings.ReplaceAll(res, "false", "False")
	// res = strings.ReplaceAll(res, "true", "True")
	return res
}

func loadModulesChunk() (string, error) {
	raw, err := ioutil.ReadFile("build/ansible_modules.txt")
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

func Execute(module string, params interface{}, result interface{}) (string, error) {
	raw, err := loadModulesChunk()
	if err != nil {
		return "", err
	}
	now := time.Now()
	bin, err := exec.LookPath("python")
	if err != nil {
		return "", err
	}
	d := Data{
		// Shebang:   "#!/usr/local/Cellar/ansible/6.0.0/libexec/bin/python3.10",
		Shebang:   "#!" + bin,
		ZipData:   string(raw),
		Year:      now.Year(),
		Month:     int(now.Month()),
		Day:       now.Day(),
		Hour:      now.Hour(),
		Minute:    now.Minute(),
		Second:    now.Second(),
		ModuleFqn: "ansible.modules." + module,
		Params: ArgstoString(Args{
			ModuleArgs: params,
		}),
		AnsibleModule: "ansible" + module,
	}

	t, err := template.New("ansible_exec").Parse(ansibleTemplate)
	if err != nil {
		return "", err
	}

	f, err := os.CreateTemp("", d.AnsibleModule)
	if err != nil {
		return "", err
	}
	defer func() {
		f.Close()
		os.RemoveAll(f.Name())
	}()
	err = t.Execute(f, d)
	if err != nil {
		return "", err
	}
	err = os.Chmod(f.Name(), 0700)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer

	cmd := exec.Command(f.Name())
	cmd.Stdin = os.Stdout
	cmd.Stdout = &buf
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	de := json.NewDecoder(&buf)
	return buf.String(), de.Decode(result)

}
