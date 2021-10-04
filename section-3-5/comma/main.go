// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//Exercise:
// change de program to work with Buffer
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	sLen := len(s)
	if sLen <= 3 {
		return s
	}
	lastDigits := s[sLen-3 : sLen]
	firstDigits := s[:sLen-3]
	fmt.Fprintf(&buf, "%s", firstDigits)
	buf.WriteString(",")
	fmt.Fprintf(&buf, "%s", lastDigits)
	return buf.String()

}
