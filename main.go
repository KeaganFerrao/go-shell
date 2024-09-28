package main

import (
	"bufio"
	"fmt"
	"goshell/cmd"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("shell>")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			continue
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Bye!")
			break
		}

		if input == "" {
			continue
		}

		data := strings.Split(input, " ")
		command := data[0]
		commandArgs := data[1:]

		c := cmd.Command{
			Arguments: commandArgs,
		}

		switch command {
		case "pwd":
			c.Pwd()
		case "ls":
			c.Ls()
		case "touch":
			c.Touch()
		case "mkdir":
			c.Mkdir()
		default:
			fmt.Printf("Command %v not supported\n", command)
		}
	}
}
