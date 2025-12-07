package main

import ("fmt")

func customer_names(names ...string) {
	for index, values := range names {
		fmt.Println(index, values)
	}
}

func main() {
	customer_names("Apple", "Mango", "Banana", "Pineapple", "Lichi", "Guava")
}

// 0 Apple
// 1 Mango
// 2 Banana
// 3 Pineapple
// 4 Lichi
// 5 Guava
