package main

import "fmt"

func main() {
outerloop:
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan ke - ", i+1)
		for j := 0; j < 5; j++ {
			if i == 4 {
				break outerloop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}
