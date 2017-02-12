package main

import (
	"fmt"
	"os"
	"sort"
)

const defaultArg = "slices"

type lang struct {
	name string
	rank int
}

func main() {
	fmt.Printf("Starting application with %v\n", os.Args)

	arg := defaultArg
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	switch arg {
	case "slices":
		fallthrough
	default:
		slices()
	}
}

func slices() {
	langs := []lang{
		{"Java", 1},
		{"Javascript", 7},
		{"C", 2},
		{"Go", 14},
		{"Python", 5},
	}
	fmt.Printf("Initial slice: %v\n", langs)

	sort.Slice(langs, func(i, j int) bool {
		return langs[i].rank < langs[j].rank
	})
	fmt.Printf("Slice sorted after rank: %v\n", langs)

	sort.Sort(byName(langs))
	fmt.Printf("Slice sorted after name: %v\n", langs)

}

type byName []lang

func (langs byName) Len() int {
	return len(langs)
}

func (langs byName) Less(i, j int) bool {
	return langs[i].name < langs[j].name
}

func (langs byName) Swap(i, j int) {
	langs[i], langs[j] = langs[j], langs[i]
}

