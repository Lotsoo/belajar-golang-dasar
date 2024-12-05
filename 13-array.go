package main

import "fmt"

func main() {

	var cars [2]string  // inisialisasi array dengan 2 elemen
	cars[0] = "Ioniq 5" // mengisi elemen dengan cara mengakses indexnya
	cars[1] = "wuling binguo ev"

	names := [3]string{"lotso", "huggin", "bear"} // [3] => isinya ada 3 elemen, array gaya horizontal

	fruits := [5]string{ // [5] => isinya ada 5 elemen, array gaya vertikal
		"avocado",
		"grape",
		"orange",
		"melon",
		"watermelon",
	}

	for i, car := range cars {
		fmt.Printf("%d. %s\n", i, car)
	}

	fmt.Println("==================================================================")

	for _, name := range names {
		fmt.Println(name)
	}

	fmt.Println("==================================================================")

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

}
