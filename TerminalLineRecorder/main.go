package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/term"
)

func main() {
	// Create a reader
	reader := bufio.NewReader(os.Stdin)

	//Create a file path:
	err := os.Mkdir("./Log", os.ModePerm)

	check(err)

	path := filepath.Join("./Log", "TerminalLog.txt")
	file, _ := os.Create(path) // Creating the file and assinging it the variable file

	// Create a writer
	writer := bufio.NewWriter(file)

	// The rest.
	fmt.Println("Enter in commands you were to run in the terminal, this go file acts as a middle man recorder: \n")
	input, _ := reader.ReadString('.')

	var inputs []string

	cleanedInput := strings.Split(inputs, "\n")

	writer.WriteString(cleanedInput)

	term.NewTerminal(_, input) * Terminal

	defer writer.Flush()
	defer file.Close()

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
