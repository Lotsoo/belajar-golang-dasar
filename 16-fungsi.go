package main

import "fmt"

func main() {
	name := "lotso"
	sayHello(name) // mengirim argumen/data pada func sayHello
}

func sayHello(message string) { // menampung parameter message
	fmt.Println(message) // cetak parameter message
}
