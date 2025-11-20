package main
import ("fmt")

func main() {
	fruit := [4]string{"apple", "mango", "pineapple", "lichi"}

	for frt := 0; frt < len(fruit); frt++ {
		// fmt.Println(frt)  // 0  1  2  3

		fmt.Println(frt, fruit[frt])  // apple  mango  pineapple  lichi  
	}
}
