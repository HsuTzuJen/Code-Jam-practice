package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var T int
	var n int
	var d int

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d %d\n", &n, &d)

		sizeCnt := map[float64][][2]int{}
		slices := make([]int, 0, n)

		s := 0
		scanf("%d", &s)
		slices = append(slices, s)
		for cnt := 1; cnt <= d; cnt++ {
			size := float64(s) / float64(cnt)
			sizeCnt[size] = append(sizeCnt[size], [2]int{cnt, s})
		}

		for i := 1; i < n; i++ {
			scanf(" %d", &s)
			slices = append(slices, s)
			for cnt := 1; cnt <= d; cnt++ {
				size := float64(s) / float64(cnt)
				sizeCnt[size] = append(sizeCnt[size], [2]int{cnt, s})
			}
		}
		scanf("\n")

		sort.Ints(slices)

		ans := d - 1
		for size, arr := range sizeCnt {
			sort.Slice(arr, func(i, j int) bool { return arr[i][0] < arr[j][0] })

			totalCuts := 0
			dtmp := 0
			for _, cnt := range arr {
				if need := d - dtmp; need < cnt[0] {
					totalCuts += need
					dtmp += need
					break
				} else {
					totalCuts += cnt[0] - 1
					if dtmp += cnt[0]; dtmp == d {
						break
					}
				}
			}
			if dtmp < d {
				j := 0
				for _, s := range slices {
					if j < len(arr) && s == arr[j][1] {
						j++
						continue
					}
					if float64(s) < size {
						continue
					}

					cuts := min(int(float64(s)/size), d-dtmp)
					totalCuts += cuts
					if dtmp += cuts; dtmp == d {
						break
					}
				}
			}
			if dtmp == d && totalCuts < ans {
				ans = totalCuts
			}
		}

		printf("Case #%d: %d\n", t, int(ans))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
