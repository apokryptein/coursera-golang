// animals.go
// Author: Tomlinson

package main

import (
	"fmt"
)


type Animal struct {
	food string
	locomotion string
	noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}



func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	var animal, action string

	fmt.Println("\nPress ^C to exit")
	fmt.Println("On one line, enter the animal name, a space, then an action")
	fmt.Println("Animals: cow, bird, or snake")
	fmt.Println("Actions: eat, move, or speak")
	fmt.Println("For example: bird speak\n")

	for {
		fmt.Printf("> ")
		fmt.Scan(&animal, &action)

		if animal != "cow" && animal != "bird" && animal != "snake" {
			fmt.Println("Animal not in our database.  Please try again...")
		}

		if action != "eat" && action != "move" && action != "speak" {
			fmt.Println("Action not in our database.  Please try again...")
		}

		if animal == "cow" {
			if action == "eat" {
				cow.Eat()
			}

			if action == "move" {
				cow.Move()
			}

			if action == "speak" {
				cow.Speak()
			}
		}

		if animal == "bird" {
			if action == "eat" {
				bird.Eat()
			}

			if action == "move" {
				bird.Move()
			}

			if action == "speak" {
				bird.Speak()
			}
		}

		if animal == "snake" {
			if action == "eat" {
				snake.Eat()
			}

			if action == "move" {
				snake.Move()
			}

			if action == "speak" {
				snake.Speak()
			}
		}				

	} // end for loop
} // end main() func





