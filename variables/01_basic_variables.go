package main

import "fmt"

func main(){
	var apple int
	apple = 11

	var mango float64
	mango = 11.22

	var lichi bool
	lichi = true

	var pine string
	pine = "rose"

	fmt.Printf("%v %.2f  %v  %q\n", apple, mango, lichi, pine)
}
