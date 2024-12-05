package main

import "fmt"

func main() {
	var students = map[int]string{
		1: "lotso", // key 1 & value: lotso dan dipisahkan oleh ":"
		2: "huggin",
		3: "bear",
	}

	fmt.Printf("%s\n", students[1]) // cara mengakses value pada map, menggunakan key

	fmt.Println("==================================================================")

	for _, student := range students {
		fmt.Printf("%s\n", student)
	}

	fmt.Println("==================================================================")

	delete(students, 2) // menghapus value "huggin" dengan key: 2

	for _, student := range students {
		fmt.Printf("%s\n", student)
	}

	fmt.Println("==================================================================")

	value, isExist := students[2] // mencari value dengan key, variable isExist bersifat opsional
	if isExist {
		fmt.Printf("%s\n", value)
	} else {
		fmt.Println("Tidak ada")
	}

	fmt.Println("==================================================================")

	// slice of map
	fruits := []map[string]string{
		{"name": "grape"},
		{"name": "orange"},
		{"name": "guava"},
	}

	for _, fruit := range fruits {
		fmt.Printf("%s\n", fruit["name"]) // mengakses tiap map dengan key "name"
	}

}
