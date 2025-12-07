package main 

import ("fmt"; "os"; "bufio")

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name:- ")

	// comma ok || err ok 
	input, _ := reader.ReadString('\n')
	fmt.Println("Welcome,", input)
}
