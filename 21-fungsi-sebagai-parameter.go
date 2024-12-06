package main

import "fmt"

func main() {

	angka1, angka2 := 5, 2

	hasilPenjumlahan := operasiMatematika(angka1, angka2, tambah) // memanggil fungsi tambah untuk mengisi callback

	hasilPengurangan := operasiMatematika(angka1, angka2, func(i1, i2 int) int { // membuat closure atau anonymous func
		return angka1 - angka2
	})

	fmt.Printf("Hasil penjumlahan %v + %v = %v\n", angka1, angka2, hasilPenjumlahan)
	fmt.Printf("Hasil pengurangan %v - %v = %v", angka1, angka2, hasilPengurangan)

}

func operasiMatematika(angka1, angka2 int, callback func(int, int) int) int { // membuat fungsi didalam parameter
	return callback(angka1, angka2) // memanggil parameter callback untuk dieksekusi
}

func tambah(angka1, angka2 int) int { // membuat fungsi tambah
	return angka1 + angka2
}
