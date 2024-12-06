package main

import "fmt"

func main() {
	number := 5

	fmt.Printf("before increment:%v\n", number)

	incrementNumber(&number)

	fmt.Println("==================================================================")

	fmt.Printf("after increment:%d\n", number)
}

func incrementNumber(number *int) {
	*number++
}
