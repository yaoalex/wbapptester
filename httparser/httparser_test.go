package httparser_test

import (
	"fmt"
	"testing"

	"github.com/yaoalex/wbapptester/httparser"
)

func equal(a, b []httparser.FunctionInfo) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.Name != b[i].Name {
			return false
		}
		for i2, v2 := range v.MuxVars {
			if v2 != b[i].MuxVars[i2] {
				return false
			}
		}
	}
	return true
}

func TestParseFunctions(t *testing.T) {
	testCases := []struct {
		Name                string
		FilePath            string
		ExpectedPackageName string
		ExpectedMux         bool
		ExpectedFuncInfo    []httparser.FunctionInfo
	}{
		{
			Name:                "ParseFunctions: Get with mux router vars",
			FilePath:            "_testfiles/testapi1.go",
			ExpectedPackageName: "testapi",
			ExpectedMux:         true,
			ExpectedFuncInfo: []httparser.FunctionInfo{
				httparser.FunctionInfo{
					Name: "Get",
					MuxVars: []string{
						"return_status",
					},
				},
			},
		},
		{
			Name:                "ParseFunctions: Post with no mux router vars",
			FilePath:            "_testfiles/testapi2.go",
			ExpectedPackageName: "testapi",
			ExpectedMux:         false,
			ExpectedFuncInfo: []httparser.FunctionInfo{
				httparser.FunctionInfo{
					Name: "Post",
				},
			},
		},
		{
			Name:                "ParseFunctions: No http handlers in the file",
			FilePath:            "_testfiles/testapi3.go",
			ExpectedPackageName: "testapi",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			templateValues := httparser.ParseFunctions(tc.FilePath)
			if tc.ExpectedPackageName != templateValues.PackageName {
				t.Errorf("Function returned wrong package name: got %v want %v",
					templateValues.PackageName, tc.ExpectedPackageName)
			}
			if tc.ExpectedMux != templateValues.ContainsMux {
				t.Errorf("Function returned incorrect ContainsMux: got %v want %v",
					templateValues.ContainsMux, tc.ExpectedMux)
			}
			if !equal(tc.ExpectedFuncInfo, templateValues.FuncInfo) {
				t.Errorf("Function returned incorrect FunctionInfo: got %v want %v",
					templateValues.FuncInfo, tc.ExpectedFuncInfo)
			}
			fmt.Println(templateValues.FuncInfo)
		})
	}
}
