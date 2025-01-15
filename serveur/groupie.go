package engine

import (
	"html/template"
	"net/http"
)

func (base *Engine) Groupie(w http.ResponseWriter, r *http.Request) {
	// Définir le chemin complet vers le fichier Groupie.html
	tmpl := template.Must(template.ParseFiles("template/Groupie.html"))

	// Je crée une structure de données pour les données du template
	data := Engine{}

	// J'execute le template avec les données
	tmpl.Execute(w, data)
}

