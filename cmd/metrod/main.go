package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	m := NewMain()
	if err := m.Run(os.Args[1:]...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}
}

type Main struct {
	Stdin io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func NewMain() *Main {
	return &Main{Stdin: os.Stdin,Stdout: os.Stdout, Stderr: os.Stderr}
}

func (m Main) Run(args ...string) interface{} {
	name, args :=
}