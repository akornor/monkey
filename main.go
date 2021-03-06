package main

import (
	"fmt"
	"github.com/akornor/monkey/repl"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func init() {
	// Silence logs when testing
	log.SetOutput(ioutil.Discard)
}

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands.\n")
	repl.Start(os.Stdin, os.Stdout)
}
