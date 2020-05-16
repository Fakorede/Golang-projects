package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		sum     map[string]int
		domains []string
		total   int
		lines   int
	)

	sum = make(map[string]int)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		lines++

		fields := strings.Fields(in.Text())

		// ensure field is 2 i.e domain and visits
		if len(fields) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", fields, lines)
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])

		// if the visits field isnt a string
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", fields[1], lines)
			return
		}

		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}

		total += visits
		sum[domain] += visits
	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(domains)

	// It isn't a good idea to loop over maps in Go
	// for domain, visits := range sum {
	// 	fmt.Printf("%-30s %10d\n", domain, visits)
	// }
	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}

	fmt.Printf("\n%-30s %10d\n", "TOTAL", total)

	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}
