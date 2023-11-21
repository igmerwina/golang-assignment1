package main

import (
	"bufio"
	"fmt"
	"os"
	"pgd-sesi3/function"
	"strings"
)

func main() {
	var seed = function.SeedData()

	fmt.Print("Enter something: ")

	reader := bufio.NewReader(os.Stdin)

	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	userInput = strings.TrimSpace(userInput)

	for _, v := range seed {
		if strings.ToLower(userInput) == "" || strings.ToLower(userInputuserInput != v.Nama {
			fmt.Println("Data Tidak Ditemukan! \n")
			return
		}

		if userInput == v.Nama {
			fmt.Println(v)
			return
		}
	}
}
