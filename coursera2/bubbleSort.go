package main

import "fmt"

func Swap(inputSlice *[10]int, index int){
    swap := (*inputSlice)[index]
    (*inputSlice)[index] = (*inputSlice)[index-1]
    (*inputSlice)[index-1] = swap
}

func BubbleSort(chanForSortedArray chan [10]int, input [10]int){
    for i := len(input); i > 0; i-- {
        for j := 1; j < i; j++ {
            if input[j-1] > input[j] {
                Swap(&input, j)
            }
        }
    }
    chanForSortedArray <- input
}

func main() {
    testArray := [10]int{18, 8, -3, -7, 4, 31, 7, 81, -23, 1}
    chanForSortedArray := make(chan [10]int)

    go BubbleSort(chanForSortedArray, testArray)

    sortedArray := <-chanForSortedArray
    fmt.Println("Test array: ", testArray)
    fmt.Println("Sorted array: ", sortedArray)
}