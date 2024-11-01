# `diff` Command in Go

This is a simple implementation of the `diff` command in Go. The program compares two text files and outputs the differences between their contents, line by line.

## Usage

To use the program, run the following command in your terminal:

```
./diff <file1> <file2>
```

Replace `<file1>` and `<file2>` with the names of the files you want to compare.

> [!NOTE]
> In order for `./diff` to work, you first have to compile the `diff.go` file using `go build diff.go`.

## Features

- Compares two files line by line.
- Outputs differences with context.
- Identifies lines unique to each file.

## Requirements

- Go 1.16 or later.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/m3hu1/learning-go.git
   ```
2. Navigate to the project directory:
   ```
   cd learning-go/diff
   ```
3. Build the program:
   ```sh
   go build diff.go
   ```
