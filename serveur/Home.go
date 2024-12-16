package engine

import (
	"html/template"
	"net/http"
)

func (jeu *Engine) Difficult(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier difficult.html
	tmpl := template.Must(template.ParseFiles("/home/mael/go/groupie_tracker/template/Home.html"))
	tmpl.Execute(w, nil)

		// Je redirige vers la page de jeu facile
		if r.Method == "POST" {
			buttonValue := r.FormValue("button")
			if buttonValue == "bouton_home" {
				http.Redirect(w, r, "/groupie", http.StatusFound)
			}
		}
	}