package main
import ("fmt")

func main() {
	fruit := map[string]int {
		"lichi": 11,
		"guava": 22,
		"strawberry":33,
	}
	
	for prop, value := range fruit {
		fmt.Println(prop, value)
	}
}
