package reader

import (
	"bufio"
	"os"
)

func Lines(filepath string) []string {
	var lines []string

	file, _ := os.Open(filepath)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func AllText(filepath string) string {
	b, _ := os.ReadFile(filepath)
	return string(b)
}
