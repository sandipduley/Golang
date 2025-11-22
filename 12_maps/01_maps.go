package main
import ("fmt")

func main() {
	fruit := map[string]int {
		"lichi": 11,
		"guava": 22,
		"strawberry":33,
	}

	fmt.Println(fruit)
	fruit["cherry"] = 44

	fmt.Println(fruit["lichi"])
	fmt.Println(fruit["guava"])
	fmt.Println(fruit["strawberry"])
	fmt.Println(fruit["cherry"])
}
