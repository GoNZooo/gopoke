package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (p pokee) String() string {
	return fmt.Sprintf("%s [%s]", p.Name, p.Url)
}

func (pr pokeresult) String() string {
	return fmt.Sprintf("%s\t%d in %s", pr.Name, pr.Readsize, pr.Duration)
}

// pokee is a name for a ping site and a url to fetch.
type pokee struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// pokeresult is a listing of a name, bytes read and a fetch time
type pokeresult struct {
	Name     string        `json:"name"`
	Readsize int           `json:"readsize"`
	Duration time.Duration `json:"duration"`
}

// poke fetches a page and returns the amount of characters read and the time it took to fetch them.
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

// readpokees reads all declared pingsites from the given pokeefile and returns an array of pokees.
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

// PokeAll assembles all pingsites read from the given pokeefile and pings them, then returns the results.
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
