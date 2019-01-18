// makejson.go
// Author: Tomlinson

package main

import (
	"fmt"
	"encoding/json"	
	"bufio"
	"os"
	"strings"
)

func main() {
	userMap := make(map[string]string)

	inputReader := bufio.NewReader(os.Stdin)


	fmt.Println("Please enter the following information")
	fmt.Printf("Name: ")
	userName, _ := inputReader.ReadString('\n')
	fmt.Printf("Address: ")
	userAddress, _ := inputReader.ReadString('\n')

	userMap["name"] = strings.TrimSuffix(userName, "\n")
	userMap["address"] = strings.TrimSuffix(userAddress, "\n")

	userProf, _ := json.Marshal(userMap)

	fmt.Println(string(userProf))
}