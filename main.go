package main

import (
	"encoding/json"

	"github.com/gopherjs/gopherjs/js"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func main() {
	js.Module.Get("exports").Set("parse", parse)
}

type Variable struct {
	Name string `hcl:"name,label"`
}

type Output struct {
	Name string `hcl:"name,label"`
}

type Resource struct {
	Type      string         `hcl:"type,label"`
	Name      string         `hcl:"name,label"`
	Config    hcl.Body       `hcl:",remain"`
	DependsOn hcl.Expression `hcl:"depends_on,attr"`
}

type Module struct {
	Name      string         `hcl:"name,label"`
	Providers hcl.Expression `hcl:"providers"`
}

type Root struct {
	Outputs   []*Output   `hcl:"output,block"`
	Variables []*Variable `hcl:"variable,block"`
	Resources []*Resource `hcl:"resource,block"`
	Modules   []*Module   `hcl:"module,block"`
}

// Parse a HCL string into a JSON object
func parse(input string) (output *js.Object) {
	parsed, err := parseHcl(input)
	if err != nil {
		return
	}

	output = js.Global.Get("JSON").Call("parse", string(parsed))

	return
}

func parseHcl(input string) (string, error) {
	parser := hclparse.NewParser()
	hclFile, _ := parser.ParseHCL([]byte(input), "<input>")

	module := tfconfig.NewModule("<virtual>")
	tfconfig.LoadModuleFromFile(hclFile, module)

	moduleJSON, err := json.Marshal(module)

	if err != nil {
		return "", err
	}

	return string(moduleJSON), nil
}
