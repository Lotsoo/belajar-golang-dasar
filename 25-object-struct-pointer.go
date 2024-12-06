package main

import "fmt"

type Fruit struct {
	name string
}

func main() {
	f1 := Fruit{
		name: "Guava",
	}

	fmt.Printf("name:%v\n", f1.name)
	fmt.Println("==================================================================")

	var f2 *Fruit = &f1 // f2 pointing ke f1
	f2.name = "Carrot"  // mengganti value f2 tanpa harus menggunakan * (dereferensi)

	fmt.Printf("name:%v\n", f1.name)
	fmt.Printf("name:%v\n", f2.name)
}
