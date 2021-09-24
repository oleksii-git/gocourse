package main

import (
    "fmt"
    "strings"
    "log"
    "os"
    "bufio"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string:")
	input, err := reader.ReadString('\n')
	if err != nil {
	    log.Fatal(err)
    }
	input = strings.Replace(input, "\n", "", -1)

    search := ((strings.Contains(input, "a") || strings.Contains(input, "A")) &&
               (strings.HasPrefix(input, "i") || strings.HasPrefix(input, "I")) &&
               (strings.HasSuffix(input, "n") || strings.HasSuffix(input, "N")))

   	if search{
   	    fmt.Println("Found!")
   	} else{
   	    fmt.Println("Not Found!")
   	}

}