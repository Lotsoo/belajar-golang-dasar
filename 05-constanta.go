package main

import "fmt"

func main() {
	const firstName = "Lotso" // nilai constanta tidak bisa diubah
	fmt.Printf("First name: %s", firstName)

	firstName = "Bullseye" // error
	fmt.Printf("First name: %s", firstName)
}
