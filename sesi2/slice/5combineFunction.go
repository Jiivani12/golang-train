package main

import "fmt"

func main() {

	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	fruits1 = append(fruits1[:3], "rambutan")

	fmt.Printf("%#v\n", fruits1)

	//kalau mau menambahkan index rambutan diantara index lainnya pake yg bawah
	index := 3
	fruits1 = append(fruits1[:index], append([]string{"rambutan"}, fruits1[index:]...)...)

	fmt.Printf("%#v\n", fruits1)
}
