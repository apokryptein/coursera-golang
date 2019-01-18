// Golang course, Module 2
// trunc.go
// Author: Tomlinson

package main

import "fmt"

func main() {
	var truncNum float64
	fmt.Printf("Floating point number: ")
	fmt.Scan(&truncNum)
	fmt.Println(int64(truncNum))
}
