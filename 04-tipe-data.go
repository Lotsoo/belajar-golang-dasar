package main

import "fmt"

func main() {
	name := "lotso" // string
	fmt.Printf("Name: %s\n", name)

	age := 10 // int
	fmt.Printf("Age: %d\n", age)

	height := 185.543 // float
	fmt.Printf("Height: %f\n", height)
	fmt.Printf("Height: %.2f\n", height) // mengambil 2 nilai setelah koma

	isMarried := false // bool
	fmt.Printf("Is married: %t", isMarried)

}
