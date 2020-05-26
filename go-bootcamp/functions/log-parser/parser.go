package main

import (
	"fmt"
	"strconv"
	"strings"
)

// store the parsed result for a domain
type result struct {
	domain string
	visits int
	// add more metrics if needed.
}

// keeps track of the parsing
type parser struct {
	sum     map[string]result
	domains []string
	total   int
	lines   int
}

func newParser() parser {
	return parser{sum: make(map[string]result)}
}

func parse(p parser, line string) (parsed result, err error) {

	fields := strings.Fields(line)

	// ensure field is 2 i.e domain and visits
	if len(fields) != 2 {
		err = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return
	}

	parsed.domain = fields[0]
	parsed.visits, err = strconv.Atoi(fields[1])

	// if the visits field isnt a string
	if parsed.visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.lines)
		return
	}

	return
}
