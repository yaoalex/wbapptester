package main

// TemplateValues contains information to be passed into the test template
type TemplateValues = struct {
	FuncInfo    []FunctionInfo
	PackageName string
}

const outputTemplate = `package {{.PackageName}}

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// THIS IS GENERATED CODE BY WEBAPPTESTER
// you will need to edit this code to suit your API's needs

{{range $funcinfo := .FuncInfo}} func Test{{$funcinfo.Name}} (t *testing.T) {
	testCases := []struct {
		Name string
		ExpectedStatus int
		MuxVars map[string]string
	}{
		{
			Name: "{{$funcinfo.Name}}: valid test case",
			ExpectedStatus: http.StatusOK,
			MuxVars:        map[string]string  {
				{{range $muxvar := $funcinfo.MuxVars}} "{{$muxvar}}" : "valid_value", 
				{{end}}
			},
		},
		{
			Name: "{{$funcinfo.Name}}: invalid test case",
			ExpectedStatus: http.StatusBadRequest,
			MuxVars:        map[string]string  {
				{{range $muxvar := $funcinfo.MuxVars}} "{{$muxvar}}" : "invalid_value", 
				{{end}}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T){
			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc({{$funcinfo.Name}})

            req = mux.SetURLVars(req, tc.MuxVars)

			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tc.ExpectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.ExpectedStatus)
			}
		})
	}
}
{{end}}
`
