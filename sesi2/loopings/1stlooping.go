package main

import "fmt"

func main() {
	//for i := 0; i < 3; I++ { //DO NOT TRY THE INFINITE LOOP IT WAS DISASTER LMAOO
		fmt.Println("angka", i)
	}

	//second way of looping
	var i = 0

	for i < 3 {
		fmt.Println("Angka", i)
		i++
	}

	//third way of looping
	var i = 0

	for {
		fmt.Println("Angka", i)

		i++
		if i == 8 {
			break
		}
	}
}
