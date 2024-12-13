package engine


import (
    "encoding/json"
    "fmt"
    "io/ioutil"
	"github.com/gin-gonic/gin"
)

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

    response := Engine{
        Artists:   artists,
        Locations: locations,
        Dates:     dates,
        Relation:  relation,
    }

    c.JSON(http.StatusOK, response)
}