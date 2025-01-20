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

func (base *Engine) Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/Home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var GroupList list
	json.Unmarshal(responseData, &GroupList.Lists)
	t.Execute(w, GroupList)
}

func Groupie(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/Groupie.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var GroupList []Engine
	json.Unmarshal(responseData, &GroupList)

	id, _ := strconv.Atoi(r.FormValue("id"))
	t.Execute(w, GroupList[id-1])
}

func Credit(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/Credit.html"))
	tmpl.Execute(w, nil)
}