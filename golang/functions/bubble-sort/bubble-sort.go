// bubble-sort.go
// Author: Tomlinson

package main

import (
	"fmt"
)

func main() {
	var intNum int
	fmt.Printf("\nNumber of integers to sort (max 10)? ")
	fmt.Scan(&intNum)

	intSlice := make([]int, intNum)

	fmt.Println("Enter integers on one line separated by spaces, then press ENTER")
	fmt.Printf("Integers: ")
	for i := range intSlice {
		//fmt.Printf("Integer %d: ", i + 1)
		fmt.Scan(&intSlice[i])
	}

	BubbleSort(intSlice)
	fmt.Println(intSlice)
} //end main functionß


func BubbleSort(numSlice []int) {

	for i := 0; i < len(numSlice); i++ {
		for j := 0; j < len(numSlice) - i - 1; j++ {
			if numSlice[j] > numSlice[j + 1] {
				Swap(numSlice, j)
			} //endif
		} //end inner for
	} //end outer for
} //end BubbleSort functionß


func Swap(swapSlice []int, i int) {
	swapSlice[i], swapSlice[i + 1] = swapSlice[i + 1], swapSlice[i]
} //end Swap functionß
