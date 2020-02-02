package main

const outputTemplate = `package {{.PackageName}}_test

import "testing"

{{range .FuncNames -}}
func Test{{.}} (t *testing.T) {
	testCases := []struct {
		name string
	}{
		{name: "{{.}}: valid test case"},
		{name: "{{.}}: invalid test case"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){

		})
	}
}
{{end -}}
`
