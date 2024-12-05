package main

import "fmt"

func main() {
	angka1, angka2 := 5, 2

	hasilPenjumlahan := hitungPenjumlahan(angka1, angka2)
	fmt.Printf("Penjumlahan %d + %d = %d\n", angka1, angka2, hasilPenjumlahan)
	
	fmt.Println("==================================================================")

	hasilPengurangan := hitungPengurangan(angka1, angka2)
	fmt.Printf("Pengurangan %d - %d = %d", angka1, angka2, hasilPengurangan)

}

func hitungPenjumlahan(angka1 int, angka2 int) int {
	hasil := angka1 + angka2
	return hasil
}

func hitungPengurangan(angka1, angka2 int) int  { // jika parameter memiliki tipe data yang sama bisa disingkat seperti itu
	hasil := angka1 - angka2
	return hasil
}