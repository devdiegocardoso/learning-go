package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	fmt.Println(states)
	states["PR"] = "Parana"
	states["SC"] = "Santa Catarina"
	states["RS"] = "Rio Grande do Sul"
	fmt.Println(states)

	fmt.Println(states["PR"])
	parana := states["PR"]
	fmt.Println(parana)

	delete(states, "SC")
	states["SP"] = "Sao Paulo"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}

	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
		i++
	}
}
