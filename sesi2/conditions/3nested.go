package main

import "fmt"

func main() {
	var score = 6

	if score >= 7 {
		switch score {
		case 10: //score 10
			fmt.Println("Wohooo! Good Job!")
		default: //score 7-9
			fmt.Println("Awesome")
		}
	} else {
		if score >= 5 { //score 5-6
			fmt.Println("Ya, You good la")
		} else if score >= 3 { //score 3-4
			fmt.Println("you sure you doing your best?")
		} else { //score 1-2
			fmt.Println("stop playing around, keep ur head up!")
			if score == 0 { //score 0
				fmt.Println("see you next semester")
			}
		}
	}
}
