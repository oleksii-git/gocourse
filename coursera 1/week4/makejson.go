package main

import (
    "fmt"
    "log"
    "encoding/json"
)

func main() {
    m := make(map[string]string)
    var input string

    fmt.Println("Enter a name:")
    fmt.Scan(&input)
    m["name"] = input

    fmt.Println("Enter an address:")
    fmt.Scan(&input)
    m["address"] = input

    jsonString, error1 := json.Marshal(m)
    if error1 != nil{
        log.Fatal("Invalid input data! Can not convert map to json")
    }
    fmt.Println(string(jsonString))
}