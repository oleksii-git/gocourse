package main

import "fmt"

func main() {
	var input float32
	fmt.Println("Enter a floating point number:")
	fmt.Scan(&input)
	fmt.Println("Your truncated number is", int(input))
}