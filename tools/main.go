package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"flag"

	"github.com/go-python/gpython/ast"
	"github.com/go-python/gpython/parser"
	"github.com/go-python/gpython/py"
	yaml "gopkg.in/yaml.v3"
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
	modules, err := parseModules(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, m := range modules {
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
}

func parseModules(dir string) ([]Module, error) {
	modules := []Module{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return modules, err
	}
	for _, file := range files {
		p := filepath.Join(dir, file.Name())
		if !strings.HasPrefix(file.Name(), "_") {
			m := Module{Path: p}
			err := m.parse()
			if err != nil {
				// skip this invalid module
				if err.Error() == "skip" {
					fmt.Println("skipped: " + p)
					continue
				}
				return modules, err
			}
			modules = append(modules, m)
		}
	}
	return modules, nil
}

type Module struct {
	NormalizedName   string
	ModuleName       string             `yaml:"module"`
	ShortDescription string             `yaml:"short_description"`
	Description      []string           `yaml:"description"`
	Options          map[string]*Option `yaml:"options"`
	Returns          map[string]*Return
	Path             string
	Documentation    string
	Return           string
}

type Option struct {
	NormalizedName string
	StructTag      string
	GoType         string
	GoElements     string

	Description []string    `yaml:"description"`
	Type        string      `default:"str" yaml:"type"`
	Required    bool        `default:"no" yaml:"required"`
	Default     interface{} `yaml:"default"`
	Elements    string      `yaml:"elements"`
}

type Return struct {
	NormalizedName string
	Description    interface{} `yaml:"description"`
	Returned       string      `yaml:"returned"`
	// complex type is not supported...
	Type   string      `default:"str" yaml:"type"`
	Sample interface{} `yaml:"sample"`

	GoType    string
	StructTag string
}

func (m *Module) parse() error {
	f, _ := os.Open(m.Path)
	Ast, err := parser.Parse(f, m.Path, py.ExecMode)
	if err != nil {
		return err
	}
	defer f.Close()
	switch node := Ast.(type) {
	case *ast.Module:
		for _, stmt := range node.Body {
			err = m.parseStmt(stmt)
			if err != nil {
				return err
			}
		}
	}
	err = yaml.Unmarshal([]byte(m.Documentation), m)
	if err != nil {
		return fmt.Errorf("%s: %s", m.Path, err)
	}

	if m.ModuleName == "" {
		return fmt.Errorf("skip")
	}
	m.Returns = map[string]*Return{}
	err = yaml.Unmarshal([]byte(m.Return), m.Returns)
	if err != nil {
		return fmt.Errorf("Return: %s: %s", m.Path, err)
	}
	return m.normalize()
}

func (m *Module) normalizeName(val string) string {
	val = strings.ReplaceAll(val, "-", "_")
	vals := strings.Split(val, "_")
	for i, v := range vals {
		vals[i] = strings.Title(v)
	}
	return strings.Join(vals, "")
}

func toGoType(ty string, elementType string) string {
	switch ty {
	case "path":
		return "string"
	case "float":
		return "float64"
	case "int":
		return "int"
	case "str":
		return "string"
	case "list":
		elType := "map[string]interface{}"
		if elementType != "" {
			elType = toGoType(elementType, "")
		}
		return "[]" + elType
	case "complex":
		return "interface{}"
	case "dict":
		return "map[string]interface{}"
	case "raw":
		return "string"
	case "bool":
		return "bool"
	case "":
		return "string"
	default:
		panic("Not supported type: " + ty)
	}
}
func (m *Module) normalize() error {
	m.NormalizedName = m.normalizeName(m.ModuleName)
	for name, o := range m.Options {
		o.NormalizedName = m.normalizeName(name)
		o.StructTag = "`yaml:\"" + name + ",omitempty\" json:\"" + name + ",omitempty\"`"
		o.GoType = toGoType(o.Type, o.Elements)
	}
	for name, r := range m.Returns {
		r.GoType = toGoType(r.Type, "")
		r.NormalizedName = m.normalizeName(name)
		r.StructTag = "`yaml:\"" + name + ",omitempty\" json:\"" + name + ",omitempty\"`"
	}
	return nil
}
func (m *Module) parseStmt(stmt ast.Stmt) error {
	switch node := stmt.(type) {
	case *ast.Assign:
		id, val, err := m.parseAssign(node)
		if err != nil {
			return err
		}
		switch id {
		case "DOCUMENTATION":
			m.Documentation = val
		case "RETURN":
			m.Return = val
		}
	}
	return nil
}

func (m *Module) parseAssign(node *ast.Assign) (id string, value string, err error) {
	if len(node.Targets) != 1 {
		return
	}
	switch tn := node.Targets[0].(type) {
	case *ast.Name:
		id = string(tn.Id)
	}
	switch valN := node.Value.(type) {
	case *ast.Str:
		value = string(valN.S)
	}
	return
}

var (
	moduleTemplate = `
// Autogenerated
package module

import (
	"aig/pkg/ansible"
	
)

type {{ .NormalizedName }} struct {
	ModuleName string
	Options {{ .NormalizedName }}Options
	Result {{ .NormalizedName }}Result
}

type {{ .NormalizedName }}Options struct {
	{{range $name, $opt := .Options }}
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
	raw, err := ansible.Execute(m.ModuleName, m.Options, &m.Result)
	m.Result.Raw = raw
	return err
}

func New{{.NormalizedName}}() *{{.NormalizedName}} {
	return &{{.NormalizedName}} {
		ModuleName: "{{.ModuleName}}",
	}
}
	`
)
