package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

type fullName struct{
    fname string
    lname string
}

func main() {
    var input string

    fmt.Println("Enter a file name:")
    fmt.Scan(&input)
    dat, err := ioutil.ReadFile(input)
    if err != nil {
        log.Fatal(err)
    }
    file := string(dat)
    items := strings.Split(file, "\n")
    structuredItems := []fullName{}
    item := []string{}
    for i := 0; i < len(items); i++{
        item = strings.Split(items[i], " ")
        structuredItems = append(structuredItems, fullName{fname: item[0], lname: item[1]})
        fmt.Printf("Line: %d First name: %s Last name: %s\n", i+1, item[0], item[1])
    }
    //fmt.Printf("%+v", structuredItems)
}