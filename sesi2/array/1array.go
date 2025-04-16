package main

import "fmt"

func main() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var strings = [3]string{"Airell", "Nanda", "Mailo"}

	fmt.Printf("%#v\n", numbers) //verb %#v untuk memformat tipe data array
	fmt.Printf("%#v\n", strings)
}
