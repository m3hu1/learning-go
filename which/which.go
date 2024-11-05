package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	path, err := exec.LookPath(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Println(path)
}
