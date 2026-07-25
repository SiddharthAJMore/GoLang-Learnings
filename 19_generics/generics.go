package main

import "fmt"

func printElements[T int | string](elements []T) {
	// following generic declarations can be used
	//  [T any], [T interface{}], [T anyotherType], [T int | string | otherType1 | otherType2], [T comparable] //comparable internally includes int, boolean, etc. general data types

	for _, element := range elements {
		fmt.Println(element)
	}
}

func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

type stack[T any] struct {
	elements []T
}

func main() {
	//nums := []int{1, 2, 3, 4, 5, 6}

	names := []string{"sidd", "dev", "harsh"}
	printElements(names)

	fmt.Println("-------------")

	myStack := stack[int]{
		elements: []int{1, 2, 3},
	}

	myStack2 := stack[string]{
		elements: []string{"name1", "name2"},
	}

	fmt.Println(myStack)
	fmt.Println(myStack2)

	names2 := []string{"sidd2", "dev2", "harsh2"}
	printSlice(names2, "more")
}
