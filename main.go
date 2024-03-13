package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s!\n", user.Username)
	fmt.Print("Monkey Interpreter v0.1\n")

	repl.Start(os.Stdin, os.Stdout)
}
