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
	ansibleTag := flag.String("V", "devel", "Ansible Version Tag")
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
	p := parser.Parser{
		Dir:        *modulePath,
		AnsibleTag: *ansibleTag,
	}
	spec, err := p.Parse()
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
		t, err := template.New("moduleTemplate:" + m.ModuleName).Parse(moduleTemplate)
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

		
		type factory func() types.Module

		func addModuleFactory(ty string, f factory) {
			factories[ty] = f
		}
		
		func ModuleByName(ty string) types.Module {
			return factories[ty]()
		}
	`

	moduleTemplate = `
		
		package module

		// Autogenerated file

		import (
			"github.com/weakpixel/aig/pkg/types"
			"fmt"
		)

		func init() {
			addModuleFactory("{{.ModuleName}}", func() types.Module {
				return New{{.NormalizedName}}()
			})
		}

		//
		// {{.NormalizedName}} ({{ .ModuleName }}) - {{ .ShortDescription }}
		// 
		func New{{.NormalizedName}}() *{{.NormalizedName}} {
			module := {{.NormalizedName}} {}
			// Create dynamic param values
			paramValues := map[string]types.Value{}
			{{- range $name, $opt := .Params }}
			    {{- if eq $opt.GoType "string" }}
				paramValues["{{ $name }}"] = types.NewStringValue(&module.Params.{{ $opt.NormalizedName }})
				{{- else if eq $opt.GoType "bool" }}
				paramValues["{{ $name }}"] = types.NewBoolValue(&module.Params.{{ $opt.NormalizedName }})
				{{- else if eq $opt.GoType "int" }}
				paramValues["{{ $name }}"] = types.NewIntValue(&module.Params.{{ $opt.NormalizedName }})
				{{- else if eq $opt.GoType "[]string" }}
				paramValues["{{ $name }}"] = types.NewStringArrayValue(&module.Params.{{ $opt.NormalizedName }})
				{{- else if eq $opt.GoType "float64" }}
				paramValues["{{ $name }}"] = types.NewFloat64Value(&module.Params.{{ $opt.NormalizedName }})
				{{- else if eq $opt.GoType "map[string]string" }}
				paramValues["{{ $name }}"] = types.NewStringMapValue(&module.Params.{{ $opt.NormalizedName }})
				{{- else }}
				// NOT SUPPORTED: {{ $name }} {{ $opt.NormalizedName }} {{$opt.GoType}}
				{{- end -}}
			{{ end }}
			module.Params.values = paramValues

			// Create dynamic result values
			resultValues := map[string]types.Value{}
			{{range $name, $opt := .Returns }}
				{{- if eq $opt.GoType "string" }}
				resultValues["{{ $name }}"] = types.NewStringValue(&module.Result.{{ $opt.NormalizedName }})
				{{- else if eq $opt.GoType "bool" }}
				resultValues["{{ $name }}"] = types.NewBoolValue(&module.Result.{{ $opt.NormalizedName }})
				{{- else if eq $opt.GoType "int" }}
				resultValues["{{ $name }}"] = types.NewIntValue(&module.Result.{{ $opt.NormalizedName }})
				{{- else if eq $opt.GoType "[]string" }}
				resultValues["{{ $name }}"] = types.NewStringArrayValue(&module.Result.{{ $opt.NormalizedName }})
				{{- else if eq $opt.GoType "float64" }}
				resultValues["{{ $name }}"] = types.NewFloat64Value(&module.Result.{{ $opt.NormalizedName }})
				{{- else if eq $opt.GoType "map[string]string" }}
				resultValues["{{ $name }}"] = types.NewStringMapValue(&module.Result.{{ $opt.NormalizedName }})
				{{- else }}
				// NOT SUPPORTED: {{ $name }} {{ $opt.NormalizedName }} {{$opt.GoType}}
				{{- end -}}
			{{ end }}
			module.Result.values = resultValues

			return &module
		}

		// {{.NormalizedName}} ({{ .ModuleName }}) - {{ .ShortDescription }}
		//
		{{- range $d := .Description }}
		// {{ $d }}
		//
		{{- end }}
		//
		// Source: {{ .SourceLink }}
		type {{ .NormalizedName }} struct {
			Params {{ .NormalizedName }}Params
			Result {{ .NormalizedName }}Result
		}

		type {{ .NormalizedName }}Params struct {
			{{range $name, $opt := .Params }}
				// {{ $opt.NormalizedName }} 
				{{- range $d := $opt.Description }}
				// {{ $d }}
				{{- end }}
				//
				// Default: {{ $opt.Default }}
				// Required: {{ $opt.Required }}
				{{ $opt.NormalizedName }} {{ $opt.GoType }} {{ $opt.StructTag }}
				
			{{ end }}
			values map[string]types.Value
		}

		func (p *{{ .NormalizedName }}Params) Names() []string {
			names := []string{}
			for name := range p.values {
				names = append(names, name)
			}
			return names
		}

		func (p *{{ .NormalizedName }}Params) Set(name string, value interface{}) error {
			v, ok := p.values[name]
			if !ok {
				return fmt.Errorf("no param with name %q", name)
			}
			return v.Set(value)
		}

		func (p *{{ .NormalizedName }}Params) Get(name string) (interface{}, error) {
			v, ok := p.values[name]
			if !ok {
				return nil, fmt.Errorf("no param with name %q", name)
			}
			return v.Get(), nil
		}

		type {{ .NormalizedName }}Result struct {
			types.CommonReturn
			Raw string 
			{{range $name, $opt := .Returns }}
				// {{ $opt.NormalizedName }} 
				{{- range $d := $opt.Desc }}
				// {{ $d }}
				{{- end }}
				{{ $opt.NormalizedName }} {{ $opt.GoType }} {{ $opt.StructTag }}
			{{ end }}
			values map[string]types.Value
		}

		func (r *{{ .NormalizedName }}Result) Names() []string {
			names := []string{}
			for name := range r.values {
				names = append(names, name)
			}
			return names
		}

		func (r *{{ .NormalizedName }}Result) Set(name string, value interface{}) error {
			v, ok := r.values[name]
			if !ok {
				return fmt.Errorf("no param with name %q", name)
			}
			return v.Set(value)
		}
		
		func (r *{{ .NormalizedName }}Result) Get(name string) (interface{}, error) {
			v, ok := r.values[name]
			if !ok {
				return nil, fmt.Errorf("no param with name %q", name)
			}
			return v.Get(), nil
		}

		func (m *{{ .NormalizedName }}) GetResult() types.Result {
			return &m.Result
		}
		
		func (m *{{ .NormalizedName }}) GetResultRaw() string {
			return m.Result.Raw
		}

		func (m *{{ .NormalizedName }}) GetCommonResult() types.CommonReturn {
			return m.Result.CommonReturn
		}

		func (m *{{ .NormalizedName }}) GetParams() types.Params {
			return &m.Params
		}

		func (m *{{ .NormalizedName }}) GetType() string {
			return "{{ .ModuleName }}"
		}

	`
)
