package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("To end the output Write . and then click enter")
	fmt.Print("Enter text to write to file a file called \"OutputFromTerminal.txt\": \n")

	input, _ := reader.ReadString('.')

	path := filepath.Join("./", "OutputFromTerminal.txt")
	file, _ := os.Create(path)

	writer := bufio.NewWriter(file)
	writer.WriteString(input)

	defer file.Close()
	defer writer.Flush()

}
