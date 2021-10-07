package main

import (
	"fmt"
	"sync"
)

//in main you can see the example of race condition
//set the values for TestMap in goroutine(line 19) and in main goroutine(line 24)
//this case the problem of working with one data structure in the concurrent goroutines which is called race condition

func main() {
	var wg sync.WaitGroup
	//var m sync.Mutex
	var TestMap = make(map[string]string)
	wg.Add(1)
	go func() {
		//m.Lock()
		TestMap["1"] = "a" // conflicting access with the main routine.
		//m.Unlock()
		wg.Done()
	}()
	TestMap["2"] = "b"
	wg.Wait()
	fmt.Println("testMap: ", TestMap)
}

// you can test it using this commands, in result of this commands you will receive WARNING about race condition:
//1. go build -race goroutineRaceCond.go
//2. ./goroutineRaceCond

//if you uncomment all commented lines, the problem of race condition will be solved
