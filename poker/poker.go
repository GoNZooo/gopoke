// Package poker is a package for fetching a specified set of websites
// and logging the response time of each site.
package poker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func (p pokee) String() string {
	return fmt.Sprintf("%s [%s]", p.Name, p.URL)
}

func scalesize(readsize int, unit string) (int, string) {
	if readsize < 1000 {
		return readsize, unit
	}

	switch unit {
	case "B":
		return scalesize(readsize/1024, "kB")
	case "kB":
		return scalesize(readsize/1024, "MB")
	case "MB":
		return scalesize(readsize/1024, "GB")
	}

	return 0, "B"
}

func (pr Pokeresult) String() string {
	readsizenumber, readsizeunit := scalesize(pr.Readsize, "B")

	nametabsize := 2 - (len(pr.Name) / 8)
	nametabs := ""
	readtabsize := 2 - (len(fmt.Sprintf("%d %s", readsizenumber, readsizeunit)) / 8)
	readtabs := ""

	for i := 0; i <= nametabsize; i++ {
		nametabs += "\t"
	}
	for i := 0; i <= readtabsize; i++ {
		readtabs += "\t"
	}

	return fmt.Sprintf("%s%s%d %s%s%s", pr.Name, nametabs, readsizenumber, readsizeunit, readtabs, pr.Duration)
}

// pokee is a name for a ping site and a url to fetch.
type pokee struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Pokeresult is a listing of a name, bytes read and a fetch time
type Pokeresult struct {
	Name     string        `json:"name"`
	Readsize int           `json:"readsize"`
	Duration time.Duration `json:"duration"`
}

// poke fetches a page and returns the amount of characters read and the time it took to fetch them.
func poke(p pokee, responsechannel chan Pokeresult, wg *sync.WaitGroup) {
	// Add 1 to wg; this increments lock counter
	// Deferred call is a decrementation of this counter
	wg.Add(1)
	defer wg.Done()

	internalchannel := make(chan Pokeresult)
	start := time.Now()

	go func() {
		response, err := http.Get(p.URL)
		datalength := 0
		if err != nil {
			datalength = -1
		} else {
			defer response.Body.Close()

			data, err := ioutil.ReadAll(response.Body)
			if err != nil {
				datalength = -1
			} else {
				datalength = len(data)
			}
		}

		internalchannel <- Pokeresult{p.Name, datalength, time.Now().Sub(start)}
	}()

	select {
	case result := <-internalchannel:
		responsechannel <- result
	case <-time.After(time.Second * 10):
		responsechannel <- Pokeresult{p.Name, -1, time.Now().Sub(start)}
	}
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
func PokeAll(pokeefile string) chan Pokeresult {
	// WaitGroup to know when to close the channel
	// channel will be closed when all sites are poked
	var wg sync.WaitGroup
	var results = make(chan Pokeresult)

	ps := readpokees(pokeefile)

	for _, p := range ps {
		go poke(p, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Return resultchannel
	return results
}
