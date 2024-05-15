package main

import (
	"fmt"
	"monkey-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}	
	fmt.Printf("Hello " + user.Username + "! This is the Monkey programming language! ")
	fmt.Printf("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
	// repl.Start(os.Stdin, os.Stdout)

	// Start the REPL
}