# `wc` command in Go

This is a simple implementation of the `wc` (word count) command in Go. The program reads one or more text files and outputs the number of lines, words, and characters in each file.

## Usage

To use the program, run the following command in your terminal:

```
./wc <file1> <file2> <file3> ...
```

Replace `<file1>`, `<file2>`, etc., with the names of the files you want to analyze.

> [!NOTE]
> In order for `./wc` to work, you first have to compile the `wc.go` file using `go build wc.go`.

## Features

- Counts the number of lines, words, and characters in each file.
- Supports multiple file inputs.
- Outputs a summary for each file.

## Requirements

- Go 1.16 or later.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/m3hu1/learning-go.git
   ```
2. Navigate to the project directory:
   ```
   cd learning-go/wc
   ```
3. Build the program:
   ```sh
   go build wc.go
   ```
