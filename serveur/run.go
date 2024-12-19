package engine

import (
    "fmt"
    "net/http"
)

func Run(jeu *Engine) {
    http.HandleFunc("/", jeu.Handler)
    http.HandleFunc("/groupie", groupieHandler)
    http.HandleFunc("/Credit", CreditHandler)
    fmt.Println("Server is listening on port http://localhost:3002")
    if err := http.ListenAndServe(":3002", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
