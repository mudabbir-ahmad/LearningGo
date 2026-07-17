package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 1. Ask for Username
	fmt.Print("Username: ")
	usernameInput, err := reader.ReadString('\n')

	check(err)
	username := strings.TrimSpace(usernameInput)

	// 2. Ask for Password (with asterisk masking)
	fmt.Print("Password: ")
	passwordIn, _ := reader.ReadString('\n')
	password := string(passwordIn)

	// 3. Output results
	fmt.Println("\n--- Result ---")
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)

	fmt.Print("The delim should cause you to end the line whenever you press back slash, then hit enter: \n")
	test, _ := reader.ReadString('\\')
	fmt.Println("You entered:", test)
}
