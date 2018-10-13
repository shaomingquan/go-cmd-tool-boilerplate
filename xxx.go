package main

import (
	"fmt"
)

var xxx commandHandler

func init() {
	xxx := func(command string, params []string) {
		fmt.Println(`
			just print it
		`)
	}

	handlers["xxx"] = xxx
}
