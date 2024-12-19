package engine

import (
	"html/template"
	"net/http"
)

func (base *Engine) Home(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier Home.html
	tmpl := template.Must(template.ParseFiles("template/Home.html"))
	
	// J'execute le template avec les données
	tmpl.Execute(w, nil)

		// Je redirige vers la page Groupie.html
	if r.Method == "POST" {
		buttonValue := r.FormValue("button")
		if buttonValue == "groupie" {
				http.Redirect(w, r, "/Groupie", http.StatusFound)
		}
	}
}