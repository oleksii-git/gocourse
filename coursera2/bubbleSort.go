package main

import (
    "fmt"
    "strconv"
)

func Swap(inputSlice *[]int, index int){
    swap := (*inputSlice)[index]
    (*inputSlice)[index] = (*inputSlice)[index-1]
    (*inputSlice)[index-1] = swap
}

func BubbleSort(input *[]int){
    for i := len(*input); i > 0; i-- {
        for j := 1; j < i; j++ {
            if (*input)[j-1] > (*input)[j] {
                Swap(input, j)
            }
        }
    }
}

func main() {
    testArray := []int{}
    var input string

    for i := 0; i < 10; i++ {
        fmt.Println("Enter a integer number to add it to array or 'X' to exit entering the array elements:")
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
                    testArray = append(testArray, item)
                }
            }
        }
    }


    fmt.Println("Input array: ", testArray)

    BubbleSort(&testArray)

    fmt.Println("Sorted array: ", testArray)
}