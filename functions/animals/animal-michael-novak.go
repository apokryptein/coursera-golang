package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

var animals = map[string]Animal{
	"cow":   Animal{"grass", "walk", "moo"},
	"bird":  Animal{"worms", "fly", "peep"},
	"snake": Animal{"mice", "slither", "hsss"},
}

// Map of functions with receivers
// https://stackoverflow.com/questions/45472816/golang-a-map-of-functions-with-a-receiver
var commands = map[string]func(Animal){
	"eat":   (Animal).Eat,
	"move":  (Animal).Move,
	"speak": (Animal).Speak,
}

func input(label string) (string, error) {
	fmt.Printf("%s", label)
	stdin := bufio.NewReader(os.Stdin)
	s, err := stdin.ReadString('\n')
	if err != nil {
		return "", err
	}
	s = strings.TrimRight(s, "\r\n")
	return s, nil
}

func main() {
	for {
		in, err := input("> ")
		if err != nil {
			log.Println(err)
			continue
		}
		if in == "X" {
			fmt.Println("Bye")
			break
		}
		cmd := strings.Fields(in)
		if len(cmd) < 2 {
			fmt.Println("Request must contains animal (any from cow, bird or snake) and information requested (any from eat, move or speak)")
			continue
		}
		animal, ok := animals[cmd[0]]
		if !ok {
			fmt.Println("Unknown animal -", cmd[0])
			continue
		}
		information, ok := commands[cmd[1]]
		if !ok {
			fmt.Println("Unknown information requested -", cmd[1])
			continue
		}
		information(animal)
	}
}
