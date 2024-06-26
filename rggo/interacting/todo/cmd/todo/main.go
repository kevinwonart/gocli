package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"gocli/rggo/interacting/todo"
)

var todoFileName = ".todo.json"

func main() {

	add := flag.Bool("add", false, "Add task to the Todo List")
	list := flag.Bool("list", false, "List all tasks")
	clean := flag.Bool("clean", false, "List unfinished tasks if set true")
	complete := flag.Int("complete", 0, "Items to be completed")
	delete := flag.Int("delete", 0, "Delete todo item")
	verbose := flag.Bool("verbose", false, "Set verbose")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"Developed following the Pragmatic Bookshelf:\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
		fmt.Fprintln(flag.CommandLine.Output(), "\nto set clean or verbose \"-clean=true\" \"-verbose=true\"")
	}

	flag.Parse()

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		fmt.Print(l.PrintList(*verbose, *clean))
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *delete > 0:
		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stderr, "Invalid Option")
		os.Exit(1)
	}
}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}

	return s.Text(), nil
}
