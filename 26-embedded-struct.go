package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Studentt struct {
	grade  int
	Person // memasukan struct person di struct studentt
}

func main() {
	var s1 Studentt
	s1.grade = 3
	s1.name = "lotso"
	s1.age = 10

	fmt.Printf("name:%s\n", s1.name)
	fmt.Printf("age:%d\n", s1.age)
	fmt.Printf("grade:%d\n", s1.grade)

}
