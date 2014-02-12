package ttt

import (
	"fmt"
	"io"
)

type ConsoleIO struct {
	Writer io.Writer
	Reader io.Reader
}

func (c ConsoleIO) Read() string {
	var result string
	fmt.Fscanln(c.Reader, &result)
	return string(result)
}

func (c ConsoleIO) Println(input string) {
	fmt.Fprintln(c.Writer, input)
}

func (c ConsoleIO) Printf(input string, a ...interface{}) {
	fmt.Fprintf(c.Writer, input, a...)
}
