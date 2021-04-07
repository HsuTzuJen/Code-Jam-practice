package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

var precompute = [21][21][]int{}

func init() {
	visited := [400]bool{}

	for r := 2; r < 21; r++ {
		for c := 2; c < 21; c++ {
			canvis := make([][]int, r*c)
			for i := 0; i < r; i++ {
				for j := 0; j < c; j++ {
					canvistmp := []int{}
					for x := 0; x < r; x++ {
						for y := 0; y < c; y++ {
							if x == i || y == j || x-y == i-j || x+y == i+j {
								continue
							}
							canvistmp = append(canvistmp, x*c+y)
						}
					}
					rand.Shuffle(len(canvistmp), func(i, j int) { canvistmp[i], canvistmp[j] = canvistmp[j], canvistmp[i] })
					canvis[i*c+j] = canvistmp
				}
			}

			target := r * c
			targetPath := make([]int, 0, target)

			var dfs func(u int, path []int) bool
			dfs = func(u int, path []int) bool {
				if path = append(path, u); len(path) == target {
					targetPath = path
					return true
				}
				visited[u] = true
				defer func() { visited[u] = false }()

				for _, v := range canvis[u] {
					if visited[v] {
						continue
					}
					if dfs(v, path) {
						return true
					}
				}
				return false
			}

			for x := 0; x < r; x++ {
				for y := 0; y < c; y++ {
					if dfs(x*c+y, targetPath) {
						goto next
					}
				}
			}
		next:
			precompute[r][c] = targetPath
		}
	}
}

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var T, r, c int

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d %d\n", &r, &c)
		if len(precompute[r][c]) == 0 {
			printf("Case #%d: IMPOSSIBLE\n", t)
		} else {
			printf("Case #%d: POSSIBLE\n", t)
			for _, u := range precompute[r][c] {
				printf("%d %d\n", u/c+1, u%c+1)
			}
		}
	}
}
