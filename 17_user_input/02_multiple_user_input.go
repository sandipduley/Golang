package main

import (
	"fmt"
)

func main() {
	var (
		whichFruit string
		quantity   float64
	)

	const perPiecePrice = 11.22

	fmt.Print("Enter the name of the fruit:- ")
	fmt.Scan(&whichFruit)

	fmt.Print("Enter quantity:- ")
	fmt.Scan(&quantity)

	totalAmount := quantity * perPiecePrice
	fmt.Println(totalAmount)
}
