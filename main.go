package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

type Name string

func (n *Name) String() string {
	return fmt.Sprint(*n)
}

func (n *Name) Set(value string) error {
	if len(*n) > 0 {
		return errors.New("name flag already set")
	}
	*n = Name(value)
	return nil
}

var name Name

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		return
	}
	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.Var(&name, "name", "Go tour")
		_ = goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.Var(&name, "n", "php tour")
		_ = phpCmd.Parse(args[1:])
	}
	log.Printf("name :%s", name)
}
