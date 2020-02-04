package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// FunctionInfo contains information about the http handler functions
type FunctionInfo struct {
	Name    string
	MuxVars []string
}

func generateTestFile(packageName, filePath string, funcInfos []FunctionInfo) {
	extension := filepath.Ext(filePath)
	basePath := filepath.Base(filePath)
	testFileName := filepath.Base(filePath)[0:len(basePath)-len(extension)] + "_test.go"
	outFile, err := os.Create(testFileName)
	if err != nil {
		fmt.Printf("Error creating test file named: %s\n", testFileName)
	}
	templateValues := TemplateValues{
		FuncInfo:    funcInfos,
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

func parseFunctions(filePath string) ([]FunctionInfo, string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)

	if err != nil {
		log.Fatal(err)
	}

	var funcInfos []FunctionInfo
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
				funcInfo := FunctionInfo{
					Name:    fmt.Sprint(t.Name),
					MuxVars: getMuxVars(t),
				}
				funcInfos = append(funcInfos, funcInfo)
			}
		}
	}
	return funcInfos, packageName
}
