package httparser

// TemplateValues contains information to be passed into the test template
type TemplateValues = struct {
	FuncInfo    []FunctionInfo
	PackageName string
	ContainsMux bool
}

// FunctionInfo contains information about the http handler functions
type FunctionInfo struct {
	Name    string
	MuxVars []string
}

const outputTemplate = `package {{.PackageName}}

import (
	"net/http"
	"net/http/httptest"
	"testing"

	{{if .ContainsMux}}
	"github.com/gorilla/mux"
	{{end}}
)

{{ $containsmux := .ContainsMux }}

// THIS IS GENERATED CODE BY WBAPPTESTER
// you will need to edit this code to suit your needs

{{range $funcinfo := .FuncInfo}} func Test{{$funcinfo.Name}} (t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name string
		ExpectedStatus int
		MuxVars map[string]string
	}{
		{
			Name: "{{$funcinfo.Name}}: valid test case",
			ExpectedStatus: http.StatusOK,
{{if $containsmux}}	MuxVars: map[string]string  {
				{{range $muxvar := $funcinfo.MuxVars}} "{{$muxvar}}" : "valid_value", 
				{{end}}
			},
			{{end}}
		},
		{
			Name: "{{$funcinfo.Name}}: invalid test case",
			ExpectedStatus: http.StatusBadRequest,
{{if $containsmux}}	MuxVars:        map[string]string  {
				{{range $muxvar := $funcinfo.MuxVars}} "{{$muxvar}}" : "invalid_value", 
				{{end}}
			},
			{{end}}
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T){
			t.Parallel()
			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc({{$funcinfo.Name}})

			{{if $containsmux}}
			req = mux.SetURLVars(req, tc.MuxVars)
			{{end}}

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
