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

func getMuxVars(t *ast.FuncDecl) []string {
	paramName := ""
	var muxVars []string
	for _, statement := range t.Body.List {
		switch s := statement.(type) {
		case *ast.AssignStmt:
			if paramName == "" {
				switch s2 := s.Rhs[0].(type) {
				case *ast.CallExpr:
					switch s3 := s2.Fun.(type) {
					case *ast.SelectorExpr:
						expr := fmt.Sprint(s3)
						if expr == "&{mux Vars}" {
							switch s4 := s.Lhs[0].(type) {
							case *ast.Ident:
								paramName = s4.Name
							}
						}
					}
				}
			} else {
				switch s5 := s.Rhs[0].(type) {
				case *ast.IndexExpr:
					switch s6 := s5.X.(type) {
					case *ast.Ident:
						if s6.Name == paramName {
							switch s7 := s5.Index.(type) {
							case *ast.BasicLit:
								cleanedValue := s7.Value[1 : len(s7.Value)-1] //remove quotes around string
								muxVars = append(muxVars, cleanedValue)
							}

						}
					}
				}
			}
		}
	}
	return muxVars
}

func parseFunctions(filePath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)

	if err != nil {
		log.Fatal(err)
	}

	var handlerFuncs []string
	var muxVars []string
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
				muxVars = getMuxVars(t)
			}
		}
	}
	if len(handlerFuncs) > 0 {
		generateTestFile(packageName, filePath, handlerFuncs, muxVars)
	}
}

func generateTestFile(packageName, filePath string, handlerFuncs, muxVars []string) {
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
		MuxVars     []string
	}{
		FuncNames:   handlerFuncs,
		PackageName: packageName,
		MuxVars:     muxVars,
	}
	fmt.Println(templateValues)
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
