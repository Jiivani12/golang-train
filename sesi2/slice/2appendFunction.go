package main

import "fmt"

func main() {

	//var fruits = make([]string, 3)
	//_ = fruits

	//fruits[0] = "apple"
	//fruits[1] = "banana"
	//fruits[2] = "mango"

	//fmt.Printf("%#v", fruits)

	//append function2

	//var fruits = make([]string, 3)

	//fruits = append(fruits, "apple", "banana", "mango")

	//fmt.Printf("%#v", fruits)

	//append function with ellipsis
	var fruits1 = []string{"apple", "banana", "mango"}

	var fruits2 = []string{"durian", "pineapple", "starfruit"}

	fruits1 = append(fruits1, fruits2...)

	fmt.Printf("%#v", fruits1)

}
