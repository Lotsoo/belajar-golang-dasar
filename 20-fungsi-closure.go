package main

import "fmt"

func main() {
	fmt.Println(penjumlahan(1, 2))
}

var penjumlahan = func(angka1, angka2 int) int { // closure => fungsi tanpa nama namun disimpan pada variable
	hasil := angka1 + angka2
	return hasil
}
