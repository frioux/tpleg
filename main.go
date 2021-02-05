package main

import (
	"os"
	"text/template"
)

var order int

func Order() int { order++; return order }

func main() {
	root := template.New("root").Funcs(template.FuncMap{"order": Order})
	tpl, err := root.ParseFiles("templates/root", "templates/inner")
	if err != nil {
		panic(err)
	}
	if err := tpl.Execute(os.Stdout, nil); err != nil {
		panic(err)
	}
}
