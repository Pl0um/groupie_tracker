package engine

import (
    "html/template"
    "net/http"
)
func CreditHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("template/credit.html"))
    tmpl.Execute(w, nil)
}
