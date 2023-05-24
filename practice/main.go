package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)
	colors = colors[1:]
	fmt.Println(colors)
	colors = colors[:len(colors)-1]
	fmt.Println(colors)

	numbers := make([]int, 5)
	numbers[0] = 123
	numbers[1] = 234
	numbers[2] = 345
	numbers[3] = 456
	numbers[4] = 567

	fmt.Println(numbers)
	numbers = append(numbers, 5)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}
