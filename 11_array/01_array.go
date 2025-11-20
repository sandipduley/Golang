package main
import ("fmt")

func main() {
	var array1 = [3]int{1, 2, 3}
	array2 := [5]int{11, 22, 33, 44, 55}
	array3 := [4]int{4, 5}
	array4 := [2]string{"apple", "mango"}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)

	fmt.Println(len(array4))  // length
	fmt.Println(cap(array3))  // capacity
}
