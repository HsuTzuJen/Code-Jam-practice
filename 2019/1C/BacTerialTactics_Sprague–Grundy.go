package main

import (
	"bufio"
	"container/heap"
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

	var grundy [16][16][16][16]int
	var dead [][]bool

	var solve func(r0, r1, c0, c1 int) int
	solve = func(r0, r1, c0, c1 int) int {
		if r0 == r1 || c0 == c1 {
			return 0
		}
		if tmp := grundy[r0][r1][c0][c1]; tmp != 0 {
			if tmp == -1 {
				return 0
			}
			return tmp
		}

		set := map[int]struct{}{}
		m := minheap{}

		for i := r0; i < r1; i++ {
			var sgb int
			for j := c0; j < c1; j++ {
				if dead[i][j] {
					goto nextr
				}
			}
			sgb = solve(r0, i, c0, c1) ^ solve(i+1, r1, c0, c1)
			if _, ok := set[sgb]; !ok {
				set[sgb] = struct{}{}
				heap.Push(&m, sgb)
			}
		nextr:
		}

		for j := c0; j < c1; j++ {
			var sgb int
			for i := r0; i < r1; i++ {
				if dead[i][j] {
					goto nextc
				}
			}
			sgb = solve(r0, r1, c0, j) ^ solve(r0, r1, j+1, c1)
			if _, ok := set[sgb]; !ok {
				set[sgb] = struct{}{}
				heap.Push(&m, sgb)
			}
		nextc:
		}

		sg := 0
		for len(m) > 0 {
			if sg < m[0] {
				break
			}
			heap.Pop(&m)
			sg++
		}
		if sg == 0 {
			grundy[r0][r1][c0][c1] = -1
		} else {
			grundy[r0][r1][c0][c1] = sg
		}
		return sg
	}

	var T int
	var r, c int

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d %d\n", &r, &c)

		dead = make([][]bool, r)
		grundy = [16][16][16][16]int{}

		for i := 0; i < r; i++ {
			s := ""
			scanf("%s\n", &s)
			dead[i] = make([]bool, c)
			for j, ch := range s {
				if ch == '#' {
					dead[i][j] = true
				}
			}
		}

		ans := 0

		if solve(0, r, 0, c) > 0 {
			for i, row := range dead {
				for _, d := range row {
					if d {
						goto nextr
					}
				}
				if solve(0, i, 0, c)^solve(i+1, r, 0, c) == 0 {
					ans += c
				}
			nextr:
			}
			for j := range dead[0] {
				for i := range dead {
					if dead[i][j] {
						goto nextc
					}
				}
				if solve(0, r, 0, j)^solve(0, r, j+1, c) == 0 {
					ans += r
				}
			nextc:
			}
		}

		printf("Case #%d: %d\n", t, ans)
	}
}