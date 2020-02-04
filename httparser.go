package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"text/template"
)

func generateTestFile(testFileName string, templateValues *TemplateValues) error {
	outFile, err := os.Create(testFileName)
	if err != nil {
		fmt.Printf("Error creating test file named: %s\n", testFileName)
	}
	tmpl := template.Must(template.New("out").Parse(outputTemplate))
	if err := tmpl.Execute(outFile, templateValues); err != nil {
		return err
	}
	if err := outFile.Close(); err != nil {
		return err
	}
	return nil
}

func parseFunctions(filePath string) *TemplateValues {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)

	if err != nil {
		log.Fatal(err)
	}

	var funcInfos []FunctionInfo
	packageName := fmt.Sprint(f.Name)
	containsMux := false

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
				muxVars := getMuxVars(t)
				if len(muxVars) > 0 {
					containsMux = true
				}
				funcInfo := FunctionInfo{
					Name:    fmt.Sprint(t.Name),
					MuxVars: muxVars,
				}
				funcInfos = append(funcInfos, funcInfo)
			}
		}
	}
	templateValues := TemplateValues{
		FuncInfo:    funcInfos,
		PackageName: packageName,
		ContainsMux: containsMux,
	}
	return &templateValues
}
