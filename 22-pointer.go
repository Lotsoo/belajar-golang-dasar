package main

import "fmt"

func main() {
	var numberA int = 5
	var numberB *int = &numberA // pointing ke numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)

	fmt.Println("numberB (value)   :", *numberB) // * digunakan untuk mengambil value yang dipointing
	fmt.Println("numberB (address) :", numberB)

	fmt.Println("==================================================================")

	numberA = 1 // mengubah nilai numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)

	fmt.Println("numberB (value)   :", *numberB) // value numberB akan berubah juga, karena pointing ke numberA
	fmt.Println("numberB (address) :", numberB)
}
