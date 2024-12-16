package engine

import (
    "fmt"
    "net/http"
    "html/template"
)

func Run(jeu *Engine) {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)

    fmt.Println("Server is listening on port http://localhost:3002")
    if err := http.ListenAndServe(":3002", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("/home/mael/go/groupie_tracker/template/Home.html"))
    tmpl.Execute(w, nil)

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the about page.")
}
