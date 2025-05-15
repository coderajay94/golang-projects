package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Welcome to the Search2dArray problem")

	elments2d := [][]int{{12, 22, 11, 14}, {53, 2, 13, 67}, {12, 34}, {23, 562}, {-1}, {}}
	target := 562

	row, column := SearchIn2dSlice(elments2d, target)
	fmt.Printf("item found at index row %v and column %v : %v \n", row, column, target)
	fmt.Println("min in 2d Array", MinIn2dSlice(elments2d))
	fmt.Println("max in 2d Array", MaxIn2dSlice(elments2d))

	//cleaner versions
	fmt.Println("min in 2d Array", MinIn2dSlice2(elments2d))
	fmt.Println("max in 2d Array", MaxIn2dSlice2(elments2d))

	//note: when element is searched you can use for i:=0; i<len(elements);i++
	//because index is what we're looking in 2d array in these scenarios
	//but max or min kind of solutions you can directly return with range like index, val := range elments
	//when only iteration is required range is enough
}

func SearchIn2dSlice(elements [][]int, target int) (int, int) {

	for row := 0; row < len(elements); row++ {
		for column := 0; column < len(elements[row]); column++ {

			if elements[row][column] == target {
				return row, column
			}
		}
	}
	return -1, -1
}

// concise version
func MaxIn2dSlice2(elments [][]int) int {
	max := math.MinInt

	for rowIndex, _ := range elments {
		for _, columnElement := range elments[rowIndex] {
			if columnElement > max {
				max = columnElement
			}
		}
	}
	return max
}

// concise version
func MinIn2dSlice2(elments [][]int) int {
	min := math.MaxInt

	for rowIndex, _ := range elments {
		for _, columnElement := range elments[rowIndex] {
			if columnElement < min {
				min = columnElement
			}
		}
	}
	return min
}

func MaxIn2dSlice(elements [][]int) int {
	max := math.MinInt
	for row := 0; row < len(elements); row++ {
		for column := 0; column < len(elements[row]); column++ {

			if elements[row][column] > max {
				max = elements[row][column]
			}
		}
	}
	return max
}

func MinIn2dSlice(elements [][]int) int {
	min := math.MaxInt
	for row := 0; row < len(elements); row++ {
		for column := 0; column < len(elements[row]); column++ {

			if elements[row][column] < min {
				min = elements[row][column]
			}
		}
	}
	return min
}
