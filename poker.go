package main

import (
    "net/http"
    "fmt"
    "log"
    "time"
    "io/ioutil"
    "encoding/json"
)

func (p pokee) String() string {
    return fmt.Sprintf("%s [%s]", p.Name, p.Url)
}

func (pr pokeresult) String() string {
    return fmt.Sprintf("%s\t%d\t%s", pr.Name, pr.Readsize, pr.Duration)
}

// A pokee is a name for a ping site and a url to fetch.
type pokee struct {
    Name string `json:"name"`
    Url string `json:"url"`
}

// A pokeresult is a listing of a name, bytes read and a fetch time
type pokeresult struct {
    Name string `json:"name"`
    Readsize int `json:"readsize"`
    Duration time.Duration `json:"duration"`
}

// Fetches a page and returns the amount of characters read and the time it took to fetch them.
func poke(p pokee) (result pokeresult) {
    start := time.Now()

    response, err := http.Get(p.Url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    result = pokeresult{p.Name, len(data), time.Now().Sub(start)}

    return
}

func readpokees(filename string) (pokees []pokee) {
    pokeedata, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    err = json.Unmarshal(pokeedata, &pokees)
    if err != nil {
        log.Fatal(err)
    }

    return
}

func PokeAll(pokeefile string) (results []pokeresult) {
    ps := readpokees(pokeefile)

    for _, p := range ps {
        results = append(results, poke(p))
    }

    return
}

func main() {
    prs := PokeAll("pokees.json")

    for _, pr := range prs {
        fmt.Printf("%s\n", pr)
    }
}
