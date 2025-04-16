package main

import "fmt"

//modify elements through index apel > apple, dst
func main() {
	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"

	fmt.Printf("%#v\n", fruits)
}
