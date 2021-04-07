package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//var f, _ = os.Open("test.txt")
//var reader *bufio.Reader = bufio.NewReader(f)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string) interface{} {
	var s string
	for {
		fmt.Fscanf(reader, "%s", &s)
		if s != "" {
			switch f {
			case "d":
				d, _ := strconv.Atoi(s)
				return d
			case "s":
				return s
			}
		}
	}
}

func getarr(n int) []int {
	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		num := scanf("d").(int)
		a = append(a, num)
	}
	return a
}


func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	T := scanf("d").(int)
	for t := 1; t <= T; t++ {
		n := scanf("d").(int)

		ws := getarr(n)
		wsmax := make([]int, n, n)
		for i, w := range ws {
			wsmax[i] = w * 6
		}

		dp := [141]int{}
		for i := 1; i < 141; i++ {
			dp[i] = 1<<63 - 1
		}

		for i, w := range ws {
			for cnt := min(139, i); cnt >= 0; cnt-- {
				if dp[cnt] <= wsmax[i] {
					dp[cnt+1] = min(dp[cnt+1], dp[cnt]+w)
				}
			}
		}

		ans := 0
		for cnt, wsum := range dp {
			if wsum != 1<<63-1 {
				ans = cnt
			}
		}

		printf("Case #%d: %d\n", t, ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}