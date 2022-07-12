package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-python/gpython/ast"
	"github.com/go-python/gpython/parser"
	"github.com/go-python/gpython/py"
	yaml "gopkg.in/yaml.v3"

	b64 "encoding/base64"
	"encoding/json"
)

func ParseModulesFromSpec(raw string) (*Spec, error) {
	spec := &Spec{
		Modules: []Module{},
	}
	jsonRaw, err := b64.StdEncoding.DecodeString(raw)
	if err != nil {
		return spec, err
	}
	err = json.Unmarshal(jsonRaw, spec)
	if err != nil {
		return spec, err
	}

	return spec, err
}

// ParseModules parses modules from Ansible source
func ParseModules(dir string) (*Spec, error) {
	spec := &Spec{
		Modules: []Module{},
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return spec, err
	}
	for _, file := range files {
		p := filepath.Join(dir, file.Name())
		if includeModule(file.Name()) {
			m := Module{Path: p}
			err := m.parse()
			if err != nil {
				// skip this invalid module
				if err.Error() == "skip" {
					fmt.Println("skipped: " + p)
					continue
				}
				return spec, err
			}
			spec.Modules = append(spec.Modules, m)
		}
	}
	return spec, nil
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

func includeModule(name string) bool {
	// exclude modueles which are known to not work
	return !strings.HasPrefix(name, "_") &&
		!strings.HasPrefix(name, "include_") &&
		!strings.HasPrefix(name, "import_") &&
		!strings.HasPrefix(name, "set_") &&
		!strings.HasPrefix(name, "fail") &&
		!strings.HasPrefix(name, "wait_for_connection.py") &&
		!strings.HasPrefix(name, "gather_facts.py") &&
		!strings.HasPrefix(name, "debug.py") &&
		!strings.HasPrefix(name, "assert.py") &&
		!strings.HasPrefix(name, "raw.py") &&
		!strings.HasPrefix(name, "fetch.py") &&
		!strings.HasPrefix(name, "add_host.py") &&
		!strings.HasPrefix(name, "script.py") &&
		!strings.HasPrefix(name, "shell.py") &&
		!strings.HasPrefix(name, "validate_argument_spec.py")
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
	for name, o := range m.Params {
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