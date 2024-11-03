package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./wc <file1> <file2> ...")
		return
	}

	totalLines := 0
	totalWords := 0
	totalChars := 0

	fmt.Printf("%8s %8s %8s \n", "lines", "words", "chars")

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file %s ❌ %v\n", filename, err)
			continue
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		lineCount := 0
		wordCount := 0
		charCount := 0

		for {
			line, err := reader.ReadString('\n')
			if err != nil && err != io.EOF {
				fmt.Printf("Error reading file %s: %v\n", filename, err)
				break
			}

			if err == io.EOF && line == "" {
				break
			}

			lineCount++
			wordCount += len(strings.Fields(line))
			charCount += len(line)

			if err == io.EOF {
				break
			}
		}

		fmt.Printf("%8d %8d %8d --> %s\n", lineCount, wordCount, charCount, filename)

		totalLines += lineCount
		totalWords += wordCount
		totalChars += charCount
	}

	if len(os.Args) > 2 {
		fmt.Printf("%8d %8d %8d --> total\n", totalLines, totalWords, totalChars)
	}
}
