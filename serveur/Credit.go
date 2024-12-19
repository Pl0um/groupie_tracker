package engine

import (
    "html/template"
    "net/http"
)
func CreditHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("template/Credit.html"))
    tmpl.Execute(w, nil)
}
