package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	 par "github.com/stevedesilva/gofuncPointers/parser"
)
// go run main.go < log.txt 
func main() {
	p := par.NewParser()

	// Scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.Lines++

		parsed, err := par.Parse(p, in.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		p = par.Update(p, parsed)
	}

	// Print the visits per domain
	sort.Strings(p.Domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.Domains {
		parsed := p.Sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.Visits)
	}

	// Print the total visits for all domains
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.Total)

	// Let's handle the error
	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}