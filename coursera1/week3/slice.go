package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
    inputItems := []int{}
	var input string
	looping := true

	for looping {
   	    fmt.Println("Enter a integer number to add it or 'X' to exit:")
        _, err := fmt.Scan(&input)
        if err != nil {
            fmt.Println("[ERROR] Incorrect input! Can not scan the input...")
        } else{
            if input == "X" {
                fmt.Println("Exit")
                break
            } else {
                item, err := strconv.Atoi(input)
                if err != nil {
                    fmt.Println("[ERROR] Incorrect input! Input is not an integer...")
                } else {
                    inputItems = append(inputItems, item)
                    sort.Ints(inputItems)
                    fmt.Printf("Your sorted slice: %v\n", inputItems)
                }
            }
        }
    }
}