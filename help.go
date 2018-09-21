package main

import (
	"fmt"
)

var help commandHandler

func init() {
	help := func(command string, params []string) {
		fmt.Println(`
			just print it
		`)
	}

	handlers["help"] = help
}
