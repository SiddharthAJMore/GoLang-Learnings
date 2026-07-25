package main

import "fmt"

// pass by value
func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum", num) // 5
}

func changeNumPointer(num *int) {
	*num = 5
	fmt.Println("In changeNumPointer", *num)
}

func main() {
	num := 1
	fmt.Println("before change", num) // 1
	changeNum(num)
	fmt.Println("after change", num) // 1

	// in go variables are passed as value to the function, and hence we cannot change the value of the original
	// 		variable inside the function.
	// But if we want to pass and change the variable inside the function, we need to send the reference
	// 		of the variable instead.

	fmt.Println("Memory location of num", &num)

	changeNumPointer(&num)
	fmt.Println("Memory location after changeNumPointer", num)

}
