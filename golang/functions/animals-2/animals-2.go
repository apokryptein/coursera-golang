// animals-2.go
// Author: Tomlinson

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// interface definition
type Animal interface {
	Eat()
	Move()
	Speak()
}

// cow struct
type Cow struct {
	food string
	locomotion string
	sound string
}

// cow functions
func (c Cow) Eat() {fmt.Println(c.food)}
func (c Cow) Move() {fmt.Println(c.locomotion)}
func (c Cow) Speak() {fmt.Println(c.sound)}


// bird struct
type Bird struct {
	food string
	locomotion string
	sound string
}

// bird functions
func (b Bird) Eat() {fmt.Println(b.food)}
func (b Bird) Move() {fmt.Println(b.locomotion)}
func (b Bird) Speak() {fmt.Println(b.sound)}


// snake struct
type Snake struct {
	food string
	locomotion string
	sound string
}

// snake functions
func (s Snake) Eat() {fmt.Println(s.food)}
func (s Snake) Move() {fmt.Println(s.locomotion)}
func (s Snake) Speak() {fmt.Println(s.sound)}


var animals = map[string]Animal{}

var commands = map[string]func(Animal){
	"eat":   (Animal).Eat,
	"move":  (Animal).Move,
	"speak": (Animal).Speak,
}

// main method
func main() {
	fmt.Println("\nEnter a new animal with the following syntax: newanimal <animal-name> <animal-type>")
	fmt.Println("Animal type choices: cow, bird, snake")
	fmt.Println("To query your animals: query <animal-name> <requested-information>")
	fmt.Println("Requested information types: eat, move, speak\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		
		if input == "exit" {
			fmt.Println("Exiting...Goodbye")
			break
		} // end if

		cmd := strings.Fields(input)
		
		switch cmd[0] {
		case "newanimal":
			switch cmd[2] {
			case "cow":
				animals[cmd[1]] = Cow{"grass", "walk", "moo"}
			case "bird":
				animals[cmd[1]] = Bird{"worms", "fly", "peep"}
			case "snake":
				animals[cmd[1]] = Snake{"mice", "slither", "hsss"}
			default:
				fmt.Println("Error processing input")
				continue
			} // end nested switch
		case "query":
			commands[cmd[2]](animals[cmd[1]])
		default:
			fmt.Println("Command not found.")
			continue
		} // end switch

	} // end for loop
} // end main func