package main

import "fmt"

func main() {
	panjang, lebar := 5, 2
	luasPersegiPanjang, kelilingPersegiPanjang := rumusPersegiPanjang(float32(panjang), float32(lebar)) // menampung 2 return
	fmt.Println("Luas persegi panjang:", luasPersegiPanjang)
	fmt.Println("Keliling persegi panjang:", kelilingPersegiPanjang)

	fmt.Println("==================================================================")

	sisi := 5
	luasPersegi, kelilingPersegi := rumusPersegi(float32(sisi))
	fmt.Println("Luas persegi:", luasPersegi)
	fmt.Println("Keliling persegi:", kelilingPersegi)

}

func rumusPersegiPanjang(p, l float32) (float32, float32) { // mendefinisikan tipe data pada return
	luas := p * l

	keliling := 2 * (p + l)

	return luas, keliling
}

func rumusPersegi(s float32) (luas, keliling float32) { // mendefinisikan 2 variable return
	luas = s * s
	keliling = 4 * s
	return // langsung return, go otomatis mengerti
}
