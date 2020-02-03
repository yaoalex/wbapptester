package main

const outputTemplate = `package {{.PackageName}}

import (
	"net/http"
	"net/http/httptest"
	"testing"
)
{{ $muxvars := .MuxVars}}
// THIS IS GENERATED CODE BY WEBAPPTESTER
// you will need to edit this code to suit your API's needs

{{range $funcname := .FuncNames}} func Test{{$funcname}} (t *testing.T) {
	testCases := []struct {
		Name string
		ExpectedStatus int
		MuxVars map[string]string
	}{
		{
			Name: "{{$funcname}}: valid test case",
			ExpectedStatus: http.StatusOK,
			MuxVars:        map[string]string  {
				{{range $muxvar := $muxvars}} "{{$muxvar}}" : "valid_value", {{end}}
			},
		},
		{
			Name: "{{$funcname}}: invalid test case",
			ExpectedStatus: http.StatusBadRequest,
			MuxVars:        map[string]string  {
				{{range $muxvar := $muxvars}} "{{$muxvar}}" : "invalid_value", 	{{end}}
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
			handler := http.HandlerFunc({{$funcname}})

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
