package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// change de program to output the result to a file
func main() {
	start := time.Now()
	fullDate := time.Now().Format("02-01-2006-15-04-05")
	fileWithResult, err := os.Create(fullDate + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fileWithResult.WriteString(fmt.Sprint(<-ch) + "\n") // receive from channel ch
	}
	fileWithResult.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
