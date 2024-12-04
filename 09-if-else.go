package main

import "fmt"

func main() {
	score := 6 // nilai score 6

	if score == 10 { //  6 = 10 ? false
		fmt.Println("Lulus dengan nilai sempurna")
	} else if score > 5 { // 6 > 5 ? true
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak lulus")
	}
}
