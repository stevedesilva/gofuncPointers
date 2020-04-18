package main

import (
	"bufio"
	"fmt"
	"os"

	par "github.com/stevedesilva/gofuncPointers/parser"
)

// go run main.go < log.txt
func main() {
	p := par.NewParser()

	// Scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {

		parsed, err := par.Parse(&p, in.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		par.Update(&p, parsed)
	}
    // no need for pointer sice we ate not changing the value
	par.Summarise(p)

	// Let's handle the error
	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}
