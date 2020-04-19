package main

import (
	"bufio"
	"os"

	par "github.com/stevedesilva/gofuncPointers/parser"
)

// go run main.go < log.txt
func main() {
	p := par.NewParser()

	// Scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		parsed := par.Parse(p, in.Text())
		par.Update(p, parsed)
	}
	// no need for pointer sice we ate not changing the value
	par.Summarise(p)

	par.DumpErrs([]error{par.Err(p), in.Err()})
}
