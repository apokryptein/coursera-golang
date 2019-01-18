// read.go
// Final project for class
// Author: Tomlinson

package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
)

func main() {
	type name struct {
		fname string
		lname string
	}

	nameSlice := make([]name, 0)
	var p1 name

	var fileName string
	fmt.Printf("Filename: ")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	inScan := bufio.NewScanner(file)

	for inScan.Scan() {
		s := strings.Split(inScan.Text(), " ")
		p1.fname = s[0]
		p1.lname = s[1]
		nameSlice = append(nameSlice, p1)
	}


	fmt.Println("\nNames printed LAST, FIRST\n")

	for _, person := range nameSlice {
		fmt.Println(person.lname + ", " + person.fname)
	}

	fmt.Println("\nComplete.  Exiting...\n")

}