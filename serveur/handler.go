package engine

import (
	"html/template"
	"net/http"
)

func (jeu *Engine) Handler(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
	tmpl := template.Must(template.ParseFiles("template/Home.html"))

	// Définir les données à passer au template
	data := map[string]interface{}{
		"title": "Welcome to Groupie Tracker",
	}

	// J'execute le template avec les données
	tmpl.Execute(w, data)
}
func (jeu *Engine) Request(w http.ResponseWriter, r *http.Request, html string) {
	boutonHome := r.Form.Get("bouton_home")
	if boutonHome == "groupie" {
		jeu.Handler(w, r)
	}
}