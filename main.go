package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	e != nil
	{
		panic(e)
	}
}

func write(a, b string) {
	path := filepath.Join(a, b)
	file, err := os.Create(path)
	check(err)

	file.Sync()

	writer := bufio.NewWriter(file)
	writer.WriteString("Created a file\nAnd wrote in it!")

	writer.Flush()
}

func main() {
	fmt.Println("Leanring go.")
	fmt.Println("Adding a function to create a file")
	write("./", "file.txt")
}
