package main

import (
	"fmt"
	"github.com/MrPaschenko/event-loop/engine"
)

type PrintCommand struct {
	output string
}

func (print *PrintCommand) Execute(loop engine.Handler) {
	fmt.Println(print.output)
}

type Delete struct {
	Args []string
}

func (d *Delete) Execute(handler engine.Handler) {
	if len(d.Args) != 3 {
		handler.Post(&PrintCommand{output: "invalid number of arguments"})
		return
	}
	str := d.Args[1]
	symb := d.Args[2]

	res := ""
	for _, c := range str {
		char := string(c)
		if char != symb {
			res += char
		}
	}
	handler.Post(&PrintCommand{output: res})
}
