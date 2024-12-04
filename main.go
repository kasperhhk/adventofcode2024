package main

import (
	"fmt"
	"khhk-adventofcode2024/reader"
)

func main() {
	fmt.Println("Line by line:")

	lines := reader.Lines("testfile.txt")
	for _, line := range lines {
		fmt.Println(line)
	}

	fmt.Println("All text:")

	fmt.Print(reader.AllText("testfile.txt"))
}
