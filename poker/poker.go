// poker is a package for fetching a specified set of websites
// and logging the response time of each site.
package poker

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

func (pr Pokeresult) String() string {
    nametabsize := 2 - (len(pr.Name) / 8)
    nametabs := ""
    readtabsize := 2 - (len(fmt.Sprintf("%d", pr.Readsize)) / 8)
    readtabs := ""

    for i := 0 ; i <= nametabsize ; i++ {
        nametabs += "\t"
    }
    for i := 0 ; i <= readtabsize ; i++ {
        readtabs += "\t"
    }

	return fmt.Sprintf("%s%s%d%s%s", pr.Name, nametabs, pr.Readsize, readtabs, pr.Duration)
}

// pokee is a name for a ping site and a url to fetch.
type pokee struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// Pokeresult is a listing of a name, bytes read and a fetch time
type Pokeresult struct {
	Name     string        `json:"name"`
	Readsize int           `json:"readsize"`
	Duration time.Duration `json:"duration"`
}

// poke fetches a page and returns the amount of characters read and the time it took to fetch them.
func poke(p pokee) (result Pokeresult) {
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

	result = Pokeresult{p.Name, len(data), time.Now().Sub(start)}

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
func PokeAll(pokeefile string) (results []Pokeresult) {
	ps := readpokees(pokeefile)

	for _, p := range ps {
		results = append(results, poke(p))
	}

	return
}

