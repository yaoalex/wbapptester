package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"reflect"

	"gopkg.in/alecthomas/kingpin.v2"
)

func test(t string, a int) {

}

func main() {
	app := kingpin.New("webapptester", "Generate tests for your HTTP handler functions")
	file := app.Arg("file", "Go file you would like to create tests for.").Required().String()
	kingpin.MustParse(app.Parse(os.Args[1:]))
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fullpath := dir + "\\" + *file

	if _, err := os.Stat(fullpath); err == nil {
		fmt.Printf("Would create tests for the file: %s\n", fullpath)
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, fullpath, nil, parser.ParseComments)
		// var typeNameBuf bytes.Buffer
		if err != nil {
			log.Fatal(err)
		}

		for _, decl := range f.Decls {
			switch t := decl.(type) {
			case *ast.FuncDecl:
				fmt.Printf("Function name: %v\n", t.Name)
				for _, param := range t.Type.Params.List {
					switch t2 := param.Type.(type) {
					case *ast.SelectorExpr:
						fmt.Println(t2.Sel)
					case *ast.StarExpr:
						fmt.Println(t2.X)
					default:
						fmt.Println(reflect.TypeOf(t2))
					}
				}
			}
		}

	} else if os.IsNotExist(err) {
		fmt.Printf("Could not find the file; %s\n", fullpath)

	} else {
		fmt.Printf("Error looking for the file; %s\n", fullpath)
	}

}
