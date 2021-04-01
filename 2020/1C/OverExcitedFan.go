package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var T int
	var x, y int
	var path string

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d %d %s\n", &x, &y, &path)

		for i, d := range path {
			switch d {
			case 'N':
				y++
			case 'S':
				y--
			case 'E':
				x++
			default:
				x--
			}
			if dis := abs(x) + abs(y); dis <= i+1 {
				printf("Case #%d: %d\n", t, i+1)
				goto next
			}

		}
		printf("Case #%d: IMPOSSIBLE\n", t)
	next:
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}
