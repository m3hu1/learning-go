# `which` command in Go

This is a simple implementation of the `which` command in Go. The program searches for the executable file associated with the given command in the directories listed in the system's PATH environment variable.

## Usage

To use the program, run the following command in your terminal:

```
./which <command>
```

Replace `<command>` with the name of the command you want to locate.

> [!NOTE]
> In order for `./which` to work, you first have to compile the `which.go` file using `go build which.go`.

## Features

- Searches for the specified command in the directories listed in the PATH environment variable.
- Outputs the full path of the executable if found.
- Handles multiple commands in a single invocation.

## Requirements

- Go 1.16 or later.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/m3hu1/learning-go.git
   ```
2. Navigate to the project directory:
   ```
   cd learning-go/which
   ```
3. Build the program:
   ```sh
   go build which.go
   ```
