package main

import ("fmt")

func main() {
	password := "1234567890"

	if len(password) > 7 { fmt.Println("Valid Password!") 
	 } else { fmt.Println("Password length must be 8 characters!") }
}
