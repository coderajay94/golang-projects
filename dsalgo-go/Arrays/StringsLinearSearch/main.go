package main

import "fmt"

func main() {
	fmt.Println("lets test string linear search")

	str := "Ajay Kumar"

	index, found := StringLinearSearch(str, "a")
	fmt.Printf("String %v found at index: %v", found, index)
}

func StringLinearSearch(str string, findme string) (int, bool) {

	if len(str) == 0 {
		return -1, false
	}

	for index, val := range str {
		if string(val) == findme {
			return index, true
		}
	}
	return -1, false
}
