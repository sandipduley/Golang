package main
import ("fmt")

func main() {
	fruit := []string{"apple", "mango", "banana"}

	for frt := 0; frt < len(fruit); frt++ {
		fmt.Println(fruit[frt])
	}

	// Alternative Way
	for index, value := range fruit {
		fmt.Println(index, value)
	} 
}
