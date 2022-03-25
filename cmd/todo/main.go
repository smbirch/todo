package main

import (
	"fmt"
	"os"
	"todo"
)

//Harcoding name of file for now
const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
