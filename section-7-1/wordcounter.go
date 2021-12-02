package main

import (
	"fmt"
	"strings"
)

type ByteCounter int

func (w *ByteCounter) Write(p []byte) (int, error) {
	*w += ByteCounter(len(p))
	return 0, nil
}

type Wordcounter int

func (w *Wordcounter) Write(p []byte) (int, error) {
	// *w += Wordcounter(len(strings.Fields(string(p))))
	return len(strings.Fields(string(p))), nil
}

func main() {
	var b ByteCounter
	b.Write([]byte("hello"))
	fmt.Println(b)

	var w Wordcounter
	w2, _ := w.Write([]byte("uma duas tres"))
	fmt.Println(w2)
}
