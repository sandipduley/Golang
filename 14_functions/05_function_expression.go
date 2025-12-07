package main

import ("fmt")

// When we store a function inside a "variable" that function is known as Function Expression  
func main() {
	add := func(num1 int, num2 int) int {
		return num1 + num2
	}

	result := add(11, 22)
	fmt.Println(result)
}
