package main

import (
	"errors"
	"fmt"
	"log"
	"tour/cmd"
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
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
