package main

import (
	"bufio"
	"fmt"
	"os"
)

// change dup2 to show the name of all files which have duplicated lines <- missing
func main() {
	counts := make(map[string][]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "test")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 : %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}

	for palavra, arrMaps := range counts {
		fmt.Println(palavra)
		for _, valor := range arrMaps {
			fmt.Println(valor)
		}
	}

}

func countLines(f *os.File, counts map[string][]map[string]int, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		arquivo := make(map[string]int)
		arquivo[arg]++
		if _, ok := counts[input.Text()]; ok {
			added := false
			for _, val := range counts[input.Text()] {
				if _, ok := val[arg]; ok {
					val[arg]++
					added = true
				}
			}
			if !added {
				counts[input.Text()] = append(counts[input.Text()], arquivo)
			}
		} else {
			counts[input.Text()] = append(counts[input.Text()], arquivo)
		}

	}
}
