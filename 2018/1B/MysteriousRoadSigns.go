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

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var T int
	var s int

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d\n", &s)

		d := make([][2]int, s)
		var a, b int

		for i := 0; i < s; i++ {
			scanf("%d %d %d\n", &d[i][0], &a, &b)
			d[i][1] = d[i][0] - b
			d[i][0] += a
		}

		maxLen := make([]int, s)

		var dfs func(L, R int)
		dfs = func(L, R int) {
			if R-L < 2 {
				maxLen[L] = max(maxLen[L], R-L+1)
				return
			}

			M := (L + R) >> 1

			dfs(L, M-1)
			dfs(M+1, R)

			for fix := 0; fix < 2; fix++ {
				for extanddir := 0; extanddir < 2; extanddir++ {
					l, r := M, M
					var m, n int
					if fix == 0 { //fix m
						m, n = d[M][0], -1<<63
					} else { //fix n
						m, n = -1<<63, d[M][1]
					}
					for l > L && (m == d[l-1][0] || n == d[l-1][1]) {
						l--
					}

					for r < R && (m == d[r+1][0] || n == d[r+1][1]) {
						r++
					}

					if extanddir == 0 {
						if l > L {
							if fix == 0 { //already fixed m
								n = d[l-1][1]
							} else { //already fixed n
								m = d[l-1][0]
							}
						}
					} else {
						if r < R {
							if fix == 0 {
								n = d[r+1][1]
							} else {
								m = d[r+1][0]
							}
						}
					}
					for l > L && (m == d[l-1][0] || n == d[l-1][1]) {
						l--
					}

					for r < R && (m == d[r+1][0] || n == d[r+1][1]) {
						r++
					}

					maxLen[l] = max(maxLen[l], r-l+1)
				}
			}
		}

		dfs(0, s-1)

		cnt4maxl := 0
		maxl := 0
		for _, l := range maxLen {
			if l > maxl {
				maxl = l
				cnt4maxl = 1
			} else if l == maxl {
				cnt4maxl++
			}
		}

		printf("Case #%d: %d %d\n", t, maxl, cnt4maxl)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}