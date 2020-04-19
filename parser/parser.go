package parser

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Result stores the r Result for a domain
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
	Lines   int               // number of r lines (for the error messages)
	Lerr    error
}

// NewParser constructs, initializes and returns a new Parser
func NewParser() *Parser {
	return &Parser{Sum: make(map[string]Result)}
}

// Parse parses a log line and returns the r Result with an error
func Parse(p *Parser, line string) (r Result) {
	if p.Lerr != nil {
		return
	}

	p.Lines++

	fields := strings.Fields(line)
	if len(fields) != 2 {
		p.Lerr = fmt.Errorf("wrong input: %v (line #%d)", fields, p.Lines)
		return
	}

	r.Domain = fields[0]

	var err error
	r.Visits, err = strconv.Atoi(fields[1])
	if r.Visits < 0 || err != nil {
		p.Lerr = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.Lines)
	}

	return
}

// Update updates the Parser for the given parsing Result
func Update(p *Parser, r Result) {
	if p.Lerr != nil {
		return
	}

	// Collect the unique domains
	if _, ok := p.Sum[r.Domain]; !ok {
		p.Domains = append(p.Domains, r.Domain)
	}

	// Keep track of total and per domain visits
	p.Total += r.Visits

	// create and assign a new copy of `visit`
	p.Sum[r.Domain] = Result{
		Domain: r.Domain,
		Visits: r.Visits + p.Sum[r.Domain].Visits,
	}
	// addressable value example
	//  p.Sum[domain].Visits += visits
	// _ = &p.Sum[domain]
	// clone := p.Sum[domain]
	// _ = &clone

}

// Summarise func
func Summarise(p *Parser) {
	// Print the visits per domain
	sort.Strings(p.Domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.Domains {
		fmt.Printf("%-30s %10d\n", domain, p.Sum[domain].Visits)
	}

	// Print the total visits for all domains
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.Total)
}

// DumpErrs handler
func DumpErrs(errs []error) {
	for _,err := range errs {
		if err != nil {
			fmt.Println("> Err:", err)
		}
	}
}

// Err func
func Err(p *Parser) error {
	return p.Lerr
}