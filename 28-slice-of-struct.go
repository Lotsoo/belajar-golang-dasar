package main

import "fmt"

type Animal struct {
	name string
	isMamalia bool
}


func main() {
	listAnimals := []Animal{
		{name: "bird", isMamalia: false},
		{name: "dolphin", isMamalia: true},
	}

	for _, animal := range listAnimals {
		fmt.Printf("name: %v, is mamalia?: %v\n",animal.name,animal.isMamalia)
	}
}