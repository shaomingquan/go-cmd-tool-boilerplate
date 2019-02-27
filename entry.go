package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type commandHandler func(command string, params []string)

var handlers = map[string]commandHandler{}

func main() {
	args := os.Args

	if args[1] == "--help" {
		flag.Usage = func() {
			// set your usage
			fmt.Println(`
				this is usage
			`)
		}
		flag.Parse()
		return
	}

	if len(args) < 2 {
		log.Fatal("must with command")
	}

	command := os.Args[1]

	params := []string{}
	if len(args) > 2 {
		params = os.Args[2:]
	}

	if _, ok := handlers[command]; ok {
		defer helloAndBye(command)()
		handlers[command](command, params)
	} else {
		log.Fatal("cmd " + command + " does't exsit")
	}
}
