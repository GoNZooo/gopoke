package main

import (
    "fmt"
    "gopoke/poker"
)

func main() {
	prs := poker.PokeAll("pokees.json")

	for _, pr := range prs {
		fmt.Printf("%s\n", pr)
	}
}
