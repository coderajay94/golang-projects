package main

import "fmt"

func main() {

	fmt.Println("Hello World from Linear Search!!")

	elements := [9]int{12, 14, 1, 34, 61, 88, 56, -11, 11}

	fmt.Println(LinearSearch(elements, 11))

}

func LinearSearch(elements [9]int, item int) int {
	if len(elements) == 0 {
		return -1
	}

	for index, _ := range elements {
		if elements[index] == item {
			return index
		}
	}

	return -1

}
