package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetApi(url string) []byte {
	artists, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	res, err := http.DefaultClient.Do(artists)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func GetApi2(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData
}

func retApi(form string) int {
	res := GetApi("https://groupietrackers.herokuapp.com/api/artists")
	var artists []struct {
		Id      int      `json:"id"`
		Name    string   `json:"name"`
		Members []string `json:"members"`
	}
	json.Unmarshal(res, &artists)
	for _, val := range artists {
		if form == val.Name {
			return val.Id
		}
		for a := 0; a < len(val.Members); a++ {
			if form == val.Members[a] {
				return val.Id
			}
		}
	}
	return 0
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func GetArtists() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	body := GetApi(url)

	var artists []Artist
	err := json.Unmarshal(body, &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}
