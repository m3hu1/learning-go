package main

import (
	"fmt"
	"os"
)

func main() {
	dirs := os.Args[1:]

	if len(dirs) == 0 {
		dirs = append(dirs, "./")
	}

	for i, dir := range dirs {
		if len(dirs) > 1 {
			fmt.Printf("%s:\n", dir)
		}

		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Printf("Error reading directory: %v\n", err)
			continue
		}

		for _, f := range files {
			fmt.Printf("%s ", f.Name())
		}

		fmt.Println()

		if i < len(dirs)-1 {
			fmt.Println()
		}
	}
}
