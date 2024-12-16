package main

import (
	e "engine/serveur"
	"net/http"
)

func main() {
	css := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	var jeu e.Engine
	e.Run(&jeu)
}