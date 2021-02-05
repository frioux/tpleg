package main

import (
	"os"
	"text/template"
)

func main() {
	root, err := template.New("root").Parse(`
{{define "begin"}} this is the default begin {{end}}
{{define "end"}} this is the default end {{end}}
{{define "inner"}} this is the default inner {{end}}

BEGIN: {{template "begin"}}

INNER: {{template "inner" }}

END: {{template "end"}}`)
	if err != nil {
		panic(err)
	}
	tpl, err := root.ParseFiles("templates/inner")
	if err != nil {
		panic(err)
	}
	if err := tpl.Execute(os.Stdout, nil); err != nil {
		panic(err)
	}
}
