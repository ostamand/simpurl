package controller

import (
	"io"
	"text/template"
)

func ShowPage(w io.Writer, data interface{}, pages ...string) {
	for i, p := range pages {
		pages[i] = "ui/html/" + p
	}
	pages = append(pages,
		"ui/html/base.layout.html",
		"ui/html/logged.partial.html",
		"ui/html/notify.partial.html",
	)
	tmpl := template.Must(template.ParseFiles(pages...))
	tmpl.Execute(w, data)
}
