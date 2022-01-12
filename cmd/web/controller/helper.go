package controller

import (
	"io"
	"net/http"
	"text/template"
)

func ShowPage(w io.Writer, data interface{}, pages ...string) {
	for i, p := range pages {
		pages[i] = "web/html/" + p
	}
	pages = append(pages,
		"web/html/base.layout.html",
		"web/html/logged.partial.html",
		"web/html/notify.partial.html",
	)
	tmpl := template.Must(template.ParseFiles(pages...))
	tmpl.Execute(w, data)
}

func AllowOrigins(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Add("Access-Control-Allow-Headers", "Content-Type, Accept")
}
