package main

import "fmt"

func main() {
	numbers := [3]int{1, 2, 3} // array

	fruits := []string{"Avocado", "Grape", "Banana"} // slice

	students := map[int]string{ // map
		1: "lotso",
		2: "Bullseye",
		3: "Forky",
	}

	// for range digunakan untuk melooping array, slice, map

	for i, fruit := range fruits { // for range wajib menggunakan 2 variable => index & value
		fmt.Printf("%d. %s\n", i, fruit)
	}

	fmt.Println("==================================================================")

	for _, number := range numbers { // jika tanpa index, menggunakan => _ (underscore)
		fmt.Println(number)
	}

	fmt.Println("==================================================================")

	for key, student := range students {
		fmt.Println("Key:", key, "Value:", student)
	}
}
