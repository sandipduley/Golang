package main

import "fmt"

func main(){
	thisIsFloat := 11.22
	changingToInt := int64(thisIsFloat)

	fmt.Println(changingToInt)
	fmt.Printf("%T \n", changingToInt)
}
