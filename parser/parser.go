package parser

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Result stores the parsed Result for a domain
type Result struct {
	Domain string
	Visits int
	// add more metrics if needed
}

// Parser keep tracks of the parsing
type Parser struct {
	Sum     map[string]Result // metrics per domain
	Domains []string          // unique domain names
	Total   int               // total visits for all domains
	Lines   int               // number of parsed lines (for the error messages)
}

// NewParser constructs, initializes and returns a new Parser
func NewParser() Parser {
	return Parser{Sum: make(map[string]Result)}
}

// Parse parses a log line and returns the parsed Result with an error
func Parse(p *Parser, line string) (parsed Result, err error) {
	p.Lines++
	fields := strings.Fields(line)
	if len(fields) != 2 {
		err = fmt.Errorf("wrong input: %v (line #%d)", fields, p.Lines)
		return
	}

	parsed.Domain = fields[0]

	parsed.Visits, err = strconv.Atoi(fields[1])
	if parsed.Visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.Lines)
		return
	}

	return
}

// Update updates the Parser for the given parsing Result
func Update(p *Parser, parsed Result) {
	domain, visits := parsed.Domain, parsed.Visits

	// Collect the unique domains
	if _, ok := p.Sum[domain]; !ok {
		p.Domains = append(p.Domains, domain)
	}

	// Keep track of total and per domain visits
	p.Total += visits

	// create and assign a new copy of `visit`
	p.Sum[domain] = Result{
		Domain: domain,
		Visits: visits + p.Sum[domain].Visits,
	}
	// addressable value example
	//  p.Sum[domain].Visits += visits
	// _ = &p.Sum[domain]
	// clone := p.Sum[domain]
	// _ = &clone

}

// Summarise func
func Summarise(p Parser) {
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
}
