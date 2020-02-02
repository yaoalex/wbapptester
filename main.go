package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/alecthomas/kingpin.v2"
)

const outputTemplate = `package {{.PackageName}}_test

import "testing"

{{range .FuncNames -}}
func Test{{.}} (t *testing.T) {
	testCases := []struct {
		name string
	}{
		{name: "valid test case"},
		{name: "invalid test case"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){

		})
	}
}

{{end -}}
`

func parseFunctions(filePath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)

	if err != nil {
		log.Fatal(err)
	}

	var handlerFuncs []string
	packageName := fmt.Sprint(f.Name)

	for _, decl := range f.Decls {
		switch t := decl.(type) {
		case *ast.FuncDecl:
			responseWriterParamExists := false
			requestParamExists := false
			for _, param := range t.Type.Params.List {
				switch t2 := param.Type.(type) {
				case *ast.SelectorExpr:
					paramName := fmt.Sprint(t2.Sel.Name)
					if paramName == "ResponseWriter" {
						responseWriterParamExists = true
					}
				case *ast.StarExpr:
					paramName := fmt.Sprint(t2.X)
					if paramName == "&{http Request}" {
						requestParamExists = true
					}
				}
			}
			if responseWriterParamExists && requestParamExists {
				handlerFuncs = append(handlerFuncs, fmt.Sprint(t.Name))
			}
		}
	}
	if len(handlerFuncs) > 0 {
		generateTestFile(packageName, filePath, handlerFuncs)
	}
}

func generateTestFile(packageName, filePath string, handlerFuncs []string) {
	extension := filepath.Ext(filePath)
	basePath := filepath.Base(filePath)
	testFileName := filepath.Base(filePath)[0:len(basePath)-len(extension)] + "_test.go"
	outFile, err := os.Create(testFileName)
	if err != nil {
		fmt.Printf("Error creating test file named: %s\n", testFileName)
	}
	var templateValues = struct {
		FuncNames   []string
		PackageName string
	}{
		FuncNames:   handlerFuncs,
		PackageName: packageName,
	}
	tmpl := template.Must(template.New("out").Parse(outputTemplate))
	if err := tmpl.Execute(outFile, templateValues); err != nil {
		panic(err)
	}
	if err := outFile.Close(); err != nil {
		panic(err)
	}
}

func main() {
	app := kingpin.New("webapptester", "Generate boilerplate code to test your HTTP handlers")
	file := app.Arg("file", "Go file you would like to create tests for.").Required().String()
	kingpin.MustParse(app.Parse(os.Args[1:]))
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath := dir + "\\" + *file

	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("Creating tests for the file at: %s\n", filePath)
		parseFunctions(filePath)

	} else if os.IsNotExist(err) {
		fmt.Printf("Could not find the file at: %s\n", filePath)

	} else {
		fmt.Printf("Error looking for the file at: %s\n", filePath)
	}

}
