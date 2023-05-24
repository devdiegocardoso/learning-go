package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value 1: ")
	v1input, _ := reader.ReadString('\n')
	v1, v1err := strconv.ParseFloat(strings.TrimSpace(v1input), 64)

	if v1err != nil {
		fmt.Println(v1err)
		panic("Value 1 must be a number!")
	}

	fmt.Print("Value 2: ")
	v2input, _ := reader.ReadString('\n')
	v2, v2err := strconv.ParseFloat(strings.TrimSpace(v2input), 64)

	if v2err != nil {
		fmt.Println(v1err)
		panic("Value 1 must be a number!")
	}

	v3 := v1 + v2
	v3 = math.Round(v3*100) / 100
	fmt.Println("The sum of", v1, "and", v2, "is", v3)

}
