package main

import (
	"bufio"
	"os"
	"strings"
	"github.com/MrPaschenko/event-loop/engine"
)

func parse(commandLine string) engine.Command{
	parts := strings.Fields(commandLine)

	if len(parts) == 0{
		return &printCommand{output: "SYNTAX ERROR: no arguments"}
	}

	switch parts[0] {
	case "print":
		if len(parts) == 2 {
			return &printCommand{output: parts[1]}
		} else if len(parts) == 1{
			return &printCommand{output: "SYNTAX ERROR: not enough arguments"}
		} else {
			return &printCommand{output: "SYNTAX ERROR: too many arguments"}
		}
	case "cat":
		if len(parts) < 3 {
			return &printCommand{output: "SYNTAX ERROR: not enough arguments"}
		} else if len(parts) > 3 {
			return &printCommand{output: "SYNTAX ERROR: too many arguments"}
		} else {
			return &catCommand{arg1: parts[1], arg2: parts[2]}
		}
	default:
		return &printCommand{output: "SYNTAX ERROR: command not found"}
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