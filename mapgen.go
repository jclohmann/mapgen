package mapgen

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type MapGen struct {
	path        string
	packageName string
	typeName    string
}

func NewMapGen(path, packageName, typeName string) MapGen {
	return MapGen{
		path:        path,
		packageName: packageName,
		typeName:    typeName,
	}
}

func (mg MapGen) Generate() error {
	filename := filepath.Join(mg.path, strings.ToLower(mg.typeName)+"-mapping.go")
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}

	data := mapGenTplData{
		PackageName: mg.packageName,
		TypeName:    mg.typeName,
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	err = mapGenTpl.Execute(file, data)

	return err
}

type mapGenTplData struct {
	PackageName string
	TypeName    string
}

var mapGenTpl *template.Template = template.Must(template.New("mapGen").Parse(`package {{.PackageName}}

type {{.TypeName}}Slice []{{.TypeName}}
type {{.TypeName}}MapFunc func({{.TypeName}}) interface{}

func (slice {{.TypeName}}Slice) Map(fn {{.TypeName}}MapFunc) []interface{} {
	var output []interface{}
	for _, item := range slice {
		out := fn(item)
		output = append(output, out)
	}
	return output
}

type {{.TypeName}}FilterFunc func({{.TypeName}}) bool

func (slice {{.TypeName}}Slice) Filter(fn {{.TypeName}}FilterFunc) {{.TypeName}}Slice {
	var output {{.TypeName}}Slice
	for _, item := range slice {
		if fn(item) {
			output = append(output, item)
		}
	}
	return output
}

type {{.TypeName}}EachFunc func({{.TypeName}})

func (slice {{.TypeName}}Slice) Each(fn {{.TypeName}}EachFunc) {
	for _, item := range slice {
		fn(item)
	}
}
`))
