package main
import ("fmt")

func main() {
	slice := make([]string, 3, 5)
	fmt.Println("length", len(slice))
	fmt.Println("capacity ", cap(slice))
	fmt.Println(slice)
}
