package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

//var f, _ = os.Open("test.txt")
//var reader *bufio.Reader = bufio.NewReader(f)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var T int
	var r, c int
	var b int

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d %d %d\n", &r, &b, &c)

		upper := 1<<63-1
		lower := 0
		cashiers := make([]cashier, c)
		for i := range cashiers {
			scanf("%d %d %d\n", &cashiers[i].m, &cashiers[i].s, &cashiers[i].p)
			//lower = min(lower, cashiers[i].p+cashiers[i].s)
			//upper = max(upper, cashiers[i].m*cashiers[i].s+cashiers[i].p)
		}

		higher := func(t int) bool {
			caps := make([]int, 0, c)
			for _, ca := range cashiers {
                caps = append(caps, max(0, min(ca.m, (t - ca.p)/ca.s)))
			}
			sort.Slice(caps, func(i, j int) bool { return caps[i] > caps[j] })

			sum := 0
			for _, cap := range caps[:r] {
				sum += cap
			}
			return sum < b
		}

		for lower < upper {
			m := (lower + upper) >> 1
			if higher(m) {
				lower = m + 1
			} else {
				upper = m
			}
		}

		printf("Case #%d: %d\n", t, lower)
	}
}

type cashier struct {
	m, s, p int
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}