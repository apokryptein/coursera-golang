// findian.go
// Go course Module 2
// Author: Tomlinson

package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	fmt.Printf("Please enter a string: ")

	inputReader := bufio.NewReader(os.Stdin)
	inputString, _ := inputReader.ReadString('\n')
	
	inputLower := strings.TrimSpace(strings.ToLower(inputString))
	
	if strings.HasPrefix(inputLower, "i") && strings.HasSuffix(inputLower, "n") && strings.Contains(inputLower, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
