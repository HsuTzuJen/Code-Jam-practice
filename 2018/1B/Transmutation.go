package main

import (
	"bufio"
	"fmt"
	"os"
)

//var f, _ = os.Open("test.txt")
//var reader *bufio.Reader = bufio.NewReader(f)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func getarr(n int) []int {
	var num int
	a := make([]int, 0, n)
	for i := 1; i < n; i++ {
		scanf("%d ", &num)
		a = append(a, num)
	}
	scanf("%d\n", &num)
	a = append(a, num)
	return a
}

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var T int
	var n int

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d\n", &n)
		formulas := make([][2]int, n, n)
		for i := range formulas {
			scanf("%d %d\n", &formulas[i][0], &formulas[i][1])
			formulas[i][0]--
			formulas[i][1]--
		}
		ms := getarr(n)

		visited := make([]bool, n)
		var dfs func(m, need int) bool
		dfs = func(m, need int) bool {
			if visited[m] {
				return false
			}
			visited[m] = true
			defer func() {
				visited[m] = false
			}()
			if ms[m] >= need {
				ms[m] -= need
				return true
			}

			stillneed := need - ms[m]
			ms[m] = 0

			if dfs(formulas[m][0], stillneed) && dfs(formulas[m][1], stillneed) {
				return true
			}

			return false
		}

		mstmp := make([]int, n, n)

		l, r := 0, 1<<63-1
		for l+1 < r {
			copy(mstmp, ms)
			m := (l + r) >> 1
			if dfs(0, m) {
				l = m
			} else {
				r = m
			}
			copy(ms, mstmp)
		}

		printf("Case #%d: %d\n", t, l)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}