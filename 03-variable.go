package main

import "fmt"

func main() {
	var firstName string = "Lotso" // deklarasi variable + inisialisasi secara langsung

	var middleName string // deklarasi variable
	middleName = "Huggin"

	lastName := "Bear" // shorthand

	fmt.Printf("Halo %s %s %s!", firstName, middleName, lastName)
}
