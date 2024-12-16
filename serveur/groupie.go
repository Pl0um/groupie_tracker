package engine

import (
    "html/template"
    "net/http"
)
func groupieHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("template/groupie.html"))
    tmpl.Execute(w, nil)
}