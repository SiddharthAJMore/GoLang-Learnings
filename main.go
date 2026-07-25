package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	// simple values
	//integer
	fmt.Println(1 + 1)
	//String
	fmt.Println("String")
	//floats
	fmt.Println(10.51)

	//boolean
	fmt.Println(true)

	//variables

	// var name string = "golang"
	// OR
	var name = "golang"
	fmt.Println(name)

	// while loop
	var i int = 0

	for i <= 3 { // while i is less than equals to 3
		fmt.Println(i)
		i++
	}

	fmt.Println("---1---")

	// while infinity
	//for {
	//	fmt.Println(1)
	//}

	// for loop normal
	for i := 1; i <= 4; i++ {
		fmt.Println(i)
	}

	fmt.Println("---2---")

	// for loop in range
	for i := range 3 {
		if i == 1 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("---3---")
	//switch statements

	var weekday string = "sunday"
	switch weekday {
	case "sunday", "saturday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	fmt.Println("---4---")

	checkDataType := func(j interface{}) {
		switch j.(type) {
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its a string")
		case bool:
			fmt.Println("Its a boolean")
		default:
			fmt.Println("other")
		}
	}

	checkDataType("wsc")

	fmt.Println("---Arrays---")

	var nums [2]int
	fmt.Println(nums)

	names := [3]string{"Siddharth", "AJ", "More"}
	fmt.Println(names)

	fmt.Println(names[0], "is good boy")

	fmt.Println("---Slices---")
	fmt.Println("---Map---")

	fmt.Println("---Functions---")

	result := add(1, 4)
	fmt.Println(result)

	lang1, lang2, _ := getLanguages() // _ will skip the storage of the lang3. Use incase when we dont want to use any one of the output
	fmt.Println(lang2, lang1)

	someFunc := func(a int) int {
		return 4
	}
	fmt.Println(processIt1(someFunc))

	fn := processIt2()
	fmt.Println(fn(1))

	fmt.Println("---Variadic functions---") // variadic functinos can get any number of params
	fmt.Println("Variadic function", sum(1, 2, 3, 4, 5, 6))

	numsForVariadic := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Variadic function spread opp", sum(numsForVariadic...))

	fmt.Println("---Closures---")

	increment := counter()
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(increment()) // 3
	//Ideally once the method is called then the variable used inside it are released after the method is completed.
	//	But incase of closure if the variable used inside the fuction returned is a global variable created outside the function,
	// 	then the value of the variable is stored and reused whenever the function is called.
}

func counter() func() int {
	var counter int = 0

	return func() int {
		counter += 1
		return counter
	}
}

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}
	return total
}

func processIt2() func(a int) int {
	return func(a int) int {
		return 2
	}
}

func processIt1(fn func(a int) int) int {
	return fn(1)
}

func add(a int, b int) int { // or (a,b int)
	return a + b
}

func getLanguages() (string, string, string) {
	return "golang", "java", "python"
}
