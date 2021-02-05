package main

import (
	"os"
	"text/template"
)

func main() {
	root := template.New("root")
	tpl, err := root.ParseFiles("templates/root", "templates/inner")
	if err != nil {
		panic(err)
	}
	if err := tpl.Execute(os.Stdout, nil); err != nil {
		panic(err)
	}
}
