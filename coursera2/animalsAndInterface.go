package main

import (
	"fmt"
)

type animal interface {
	Eat()
	Move()
	Speak()
	GetName() string
}
type Cow struct {
	Name       string
	food       string "grass"
	locomotion string "walk"
	noise      string "moo"
}

func (animal Cow) Move() {
	fmt.Println(animal.locomotion)
}
func (animal Cow) Eat() {
	fmt.Println(animal.food)
}
func (animal Cow) Speak() {
	fmt.Println(animal.noise)
}
func (animal Cow) GetName() string {
	return animal.Name
}

type Bird struct {
	Name       string
	food       string "worms"
	locomotion string "fly"
	noise      string "peep"
}

func (animal Bird) Move() {
	fmt.Println(animal.locomotion)
}
func (animal Bird) Eat() {
	fmt.Println(animal.food)
}
func (animal Bird) Speak() {
	fmt.Println(animal.noise)
}
func (animal Bird) GetName() string {
	return animal.Name
}

type Snake struct {
	Name       string
	food       string "mice"
	locomotion string "slither"
	noise      string "hsss"
}

func (animal Snake) Move() {
	fmt.Println(animal.locomotion)
}
func (animal Snake) Eat() {
	fmt.Println(animal.food)
}
func (animal Snake) Speak() {
	fmt.Println(animal.noise)
}
func (animal Snake) GetName() string {
	return animal.Name
}

func query(animal_name string, request string, animals []animal) {
	notFind := true
	for i := 0; i < len(animals); i++ {
		if animal_name == animals[i].GetName() {
			notFind = false
			switch request {
			case "move":
				animals[i].Move()
			case "eat":
				animals[i].Eat()
			case "speak":
				animals[i].Speak()
			default:
				fmt.Println("[ERROR] Incorrect <request> parameter!")
			}
		}
	}
	if notFind {
		fmt.Println("[WARNING] Not found animal with name: ", animal_name)
	}
}

func newAnimal(animal_name string, animalType string, animals *[]animal) {
	created := false
	switch animalType {
	case "cow":
		*animals = append(*animals, Cow{
			Name:       animal_name,
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		})
		created = true
	case "bird":
		*animals = append(*animals, Cow{
			Name:       animal_name,
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		})
		created = true
	case "snake":
		*animals = append(*animals, Cow{
			Name:       animal_name,
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		})
		created = true
	default:
		fmt.Println("[WARNING] Can not create animal with type: ", animalType)
	}
	if created {
		fmt.Println("[INFO] Created animal with name: ", animal_name)
	}
}

func main() {
	var animals []animal
	multiline := `Enter string in this format: <query> <animal_name> <request>
<task> can be: 'newanimal' or 'query'
<animal_name> can be: any string
<request> can be: 'move', 'eat' or 'speak' if <query> is 'query', or 'cow', 'bird', or 'snake'
>`
	var task, animal_name, request string
	for true {
		fmt.Println("<<<<<<<<<ITERATION START>>>>>>>>>")
		fmt.Print(multiline)
		_, err1 := fmt.Scan(&task)
		_, err2 := fmt.Scan(&animal_name)
		_, err3 := fmt.Scan(&request)
		if (err1 != nil) || (err2 != nil) || (err3 != nil) {
			fmt.Println("[ERROR] Incorrect input! Can not scan the input...")
		} else {
			switch task {
			case "newanimal":
				fmt.Println("[INFO] Creating new animal with the name: ", animal_name)
				newAnimal(animal_name, request, &animals)
			case "query":
				fmt.Println("[INFO] Getting requested animal`s info!")
				query(animal_name, request, animals)
			default:
				fmt.Println("[ERROR] Incorrect <task> parameter!")
			}
		}
		fmt.Println("Full array of animals: ", animals)
		fmt.Println("<<<<<<<<<ITERATION END>>>>>>>>>\n")
	}
}
