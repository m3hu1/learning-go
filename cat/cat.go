package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./cat <file1> <file2> <file3> ...")
		return
	}

	line := 1 // to count the lines

	for _, filename := range os.Args[1:] { // idx, arg; starting from idx 1 cuz idx 0 is the program name itself
		file, _ := os.Open(filename)      // *os.File, err
		defer file.Close()                // close the file when the function ends; defer schedules a func call to be exec after the surrounding func completes
		scanner := bufio.NewScanner(file) // Scanner reads line by line by default

		for scanner.Scan() { // loop as long as there are lines to read
			if scanner.Text() == "" { // to exclude the blank lines
				continue
			}
			fmt.Printf("%d: %s\n", line, scanner.Text()) // scanner.Text() returns the curr line's data (text)
			line++
		}
	}
}
