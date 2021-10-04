package main

import (
    "fmt"
    //"strconv"
)

type Animal struct {
    food, locomotion, noise string
}

func (animal Animal) Move() string {
	return animal.locomotion
}

func (animal Animal) Eat() string {
	return animal.food
}

func (animal Animal) Speak() string {
	return animal.noise
}

func main() {

    var animals = []Animal{
        Animal{
             food: "grass",
             locomotion: "walk",
             noise: "moo",
        },
        Animal{
            food: "worms",
            locomotion: "fly",
            noise: "peep",
        },
        Animal{
            food: "mice",
            locomotion: "slither",
            noise: "hsss",
        },
    }

    var input1, input2 string
    for true {
        fmt.Println("<<<<<<<<<ITERATION START>>>>>>>>>>>")
        fmt.Print("Enter string in this format: <animal> <request>\n<animal> can be: 'cow', 'bird' or 'snake'\n<request> can be: 'move', 'eat' or 'speak'\n>")
        _, err1 := fmt.Scan(&input1)
        _, err2 := fmt.Scan(&input2)
        if (err1 != nil) || (err2 != nil) {
            fmt.Println("[ERROR] Incorrect input! Can not scan the input...")
        } else{
            if (input1 == "X") || (input2 == "X") {
                fmt.Println("Exit")
                break
            } else {
                validation := true
                var selectedAnimal *Animal
                switch input1 {
                    case "cow":
                        selectedAnimal = &animals[0]
                    case "bird":
                        selectedAnimal = &animals[1]
                    case "snake":
                        selectedAnimal = &animals[2]
                    default:
                        validation = false
                        fmt.Println("[ERROR] Incorrect <animal> parameter!")
                }
                if validation{
                    switch input2 {
                        case "move":
                            fmt.Println("Requested data:", (*selectedAnimal).Move())
                        case "eat":
                            fmt.Println("Requested data:", (*selectedAnimal).Eat())
                        case "speak":
                            fmt.Println("Requested data:", (*selectedAnimal).Speak())
                        default:
                            fmt.Println("[ERROR] Incorrect <request> parameter!")
                    }
                }
            }
        }
        fmt.Println("<<<<<<<<<ITERATION END>>>>>>>>>>>")
    }


}