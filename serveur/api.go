package engine

import (
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
