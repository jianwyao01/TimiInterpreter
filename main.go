package main

import (
	"TimiInterpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is a playful programming language!\n", current.Username)
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
