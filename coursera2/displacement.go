package main

import (
    "fmt"
    "strconv"
    "time"
)

func GenDisplaceFn(acceleration float64, initialVelocity float64, initialDisplacement float64) func(inputTime float64) float64{
    return func(inputTime float64) float64{
        time.Sleep(time.Duration(inputTime) * time.Second)
        displacement := (acceleration * inputTime * inputTime) / 2 + initialVelocity * inputTime + initialDisplacement
        return displacement
    }
}

func main() {
    testArray := []float64{}
    var input string
    var inputTime float64
    var message string

    for i := 0; i < 4; i++ {
        switch i {
            case 0:
                message = "Enter acceleration: "
            case 1:
                message = "Enter initial velocity: "
            case 2:
                message = "Enter initial displacement: "
            case 3:
                message = "Enter time: "
        }
        fmt.Println(message)
        _, err := fmt.Scan(&input)
        if err != nil {
            fmt.Println("[ERROR] Incorrect input! Can not scan the input...")
            i--
            //i-- for endless loop
        } else{
            if input == "X" {
                fmt.Println("Exit")
                break
            } else {
                item, err := strconv.ParseFloat(input, 32)
                if err != nil {
                    fmt.Println("[ERROR] Incorrect input! Input is not a number...")
                    i--
                    //i-- for endless loop
                } else {
                    if i != 3 {
                        testArray = append(testArray, item)
                    } else {
                        inputTime = item
                    }

                }
            }
        }
    }


    fmt.Println("Input [acceleration, initial velocity, initial displacement]: ", testArray)
    fmt.Println("Input time: ", inputTime)

    Fn := GenDisplaceFn(testArray[0], testArray[1], testArray[2])
    fmt.Println("displacement =", Fn(inputTime))

}