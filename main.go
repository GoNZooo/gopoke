package main

import (
	"fmt"
	"os"

	"github.com/GoNZooo/gopoke/poker"
)

func pokeefilepath() (pokeefile string) {
	if os.Getenv("GOPOKE_POKEES") == "" {
		pokeefile = os.Getenv("HOME") + "/.local/share/gopoke/pokees.json"
	} else {
		pokeefile = os.Getenv("GOPOKE_POKEES")
	}

	return
}

func main() {
	prs := poker.PokeAll(pokeefilepath())

	fmt.Println("Site\t\t\tBytes read\t\tTime to read")
	fmt.Println("------------------------------------------------------------")
	for pr := range prs {
		fmt.Printf("%s\n", pr)
	}
}
