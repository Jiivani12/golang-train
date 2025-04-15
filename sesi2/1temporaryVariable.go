package main

import "fmt"

func main() {
	var currentYear = 2025

	if age := currentYear - 2009; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim") //kalo hasil dari currentYear - n = lebih dari 17, sudah boleh biki
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim") //tetapi kalau hasil dari currentYear - n = kurang dari 17, belum boleh bikin
	}
}
