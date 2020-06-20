package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/GoNZooo/gopoke/poker"
)

func getPokeeSpecifications(hostsFromArgs []string) []string {
	if len(hostsFromArgs) == 0 {
		filename := os.Getenv("HOME") + "/.local/share/gopoke/pokees.txt"

		fileContents, err := ioutil.ReadFile(filename)
		if err != nil {
			panic("Unable to open pokee file and no hosts given\n")
		}

		s := strings.ReplaceAll(string(fileContents), "\r\n", "\n")

		return strings.Split(s, "\n")
	}

	return hostsFromArgs
}

func main() {
	hosts := getPokeeSpecifications(os.Args[1:])
	if len(hosts) == 0 {
		fmt.Print("No sites specified.\n")

		os.Exit(1)
	}
	fmt.Println(hosts)

	pokeResults := poker.PokeAll(hosts)

	fmt.Println("Site\t\t\tBytes read\t\tTime to read")
	fmt.Println("------------------------------------------------------------")
	for pokeResult := range pokeResults {
		fmt.Printf("%s\n", pokeResult)
	}
}
