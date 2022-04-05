// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"todo"

// 	"github.com/gofiber/fiber/v2"
// )

// var todoFileName = ".todo.json"

// func main() {

// 	flag.Usage = func() {
// 		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed by Spencer Birch\n", os.Args[0])
// 		fmt.Fprintf(flag.CommandLine.Output(), "Usage Information: \n")
// 		flag.PrintDefaults()
// 	}

// 	task := flag.String("task", "", "Task to be included in the todo list")
// 	list := flag.Bool("list", false, "List all tasks")
// 	complete := flag.Int("complete", 0, "Item to be completed")

// 	flag.Parse()

// 	export TODO_FILENAME=newfilename.json
// 	if os.Getenv("TODO_FILENAME") != "" {
// 		todoFileName = os.Getenv("TODO_FILENAME")
// 	}

// 	l := &todo.List{}

// 	if err := l.Get(todoFileName); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(1)
// 	}

// 	switch {
// 	case *list:
// 		// CList current todo items
// 		fmt.Print(l)

// 	case *complete > 0:
// 		// Complete the given item
// 		if err := l.Complete(*complete); err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			os.Exit(1)
// 		}

// 		// Save new list
// 		if err := l.Save(todoFileName); err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			os.Exit(1)
// 		}

// 	case *task != "":
// 		// Add task
// 		l.Add(*task)

// 		// Save new list
// 		if err := l.Save(todoFileName); err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			os.Exit(1)
// 		}

// 	default:
// 		// invalid flag
// 		fmt.Fprintln(os.Stderr, "Invalid option")
// 		os.Exit(1)
// 	}
// }
