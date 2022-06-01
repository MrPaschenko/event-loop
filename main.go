package main

import (
	"bufio"
	"github.com/MrPaschenko/event-loop/engine"
	"os"
	"strings"
)

func parse(commandLine string) engine.Command {
	parts := strings.Fields(commandLine)

	if len(parts) == 0 {
		return &PrintCommand{output: "SYNTAX ERROR: no arguments"}
	}

	switch parts[0] {
	case "print":
		if len(parts) == 2 {
			return &PrintCommand{output: parts[1]}
		} else if len(parts) == 1 {
			return &PrintCommand{output: "SYNTAX ERROR: not enough arguments"}
		} else {
			return &PrintCommand{output: "SYNTAX ERROR: too many arguments"}
		}
	case "delete":
		if len(parts) < 3 {
			return &PrintCommand{output: "SYNTAX ERROR: not enough arguments"}
		} else if len(parts) > 3 {
			return &PrintCommand{output: "SYNTAX ERROR: too many arguments"}
		} else {
			return &Delete{parts}
		}
	default:
		return &PrintCommand{output: "SYNTAX ERROR: command not found"}
	}
}

func main() {
	Loop := new(engine.Loop)
	Loop.Start()

	if input, err := os.Open("inputFile.txt"); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine)
			Loop.Post(cmd)
		}
	}
	Loop.AwaitFinish()
}
