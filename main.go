package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/gin-gonic/gin"
)

type APIResponse struct {
    Artists   interface{} `json:"artists"`
    Locations interface{} `json:"locations"`
    Dates     interface{} `json:"dates"`
    Relation  interface{} `json:"relation"`
}

func fetchData(url string) (interface{}, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var result interface{}
    err = json.Unmarshal(body, &result)
    if err != nil {
        return nil, err
    }

    return result, nil
}

func getAPIData(c *gin.Context) {
    artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
    locationsURL := "https://groupietrackers.herokuapp.com/api/locations"
    datesURL := "https://groupietrackers.herokuapp.com/api/dates"
    relationURL := "https://groupietrackers.herokuapp.com/api/relation"

    artists, err := fetchData(artistsURL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching artists: %v", err)})
        return
    }

    locations, err := fetchData(locationsURL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching locations: %v", err)})
        return
    }

    dates, err := fetchData(datesURL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching dates: %v", err)})
        return
    }

    relation, err := fetchData(relationURL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching relation: %v", err)})
        return
    }

    response := APIResponse{
        Artists:   artists,
        Locations: locations,
        Dates:     dates,
        Relation:  relation,
    }

    c.JSON(http.StatusOK, response)
}

func main() {
    router := gin.Default()
    router.GET("groupietrackers.herokuapp.com/api", getAPIData)

	router.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
    })

	fmt.Println("Server starting on http://localhost:8080")
    router.Run("localhost:8080")
}