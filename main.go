package poker

import (
    "fmt"
)

func main() {
	prs := PokeAll("pokees.json")

	for _, pr := range prs {
		fmt.Printf("%s\n", pr)
	}
}
