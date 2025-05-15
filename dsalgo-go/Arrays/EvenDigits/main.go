package main

import (
	"fmt"
	"math"
	"strconv"
)

// https://leetcode.com/problems/find-numbers-with-even-number-of-digits/
func main() {

	fmt.Println("printing numbers")
	elments := []int{12, 345, 2, 16, 7896, 134}
	fmt.Println(findNumbers(elments))
	fmt.Println(CountDigits2(1000))
}

func findNumbers(nums []int) int {
	integerCount := 0
	for _, num := range nums {
		if checkEven(num) {
			integerCount++
		}
	}
	return integerCount
}

func checkEven(num int) bool {
	if CountDigits2(num)%2 == 0 {
		return true
	}
	return false
}

// UsingLog - optimised version
func CountDigits2(num int) int {
	return int(math.Log10(float64(num))) + 1
}

func CountDigits(num int) int {
	countDigits := 0

	if num == 0 {
		return 1
	}
	for {
		if num == 0 {
			break
		}
		num = num / 10
		countDigits++

	}
	return countDigits
}

func checkEven2(num int) bool {
	str := strconv.Itoa(num)
	if len(str)%2 == 0 {
		return true
	}
	return false
}
