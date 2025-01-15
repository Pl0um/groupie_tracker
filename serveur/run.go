package engine

import (
	"fmt"
	"net/http"
	"strings"
)

func Run(base *Engine) {
    // Je définis plusieur routes
    http.HandleFunc("/", base.Handler) // Une fois lancer on arrive sur la racine ou on appelle la fonction Handler
    http.HandleFunc("/Home", base.Home)
    http.HandleFunc("/Groupie", base.Groupie)
    

    http.HandleFunc("/api/artists/", func(w http.ResponseWriter, r *http.Request) {
        id := strings.TrimPrefix(r.URL.Path, "/api/artists/")
        res := GetApi("https://groupietrackers.herokuapp.com/api/artists" + id)
        w.Header().Set("Content-Type", "application/json")
        w.Write(res)
    })
    
	css := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

    // Servir les fichiers images
    images := http.FileServer(http.Dir("images"))
    http.Handle("/images/", http.StripPrefix("/images/", images))

    // Servir le fichier Groupie.html
    http.HandleFunc("/Groupie.html", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "template/Groupie.html")
    })
    
    fmt.Println("Le serveur c'est lancer ici : http://localhost:2025") // J'affiche un message pour dire que le serveur c'est lancer
    if err := http.ListenAndServe(":2025", nil); err != nil { // Je lance le serveur sur le port 2025
        fmt.Println("Désoler le serveur ne c'est pas lancer :", err) // Si il y a une erreur je l'affiche
    }
}
