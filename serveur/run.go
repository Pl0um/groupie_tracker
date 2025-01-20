package engine

import (
    "fmt"
    "net/http"
    "strings"
)

func Run(base *Engine) {
    // Je définis plusieurs routes
    http.HandleFunc("/", base.Handler) // Une fois lancé, on arrive sur la racine où on appelle la fonction Handler
    http.HandleFunc("/Home", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "template/Home.html")
    })
    http.HandleFunc("/Groupie", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "template/Groupie.html")
    })

    http.HandleFunc("/api/artists/", func(w http.ResponseWriter, r *http.Request) {
        id := strings.TrimPrefix(r.URL.Path, "/api/artists/")
        res := GetApi("https://groupietrackers.herokuapp.com/api/artists/" + id)
        w.Header().Set("Content-Type", "application/json")
        w.Write(res)
    })
    
    // Servir les fichiers CSS
    css := http.FileServer(http.Dir("css"))
    http.Handle("/css/", http.StripPrefix("/css/", css))

    // Servir les fichiers images
    images := http.FileServer(http.Dir("images"))
    http.Handle("/images/", http.StripPrefix("/images/", images))

    fmt.Println("Le serveur s'est lancé ici : http://localhost:2025") // J'affiche un message pour dire que le serveur s'est lancé
    if err := http.ListenAndServe(":2025", nil); err != nil { // Je lance le serveur sur le port 2025
        fmt.Println("Désolé, le serveur ne s'est pas lancé :", err) // Si il y a une erreur, je l'affiche
    }
}