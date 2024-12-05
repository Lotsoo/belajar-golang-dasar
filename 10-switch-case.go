package main

import (
	"fmt"
)

func main() {
	score := 6 // nilai score 6

	switch { // sama seperti if else, beda gaya syntax
	case (score == 10):
		fmt.Println("Lulus dengan nilai sempurna")
		break
	case (score > 5):
		fmt.Println("Lulus")
	default:
		{
			fmt.Println("Tidak lulus")

		}
	}
}
