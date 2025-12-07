package main

import ("fmt")

func takeNumbers(numbers ...int) {
	fmt.Println(numbers)
}

func hybridInput(anotherParam string, integer ...int) {
	fmt.Println(anotherParam, integer)
}

func main() {
	takeNumbers(11, 22, 33, 44, 55)  // [11 22 33 44 55]

	hybridInput("Apple", 66, 77, 88, 99, 00) // Apple [66 77 88 99 0]
}
