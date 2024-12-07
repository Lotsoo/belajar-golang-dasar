package main

import "fmt"

type car struct {
	name string
}

func (c car) displayNameCar() { // (c car) syntax berfungsi untuk memberitahu bahwa ini struct
	fmt.Println(c.name)
}

func main() {
	bmw := car{name: "bmw"}
	bmw.displayNameCar() // memanggil method
}
