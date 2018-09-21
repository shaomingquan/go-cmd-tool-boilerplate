package main

import (
	"flag"
	"fmt"
)

var run commandHandler

func init() {
	run = func(command string, params []string) {
		a := false
		b := 0
		flag.BoolVar(&a, "a", false, "noop")
		flag.IntVar(&b, "b", 0, "noop")
		flag.CommandLine.Parse(params)

		fmt.Printf("run~ %t %d", a, b)
	}

	// register command handler
	handlers["run"] = run
}
