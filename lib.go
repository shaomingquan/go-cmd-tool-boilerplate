package main

import "fmt"

func helloAndBye(cmd string) func() {
	fmt.Println(cmd + " start")
	return func() {
		fmt.Println(cmd + " done")
	}
}
