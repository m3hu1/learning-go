# `cat` command in Go

This is a simple implementation of the `cat` command in Go. The program reads one or more text files and outputs their contents to the standard output, line by line, while numbering each line.

## Usage

To use the program, run the following command in your terminal:

```
./cat <file1> <file2> <file3> ...
```

Replace `<file1>`, `<file2>`, etc., with the names of the files you want to read.
> [!NOTE]
> In order for `./cat` to work, you first have to compile the `cat.go` file using `go build cat.go`.

## Features

- Reads multiple files.
- Outputs each line with a line number.
- Skips blank lines.

## Requirements

- Go 1.16 or later.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/m3hu1/learning-go.git
   ```
2. Navigate to the project directory:
   ```
   cd learning-go/cat
   ```
3. Build the program:
   ```sh
   go build cat.go
   ```
