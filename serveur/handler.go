package engine

import (
	"html/template"
	"net/http"
)

func (base *Engine) Handler(w http.ResponseWriter, r *http.Request) {
	// J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier Home.html
	tmpl := template.Must(template.ParseFiles("template/Home.html"))
	
	// Je crée une structure de données pour les données du template
	data := Engine{}

	// J'execute le template avec les données
	tmpl.Execute(w, data)
}
