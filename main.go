package main

import (
    e "engine/serveur"
    "io/ioutil"
    "net/http"
)

func GetApi(url string) []byte {
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    return body
}

func main() {
    http.HandleFunc("/api/artists", func(w http.ResponseWriter, r *http.Request) {
        res := GetApi("https://groupietrackers.herokuapp.com/api/artists")
        w.Header().Set("Content-Type", "application/json")
        w.Write(res)
    })
    var jeu e.Engine
    e.Run(&jeu)
}
