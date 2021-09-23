package main

import (
	"fmt"
	"os"
)

// 1.1 change the program to show os.Args[0]
// 1.2 change the program to the index and value for each argument, one per line
// 1.3 See the differente of execution between inveficient version to output string and strings.Join <- missing
func main() {

	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("%d\t%s\n", i, os.Args[i])
	}

}
