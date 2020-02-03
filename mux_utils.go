package main

import (
	"fmt"
	"go/ast"
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
