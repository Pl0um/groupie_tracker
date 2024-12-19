package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Run(jeu *Engine) {
    http.HandleFunc("/", jeu.Handler)
    http.HandleFunc("/groupie", groupieHandler)
    http.HandleFunc("/api/artists", serveArtists)

    fmt.Println("Server is listening on port http://localhost:2025")
    if err := http.ListenAndServe(":2025", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func serveArtists(w http.ResponseWriter, r *http.Request) {
    artists, err := GetArtists()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(artists)
}