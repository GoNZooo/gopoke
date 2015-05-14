package main

import (
    "fmt"
    "gopoke/poker"
    "time"
    "os"
)

func pokeefilepath() (pokeefile string) {
    if os.Getenv("GOPOKE_POKEES") == "" {
        pokeefile = os.Getenv("HOME") + ".local/share/gopoke/pokees.json"
    } else {
        pokeefile = os.Getenv("GOPOKE_POKEES")
    }

    return
}

func main() {
    start := time.Now()
	prs := poker.PokeAll(pokeefilepath())
    diff := time.Now().Sub(start)

    fmt.Println("Site\t\t\tBytes read\tTime to read")
    fmt.Println("------------------------------------------------------------")
	for _, pr := range prs {
		fmt.Printf("%s\n", pr)
	}
    fmt.Printf("\nTotal time:\t\t%s\n", diff)
}
