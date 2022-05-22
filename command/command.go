package main

import (
	"fmt"
	"github.com/MrPaschenko/event-loop/engine"
)

type printCommand struct {
	output string
}

func (print *printCommand) Execute(loop engine.Handler) {
	fmt.Println(print.output)
}

type Delete struct {
	Args []string
}

func (d *Delete) Execute(handler engine.Handler) {
	if len(d.Args) != 2 {
		handler.Post(&printCommand{output: "less than two arguments"})
		return
	}
	str := d.Args[0]
	symb := d.Args[1]

	res := ""
	for _, c := range str {
		char := string(c)
		if char != symb {
			res += char
		}
	}
	handler.Post(&printCommand{output: res})
}
