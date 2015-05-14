package main

import (
    "fmt"
    "gopoke/poker"
    "time"
)

func main() {
    start := time.Now()
	prs := poker.PokeAll("pokees.json")
    diff := time.Now().Sub(start)

    fmt.Println("Site\t\t\tBytes read\tTime to read")
    fmt.Println("------------------------------------------------------------")
	for _, pr := range prs {
		fmt.Printf("%s\n", pr)
	}
    fmt.Printf("\nTotal time:\t\t%s\n", diff)
}
