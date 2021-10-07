package main

import (
	"fmt"
	"sort"
	"strconv"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func sortArray(array []int, chanForSlices chan []int) {
	sort.Ints(array)
	chanForSlices <- array
}

func merge2Slices(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}
	// if there is something in left slice that not in 'result' array
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	// if there is something in right slice that not in 'result' array
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}

func main() {
	var testArray []int
	var input string

	for true {
		fmt.Println("Enter an integer to add it to the testArray or 'X' to exit entering the array elements:")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("[ERROR] Incorrect input! Can not scan the input...")
		} else {
			if input == "X" {
				fmt.Println("Exit")
				break
			} else {
				item, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("[ERROR] Incorrect input! Input is not an integer...")
				} else {
					testArray = append(testArray, item)
				}
			}
		}
	}

	size := len(testArray)
	initialSize := float64(size)
	subArraySize := initialSize / 4
	subArraySizeRounded := int(subArraySize)
	roundSurplus := subArraySize - float64(subArraySizeRounded)

	fmt.Println("Initial array: ", testArray)
	chanForSlices := make(chan []int, 4)

	sizeAlign := 0.0
	// sizeAlign and roundSurplus vars are for the situation then we have an array with size(length) that can not divide on 4 without rest(such as 15 or 19)
	for i := 0; i < size; i += subArraySizeRounded {
		sizeAlign += roundSurplus
		if sizeAlign >= 1.0 {
			subArraySizeRounded++
		}

		go sortArray(testArray[i:min(i+subArraySizeRounded, size)], chanForSlices)

		if sizeAlign >= 1.0 {
			subArraySizeRounded--
			i++
			sizeAlign -= 1.0
		}
	}
	part1 := <-chanForSlices
	part2 := <-chanForSlices
	part3 := <-chanForSlices
	part4 := <-chanForSlices
	fmt.Println("Sorted parts: ", part1, part2, part3, part4)

	final := merge2Slices(merge2Slices(part1, part2), merge2Slices(part3, part4))
	fmt.Println("Sorted array: ", final)
}
