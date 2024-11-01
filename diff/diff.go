package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: ./diff <file1> <file2>")
		return
	}

	file1, _ := os.Open(os.Args[1])
	defer file1.Close()

	file2, _ := os.Open(os.Args[2])
	defer file2.Close()

	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	lineCounter := 1

	for scanner1.Scan() && scanner2.Scan() { // loop until both files have lines to read
		line1 := scanner1.Text()
		line2 := scanner2.Text()
		if line1 != line2 { // comparing the text of lines in both the files
			fmt.Printf("Difference at line %d:\n", lineCounter)
			fmt.Printf("- %s\n", line1)
			fmt.Printf("+ %s\n", line2)
		}
		lineCounter++
	}

	// if len(file1) > len(file2), then file1 will have extra lines
	for scanner1.Scan() {
		fmt.Printf("Difference at line %d:\n", lineCounter)
		fmt.Printf("- %s\n", scanner1.Text())
		lineCounter++
	}

	// if len(file2) > len(file1), then file2 will have extra lines
	for scanner2.Scan() {
		fmt.Printf("Difference at line %d:\n", lineCounter)
		fmt.Printf("+ %s\n", scanner2.Text())
		lineCounter++
	}
}
