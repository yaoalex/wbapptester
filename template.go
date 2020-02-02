package main

const outputTemplate = `package {{.PackageName}}

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// THIS IS GENERATED CODE BY WEBAPPTESTER
// you will need to edit this code to suit your API's needs

{{range .FuncNames -}}
func Test{{.}} (t *testing.T) {
	testCases := []struct {
		Name string
		ExpectedStatus int
	}{
		{
			Name: "{{.}}: valid test case",
			ExpectedStatus: http.StatusOK,
		},
		{
			Name: "{{.}}: invalid test case",
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T){
			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc({{.}})

			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tc.ExpectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.ExpectedStatus)
			}
		})
	}
}
{{end -}}
`
