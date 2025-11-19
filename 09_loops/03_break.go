package main

import ("fmt")

func main () {
	for num := 0; num < 11; num++ {
		if num == 5 {
			break  // break loop when num = 5
		}
	  fmt.Println(num)
	}
}
