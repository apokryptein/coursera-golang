// slice.go
// Takes integer input from user, stores in a slice and outputs
// contents of the slice in sorted order
// Author: Tomlinson

package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"sort"
)

func main() {
	intStore := make([]int, 3)
	inputReader := bufio.NewReader(os.Stdin)
	var i int = 0

	fmt.Println("Enter 'X' when you are finished.")


	for {
		fmt.Printf("Enter an integer: ")
		inputVal, _ := inputReader.ReadString('\n')
		inputVal = strings.TrimSuffix(inputVal, "\n")

		if inputVal == "X" {
			break
		} else if inputVal == "" {
			fmt.Println("Please enter an integer.  Try again...")
			continue
		}

		valFinal, _ := strconv.Atoi(inputVal)

		if i < 3 {
			intStore[i] = valFinal
		} else {
			intStore = append(intStore, valFinal)
		}

		i++
	}
	sort.Ints(intStore)
	fmt.Println(intStore)
}