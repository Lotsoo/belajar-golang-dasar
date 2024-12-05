package main

import "fmt"

func main() {
	fruits := []string{ // slice sama seperti array, bedanya slice fleksibel
		"avocado",
		"grape",
		"orange",
		"melon",
		"watermelon",
	}

	fmt.Println("Fruits:", fruits)
	fmt.Println("==================================================================")

	newFruits := fruits[0:2] // mengambil elemen dimulai dari index ke-0 hingga sebelum index ke-2
	fmt.Println("New fruits:", newFruits)
	fmt.Println("Panjang new fruits:", len(newFruits))
	fmt.Println("Kapasitas new fruits:", cap(newFruits)) // kapasitas reference dari fruits
	fmt.Println("==================================================================")

	newFruits2 := fruits[:] // mengambil semua elemen
	fmt.Println("New fruits2:", newFruits2)
	fmt.Println("==================================================================")

	newFruits2 = append(newFruits2, "Papaya") // menambah elemen
	fmt.Println("New fruits2:", newFruits2)

}
