package engine

import (
	"net/http"
	"html/template"
)

func (base *Engine) Home(w http.ResponseWriter, r *http.Request) {
		// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier Home.html
		tmpl := template.Must(template.ParseFiles("template/Home.html"))
	
		// Je crée une structure de données pour les données du template
		data := Engine{}
	
		// J'execute le template avec les données
		tmpl.Execute(w, data)

	
		// Je redirige vers la page Groupie.html
	if r.Method == "POST" {
		buttonValue := r.FormValue("button")
		if buttonValue == "groupie" {
				http.Redirect(w, r, "/Groupie", http.StatusFound)
		} else if buttonValue == "credit" {
				http.Redirect(w, r, "/Credit", http.StatusFound)
		} else if buttonValue == "location" {
			http.Redirect(w, r, "/Locations", http.StatusFound)
		}
	}
}
