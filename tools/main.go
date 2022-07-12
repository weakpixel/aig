package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/weakpixel/aig/pkg/parser"
	"github.com/weakpixel/aig/pkg/types"

	b64 "encoding/base64"

	"flag"
)

func main() {
	modulePath := flag.String("m", "", "Path to ansible modules")
	outPath := flag.String("o", "", "Path to go generated taks")
	flag.Parse()
	dir := *modulePath
	if dir == "" {
		fmt.Println("[ERROR] Please define ansible module path with -m <path>")
		os.Exit(1)
	}
	out := *outPath
	if out == "" {
		fmt.Println("[ERROR] Please define go output path with -o <path>")
		os.Exit(1)
	}
	spec, err := parser.ParseModules(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, m := range spec.Modules {
		outFile := filepath.Join(out, m.ModuleName+".go")
		f, err := os.Create(outFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		t, err := template.New(m.ModuleName).Parse(moduleTemplate)
		if err != nil {
			panic(err)
		}
		err = t.Execute(f, &m)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
	writeModuleSpec(spec, out)
	writeModuleZip(out)
}

func writeModuleZip(out string) {
	raw, err := ioutil.ReadFile("build/ansible_modules.zip.base64")
	if err != nil {
		panic(err)
	}
	outDir := filepath.Join(out, "src")
	err = os.MkdirAll(outDir, 0777)
	if err != nil {
		panic(err)
	}
	outFile := filepath.Join(outDir, "module-source.go")
	fmt.Println(outFile)
	constantFile, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer constantFile.Close()

	t, err := template.New("module-source.go").Parse(moduleSourceTemplate)
	if err != nil {
		panic(err)
	}
	err = t.Execute(constantFile, map[string]interface{}{
		"Source": strings.TrimSpace(string(raw)),
	})
	if err != nil {
		panic(err)
	}
}

func writeModuleSpec(spec *types.Spec, out string) {
	raw, err := json.Marshal(spec)
	if err != nil {
		panic(err)
	}
	outFile := filepath.Join(out, "spec.go")
	constantFile, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer constantFile.Close()

	t, err := template.New("spec.go").Parse(specTemplate)
	if err != nil {
		panic(err)
	}
	encodedRaw := b64.StdEncoding.EncodeToString(raw)
	err = t.Execute(constantFile, map[string]interface{}{
		"Source": encodedRaw,
	})
	if err != nil {
		panic(err)
	}
}

var (
	moduleSourceTemplate = `
		package src
		const (
			// Ansible Modules source (zippped and base64 encoded)
			ModuleSources = {{ .Source | printf "%q" }}
		)
	`

	specTemplate = `
		package module
		import (
			"github.com/weakpixel/aig/pkg/types"
			b64 "encoding/base64"
		)
		func GetSpec() (*types.Spec, error) {
			jsonRaw, err := b64.StdEncoding.DecodeString(moduleSpecJSON)
			if err != nil {
				return nil, err
			}
			return types.Parse(jsonRaw)
		}
		const (
			// ModuleSpecJson contains source model spec base64 encoded
			moduleSpecJSON = {{ .Source | printf "%q" }}
			
		)
		var (
			factories = map[string]factory{}
		)

		type Module interface {
			GetResult() interface{}
			GetResultRaw() string
			GetParams() interface{}
			GetType() string
			Run() error
		}

		type factory func() Module

		func addModuleFactory(ty string, f factory) {
			factories[ty] = f
		}
		
		func ModuleByName(ty string) Module {
			return factories[ty]()
		}
	`

	moduleTemplate = `
		// Autogenerated
		package module

		import (
			"github.com/weakpixel/aig/pkg/ansible"
			
		)

		func init() {
			addModuleFactory("{{.ModuleName}}", func() Module {
				return New{{.NormalizedName}}()
			})
		}

		type {{ .NormalizedName }} struct {
			ModuleName string
			Params {{ .NormalizedName }}Params
			Result {{ .NormalizedName }}Result
		}

		type {{ .NormalizedName }}Params struct {
			{{range $name, $opt := .Params }}
				// {{ $opt.NormalizedName }} 
				{{ $opt.NormalizedName }} {{ $opt.GoType }} {{ $opt.StructTag }}
			{{ end }}
		}

		type {{ .NormalizedName }}Result struct {
			Raw string 
			{{range $name, $opt := .Returns }}
				// {{ $opt.NormalizedName }} 
				{{ $opt.NormalizedName }} {{ $opt.GoType }} {{ $opt.StructTag }}
			{{ end }}
		}

		func (m *{{ .NormalizedName }}) Run() error {
			raw, err := ansible.Execute(m.ModuleName, m.Params, &m.Result)
			m.Result.Raw = raw
			return err
		}

		func (m *{{ .NormalizedName }}) GetResult() interface{} {
			return &m.Result
		}

		func (m *{{ .NormalizedName }}) GetResultRaw() string {
			return m.Result.Raw
		}

		func (m *{{ .NormalizedName }}) GetParams() interface{} {
			return &m.Params
		}

		func (m *{{ .NormalizedName }}) GetType() string {
			return m.ModuleName
		}

		func New{{.NormalizedName}}() *{{.NormalizedName}} {
			return &{{.NormalizedName}} {
				ModuleName: "{{.ModuleName}}",
			}
		}
	`
)
