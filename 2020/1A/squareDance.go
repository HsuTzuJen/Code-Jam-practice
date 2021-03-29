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

func main() { //using some logic with C++ which got accepted but TLE in Golang, WTF???
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var T, r, c int

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d %d\n", &r, &c)
		a, up, down, left, right := make([][]int, r, r), make([][]int, r, r), make([][]int, r, r), make([][]int, r, r), make([][]int, r, r)

		alived := make([][2]int, 0, r*c)

		total := 0
		num := 0

		for i := range a {
			a[i], up[i], down[i], left[i], right[i] = make([]int, 0, c), make([]int, c, c), make([]int, c, c), make([]int, c, c), make([]int, c, c)
			up[i][0] = i - 1
			down[i][0] = i + 1
			right[i][0] = 1
			left[i][0] = -1
			alived = append(alived, [2]int{i, 0})
			for j := 1; j < c; j++ {
				scanf("%d ", &num)
				total += num
				a[i] = append(a[i], num)
				up[i][j] = i - 1
				down[i][j] = i + 1
				right[i][j] = j + 1
				left[i][j] = j - 1
				alived = append(alived, [2]int{i, j})
			}
			scanf("%d\n", &num)
			total += num
			a[i] = append(a[i], num)
		}

		ans := total

		rm := [][2]int{}

		for {
			alivedtmp := alived[:0]
			for _, ij := range alived {
				i, j := ij[0], ij[1]

				cnt := 0
				sum := 0
				if up[i][j] >= 0 {
					sum += a[up[i][j]][j]
					cnt++
				}

				if down[i][j] < r {
					sum += a[down[i][j]][j]
					cnt++
				}

				if left[i][j] >= 0 {
					sum += a[i][left[i][j]]
					cnt++
				}

				if right[i][j] < c {
					sum += a[i][right[i][j]]
					cnt++
				}

				if sum <= a[i][j]*cnt {
					alivedtmp = append(alivedtmp, ij)
				} else {
					rm = append(rm, ij)
				}

			}
			if len(rm) == 0 {
				break
			}

			for _, ij := range rm {
				i, j := ij[0], ij[1]
				total -= a[i][j]

				if up[i][j] != -1 {
					down[up[i][j]][j] = down[i][j]
				}
				if down[i][j] != r {
					up[down[i][j]][j] = up[i][j]
				}
				if left[i][j] != -1 {
					right[i][left[i][j]] = right[i][j]
				}
				if right[i][j] != c {
					left[i][right[i][j]] = left[i][j]
				}
			}
			ans += total
			alived = alivedtmp
			rm = rm[:0]

		}
		printf("Case #%d: %d\n", t, ans)
	}
}
