package main

import ("fmt")

func add(x int, y int) int {
	return x + y	
}

func main() {
	fmt.Println(add(11, 22))  // 33

	// another way
	result := add(33, 44)  // 77
	fmt.Println(result)
}
