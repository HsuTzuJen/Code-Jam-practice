package main

import (
	"bufio"
	"fmt"
	"os"
)

var pascal = [1001][]int{}
var prefixsums = [1001][]int{}

func init() {
	pascal[0] = append(pascal[0], 1)
	for i := 1; i < 1001; i++ {
		pascal[i] = make([]int, i+1, i+1)
		prefixsums[i] = make([]int, i+1, i+1)
		pascal[i][0], pascal[i][i], prefixsums[i][0] = 1, 1, 1
		for j := 1; j < i; j++ {
			pascal[i][j] = pascal[i-1][j] + pascal[i-1][j-1]
			prefixsums[i][j] = prefixsums[i][j-1] + pascal[i][j]
		}
		prefixsums[i][i] = prefixsums[i][i-1] + pascal[i][i]
	}
}

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
	var n int

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d\n", &n)
		printf("Case #%d:\n", t)

		i, j := 0, 0
		for {
			printf("%d %d\n", i+1, j+1)
			if n -= pascal[i][j]; n == 0 {break}

			if i&1 == 1 && n >= prefixsums[i+1][j+1] { //checks from next maximum prefixsum we could go
				i++; j++; continue
			}

			if n >= prefixsums[i+1][j] { //if we could not, check next
				i++; continue
			}

			j-- //the last option
		}
	}
}
