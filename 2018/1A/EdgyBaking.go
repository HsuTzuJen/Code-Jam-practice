
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

//var f, _ = os.Open("test.txt")
//var reader *bufio.Reader = bufio.NewReader(f)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

var ws = [100]float64{}
var hs = [100]float64{}
var add = [100][2]float64{}

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var T int
	var n int
	var w, h, p float64

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d %f\n", &n, &p)

		pp := p / 2
		for i := 0; i < n; i++ {
			scanf("%f %f\n", &w, &h)
			pp -= w + h
			add[i][0] = min(w, h)
			add[i][1] = math.Sqrt(w*w+h*h) - add[i][0]
		}

		dp := make([]map[float64]float64, n)
		for i := range dp {
			dp[i] = map[float64]float64{}
		}

		var dfs func(i int, pp float64) float64
		dfs = func(i int, pp float64) float64 {
			if i == n {
				return pp
			}

			if res, ok := dp[i][pp]; ok {
				return res
			}

			res := dfs(i+1, pp)
			if pp >= add[i][0] {
				tmp := dfs(i+1, pp-add[i][0])
				if tmp -= min(tmp, add[i][1]); tmp < res {
					res = tmp
				}
			}

			dp[i][pp] = res

			return res
		}

		printf("Case #%d: %f\n", t, p-2*dfs(0, pp))
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}

	return b
}