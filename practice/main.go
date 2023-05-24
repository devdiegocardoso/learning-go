package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Age)
	poodle.Age = 9
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Age)

}

type Dog struct {
	Breed string
	Age   int
}
