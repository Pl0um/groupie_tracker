package engine

import (
    "fmt"
    "net/http"
)

func Run(base *Engine) {
    // Je définis plusieurs routes
    http.HandleFunc("/", base.Home) // Une fois lancé, on arrive sur la racine où on appelle la fonction Handler
    http.HandleFunc("/Groupie", Groupie)
    http.HandleFunc("/Credit", Credit)

    // Fournir les fichiers template
    fs := http.FileServer(http.Dir("template/"))
    http.Handle("/serv/", http.StripPrefix("/template/", fs))

    // Fournir les fichiers CSS
    css := http.FileServer(http.Dir("css"))
    http.Handle("/css/", http.StripPrefix("/css/", css))

    // Fournir les fichiers JS
    js := http.FileServer(http.Dir("js"))
    http.Handle("/js/", http.StripPrefix("/js/", js))

    fmt.Println("Le serveur s'est lancé ici : http://localhost:2026") // J'affiche un message pour dire que le serveur s'est lancé
    if err := http.ListenAndServe(":2026", nil); err != nil { // Je lance le serveur sur le port 2025
        fmt.Println("Désolé, le serveur ne s'est pas lancé :", err) // Si il y a une erreur, je l'affiche
    }
}