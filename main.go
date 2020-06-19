package main

import (
	"fmt"
	"os"

	"github.com/GoNZooo/gopoke/poker"
)

func pokeefilepath() string {
	if os.Getenv("GOPOKE_POKEES") == "" {
		return os.Getenv("HOME") + "/.local/share/gopoke/pokees.json"
	}

	return os.Getenv("GOPOKE_POKEES")
}

func main() {
	pokeResults := poker.PokeAll(pokeefilepath())

	fmt.Println("Site\t\t\tBytes read\t\tTime to read")
	fmt.Println("------------------------------------------------------------")
	for pokeResult := range pokeResults {
		fmt.Printf("%s\n", pokeResult)
	}
}
