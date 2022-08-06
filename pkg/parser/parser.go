package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/weakpixel/aig/pkg/types"

	"github.com/go-python/gpython/ast"
	"github.com/go-python/gpython/parser"
	"github.com/go-python/gpython/py"
	yaml "gopkg.in/yaml.v3"
)

type Parser struct {
	Dir        string
	AnsibleTag string
}

// ParseModules parses modules from Ansible source
func (p *Parser) Parse() (*types.Spec, error) {
	spec := &types.Spec{
		Modules: []*types.Module{},
	}
	files, err := ioutil.ReadDir(p.Dir)
	if err != nil {
		return spec, err
	}
	for _, file := range files {
		pth := filepath.Join(p.Dir, file.Name())
		if p.includeModule(file.Name()) {
			m := &types.Module{Path: pth}
			err := p.parse(m)
			if err != nil {
				// skip this invalid module
				if err.Error() == "skip" {
					fmt.Println("skipped: " + pth)
					continue
				}
				return spec, err
			}
			spec.Modules = append(spec.Modules, m)
		}
	}
	return spec, nil
}

func (p *Parser) parse(m *types.Module) error {
	f, _ := os.Open(m.Path)
	Ast, err := parser.Parse(f, m.Path, py.ExecMode)
	if err != nil {
		return err
	}
	defer f.Close()
	switch node := Ast.(type) {
	case *ast.Module:
		for _, stmt := range node.Body {
			err = p.parseStmt(m, stmt)
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
	m.Returns = map[string]*types.Return{}
	err = yaml.Unmarshal([]byte(m.Return), m.Returns)
	if err != nil {
		return fmt.Errorf("Return: %s: %s", m.Path, err)
	}
	return p.normalize(m)
}

func (p *Parser) includeModule(name string) bool {
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
		!strings.HasPrefix(name, "group_by.py") &&
		!strings.HasPrefix(name, "copy.py") &&
		!strings.HasPrefix(name, "pause.py") &&
		!strings.HasPrefix(name, "ping.py") &&
		!strings.HasPrefix(name, "service.py") &&
		!strings.HasPrefix(name, "service_facts.py") &&
		!strings.HasPrefix(name, "validate_argument_spec.py")

}

func (p *Parser) normalizeName(m *types.Module, val string) string {
	val = strings.ReplaceAll(val, "-", "_")
	vals := strings.Split(val, "_")
	for i, v := range vals {
		vals[i] = strings.Title(v)
	}
	return strings.Join(vals, "")
}

func (p *Parser) toGoType(ty string, elementType string) string {
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
			elType = p.toGoType(elementType, "")
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

func (p *Parser) normalize(m *types.Module) error {
	m.NormalizedName = p.normalizeName(m, m.ModuleName)
	for name, o := range m.Params {
		o.NormalizedName = p.normalizeName(m, name)
		o.StructTag = "`yaml:\"" + name + ",omitempty\" json:\"" + name + ",omitempty\"`"
		o.GoType = p.toGoType(o.Type, o.Elements)
	}

	for name, r := range m.Returns {
		r.GoType = p.toGoType(r.Type, "")
		r.NormalizedName = p.normalizeName(m, name)
		r.StructTag = "`yaml:\"" + name + ",omitempty\" json:\"" + name + ",omitempty\"`"
	}

	m.SourceLink = fmt.Sprintf("https://github.com/ansible/ansible/blob/%s/lib/ansible/modules/%s.py", p.AnsibleTag, m.ModuleName)
	return nil
}
func (p *Parser) parseStmt(m *types.Module, stmt ast.Stmt) error {
	switch node := stmt.(type) {
	case *ast.Assign:
		id, val, err := p.parseAssign(m, node)
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

func (p *Parser) parseAssign(m *types.Module, node *ast.Assign) (id string, value string, err error) {
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
