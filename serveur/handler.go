package engine

import (
    "encoding/json"
    "fmt"
    "html/template"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strconv"
)

// Function Home pour la page d'accueil
func (base *Engine) Home(w http.ResponseWriter, r *http.Request) {
    // Parcours le fichier template Home.html
    t, err := template.ParseFiles("./template/Home.html")
    if err != nil {
        // En cas d'erreur, renvoie une erreur interne du serveur
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // Effectue une requête GET pour récupérer les données des artistes
    response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
    if err != nil {
        // En cas d'erreur, affiche l'erreur et quitte le programme
        fmt.Print(err.Error())
        os.Exit(1)
    }

    // Lit le corps de la réponse
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        // En cas d'erreur, affiche l'erreur et quitte le programme
        log.Fatal(err)
    }
    // Déclare une variable pour stocker la liste des groupes
    var GroupList list
    // Décode les données JSON dans la variable GroupList
    json.Unmarshal(responseData, &GroupList.Lists)

    // Exécute le template avec les données des groupes
    t.Execute(w, GroupList)
}

// Function Groupie pour la page d'un groupe
func Groupie(w http.ResponseWriter, r *http.Request) {
    // Définir une map de fonctions pour être utilisées dans le template
    funcMap := template.FuncMap{
        "toJson": func(v interface{}) string {
            a, _ := json.Marshal(v)
            return string(a)
        },
        "safeJS": func(s string) template.JS {
            return template.JS(s)
        },
    }

    // Parcours le fichier template Groupie.html avec les fonctions définies
    t, err := template.New("Groupie.html").Funcs(funcMap).ParseFiles("./template/Groupie.html")
    if err != nil {
        // En cas d'erreur, renvoie une erreur interne du serveur
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Effectue une requête GET pour récupérer les données des artistes
    response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
    if err != nil {
        // En cas d'erreur, affiche l'erreur et quitte le programme
        fmt.Print(err.Error())
        os.Exit(1)
    }

    // Lit le corps de la réponse
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        // En cas d'erreur, affiche l'erreur et quitte le programme
        log.Fatal(err)
    }

    // Déclare une variable pour stocker la liste des groupes
    var GroupList []Engine
    // Décode les données JSON dans la variable GroupList
    json.Unmarshal(responseData, &GroupList)

    // Récupère l'ID de l'artiste à partir des paramètres de la requête
    id, _ := strconv.Atoi(r.FormValue("id"))
    artist := GroupList[id-1]

    // Convertit les locations de l'artiste en JSON
    locationsJSON, _ := json.Marshal(artist.Locations)

    // Exécute le template avec les données de l'artiste et les locations en JSON
    t.Execute(w, struct {
        Engine
        LocationsJSON string
    }{
        Engine:        artist,
        LocationsJSON: string(locationsJSON),
    })
}

// Function Crédit pour la page de crédit

func Credit(w http.ResponseWriter, r *http.Request) {
    // Parcours le fichier template Credit.html
    t, err := template.ParseFiles("./template/Credit.html")
    if err != nil {
        // En cas d'erreur, renvoie une erreur interne du serveur
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Exécute le template sans données supplémentaires
    t.Execute(w, nil)
}