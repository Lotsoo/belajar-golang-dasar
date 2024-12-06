package main

import "fmt"

type Student struct { // membuat struct
	name  string
	grade int
}

func main() {
	var s1 Student    // membuat variable bertipe student
	s1.name = "lotso" // mengakses properti name
	s1.grade = 3      // mengakses properti grade

	fmt.Printf("name:%v\n", s1.name)
	fmt.Printf("grade:%v\n", s1.grade)

	s2 := Student{ // inisialisasi bisa seperti ini
		name:  "bullseye",
		grade: 5,
	}

	fmt.Println("==================================================================")

	fmt.Printf("name:%v\n", s2.name)
	fmt.Printf("grade:%v\n", s2.grade)
}
