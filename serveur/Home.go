package engine

import (
	"html/template"
	"net/http"
)

func (jeu *Engine) Home(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier difficult.html
	tmpl := template.Must(template.ParseFiles("template/Home.html"))
	tmpl.Execute(w, nil)

		// Je redirige vers la page groupie
	if r.Method == "POST" {
		buttonValue := r.FormValue("button")
		if buttonValue == "bouton_home" {
				http.Redirect(w, r, "/groupie", http.StatusFound)
		}
	}
}