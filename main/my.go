package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	err := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{w: f}
	sw.writeln("Don't communicate by sharing memory, share memory by communicating.")
	sw.writeln("Concurrency is not parallelism.")
	sw.writeln("Channels orchestrate; mutexes serialize.")
	sw.writeln("The bigger the interface, the weaker the abstraction.")
	sw.writeln("Make the zero value useful.")

	return sw.err
}

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}

	_, sw.err = fmt.Fprintln(sw.w, s)
}
