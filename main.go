package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func parseFunctions(filepath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, parser.ParseComments)

	if err != nil {
		log.Fatal(err)
	}

	var handlerFuncs []string

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
	fmt.Println(handlerFuncs)
}

func main() {
	app := kingpin.New("webapptester", "Generate boilerplate code to test your HTTP handlers")
	file := app.Arg("file", "Go file you would like to create tests for.").Required().String()
	kingpin.MustParse(app.Parse(os.Args[1:]))
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filepath := dir + "\\" + *file

	if _, err := os.Stat(filepath); err == nil {
		fmt.Printf("Creating tests for the file at: %s\n", filepath)
		parseFunctions(filepath)

	} else if os.IsNotExist(err) {
		fmt.Printf("Could not find the file at: %s\n", filepath)

	} else {
		fmt.Printf("Error looking for the file at: %s\n", filepath)
	}

}
