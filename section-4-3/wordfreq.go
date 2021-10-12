package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	file := args[0]
	m := make(map[string]int)
	f, _ := os.Open(file)
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		m[s.Text()]++
		// fmt.Println(s.Text())
	}
	for k, v := range m {
		fmt.Printf("Palavra: %v Ocorrencias:%v\n", k, v)
	}
}
