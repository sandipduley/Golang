package main

import ("fmt")

func main() {
	// fmt.Printf()  --  Prints a formatted string std out
	// fmt.Sprintf()  --  Returns the formatted strings string as a value

	fmt.Printf("I am %v years old", 11)
	
	msg := fmt.Sprintf(" Hello %s, you scored %d marks!", "Sandip", 95)
	fmt.Println(msg)

}

