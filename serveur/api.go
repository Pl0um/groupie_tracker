package engine

import (
    "io/ioutil"
    "log"
    "net/http"
)

func GetApi(url string) []byte {
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    return responseData
}