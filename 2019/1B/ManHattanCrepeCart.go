
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
	var p, q int
	var x, y int
	var d string

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d %d\n", &p, &q)

		bitxw := make([]int, q+2)
		bitxe := make([]int, q+2)
		bityn := make([]int, q+2)
		bitys := make([]int, q+2)

		update := func(s int, bit []int) {
			for s < len(bit) {
				bit[s]++
				s += s & (-s)
			}
		}

		update2 := func(s int, bit []int) {
			for s > 0 {
				bit[s]++
				s -= s & (-s)
			}
		}

		get := func(x int, bit []int) (res int) {
			for x > 0 {
				res += bit[x]
				x -= x & (-x)
			}
			return
		}

		get2 := func(x int, bit []int) (res int) {
			for x < len(bit) {
				res += bit[x]
				x += x & (-x)
			}
			return
		}

		for i := 0; i < p; i++ {
			scanf("%d %d %s\n", &x, &y, &d)
			switch d {
			case "S":
				update2(y, bitys)
			case "N":
				update(y+2, bityn)
			case "W":
				update2(x, bitxw)
			default: //"E"
				update(x+2, bitxe)
			}
		}

		maxx, tx := -1, 0

		for i := 1; i < len(bitxw); i++ {
			tmp := get2(i, bitxw) + get(i, bitxe)
			if tmp > maxx {
				maxx = tmp
				tx = i
			}
		}

		maxy, ty := -1, 0

		for i := 1; i < len(bityn); i++ {
			if tmp := get2(i, bitys) + get(i, bityn); tmp > maxy {
				maxy = tmp
				ty = i
			}
		}

		printf("Case #%d: %d %d\n", t, tx-1, ty-1)
	}
}