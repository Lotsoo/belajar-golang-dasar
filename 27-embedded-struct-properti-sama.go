package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	age   int 
	grade int
	person
}

func main() {
	var s1 student
	s1.age = 10 // mengakses properti age milik student
	s1.person.age = 15 // mengakses properti age milik person

	fmt.Println(s1.age)
	fmt.Println(s1.person.age)
}
