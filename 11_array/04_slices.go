package main
import ("fmt")

func main() {
	mySlice1 := []int{}
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))
	fmt.Println(mySlice1)

	mySlice2 := []string{"apple", "mango", "banana", "pineapple"}
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))
	fmt.Println(mySlice2)

	fmt.Println(mySlice2[:])  // print all
	fmt.Println(mySlice2[1:3])  // print specific
	
	mySlice2 = append(mySlice2, "lichi", "guava")
	fmt.Println(mySlice2)
}
