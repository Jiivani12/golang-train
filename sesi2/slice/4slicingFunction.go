package main

import "fmt"

func main() {

	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var fruits2 = fruits1[1:4] //Variable fruits2 ingin mendapatkan element dari fruits1 dari index ke satu hingga index ke tiga, maka dari itu cara penulisannya adalah fruits[1:4].
	fmt.Printf("%#v\n", fruits2)

	var fruits3 = fruits1[0:] //keterangan stop dihilangkan jika ingin mendapatkan element hingga index akhir
	fmt.Printf("%#v\n", fruits3)

	var fruits4 = fruits1[:3] //keterangan start dihilangkan jika ingin mendapatkan element dr index ke nol
	fmt.Printf("%#v\n", fruits4)

	var fruits5 = fruits1[:] //sama dengan fruits1[:len(fruits1)]
	fmt.Printf("%#v\n", fruits5)
}
