package main

import (
	"log"
	"os"
	"text/template"
)

type stagiaire struct {
	Nom    string
	Prenom string
	Note   float32
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

var fm = template.FuncMap{
	"inc": func(i int) int {
		return i + 1
	},
}

func main() {

	stg := stagiaire{
		Nom:    "Rachid",
		Prenom: "Lahmami",
		Note:   16.98,
	}

	data := []stagiaire{
		stg,
		{
			Nom:    "Romaissaa",
			Prenom: "Lahmami",
			Note:   16.98,
		},
		{
			Nom:    "Imrane",
			Prenom: "Lahmami",
			Note:   14.96,
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
