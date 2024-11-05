# `ls` command in Go

This is a simple implementation of the `ls` (list) command in Go. The program lists the contents of one or more directories, displaying file and directory names along with optional details.

## Usage

To use the program, run the following command in your terminal:

```
./ls <directory1> <directory2> ...
```

Replace `<directory1>`, `<directory2>`, etc., with the paths of the directories you want to list.

> [!NOTE]
> For `./ls` to work, you first need to compile the `ls.go` file using `go build ls.go`.

## Features

- Lists files and directories in specified directories.
- Supports multiple directory inputs.

## Requirements

- Go 1.16 or later.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/m3hu1/learning-go.git
2. Navigate to the project directory:
   ```
   cd learning-go/ls
   ```
3. Build the program:
   ```sh
   go build ls.go
   ```
