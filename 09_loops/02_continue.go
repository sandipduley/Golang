package main

import ("fmt")

func main() {
	for num := 2; num < 11; num++ {
		if num % 2 != 0 {  // skip iteration and continue
			continue
		}
	fmt.Println(num)
	}
}
