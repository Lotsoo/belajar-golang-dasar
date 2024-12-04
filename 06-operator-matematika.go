package main

import "fmt"

func main() {
	a := 6
	b := 3

	penjumlahan := a + b
	fmt.Printf("Penjumlahan %d + %d = %d\n", a, b, penjumlahan)

	pengurangan := a - b
	fmt.Printf("Pengurangan %d - %d = %d\n", a, b, pengurangan)

	perkalian := a * b
	fmt.Printf("Perkalian %d * %d = %d\n", a, b, perkalian)

	pembagian := a / b
	fmt.Printf("Pembagian %d / %d = %d\n", a, b, pembagian)

	modulus := a % b
	fmt.Printf("Penjumlahan %d modulus %d = %d", a, b, modulus) // hasil bagi

}
