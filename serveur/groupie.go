package engine

import (
	"html/template"
	"net/http"
)

func (base *Engine) groupieHandler(w http.ResponseWriter, r *http.Request) {
    // J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier Groupie.html
	tmpl := template.Must(template.ParseFiles("template/Groupie.html"))

    // J'exectute le template avec les données
	tmpl.Execute(w, nil)
}