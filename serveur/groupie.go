package engine

import (
    "html/template"
    "net/http"
    "path/filepath"
)

func (base *Engine) Groupie(w http.ResponseWriter, r *http.Request) {
    // Définir le chemin complet vers le fichier Groupie.html
    tmplPath := filepath.Join("template", "Groupie.html")

    // Utiliser la bibliothèque template pour créer un modèle qui recherche mon fichier Groupie.html
    tmpl, err := template.ParseFiles(tmplPath)
    if err != nil {
        http.Error(w, "Erreur lors du chargement du modèle", http.StatusInternalServerError)
        return
    }

    // Exécuter le modèle avec des données
    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, "Erreur lors de l'exécution du modèle", http.StatusInternalServerError)
    }
}