package main

import "fmt"

func main() {
	var score = 0

	switch score {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	//with operators
	switch {
	case score == 7:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}

	}

	//fallthrough
	switch {

	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score > 4):
		fmt.Println("not bad")
		fallthrough
	case score == 4:
		fmt.Println("it is ok, but please study harder")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you don't have a good score yet")
		}
	}

}
