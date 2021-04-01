package main

import (
	"bufio"
	"fmt"
	"math/rand"
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
	var a, b int

	scanf("%d %d %d\n", &T, &a, &b)
	for t := 1; t <= T; t++ {

		askcnt := 0
		res := ""

		var x, y, sx, sy int

		var q [][2]int
		shooted := [37][37]bool{}

		push := func(i, j int) {
			ii, jj := i-x+18, j-y+18
			if shooted[ii][jj] {
				return
			}

			shooted[ii][jj] = true
			q = append(q, [2]int{i, j})
		}

		var low, high int

		for {
			askcnt++
			sx, sy = rand.Intn(int(2e9+1))-int(1e9), rand.Intn(int(2e9+1))-int(1e9)
			fmt.Println(sx, sy)
			scanf("%s\n", &res)
			if res == "CENTER" {
				goto next
			} else if res == "HIT" {
				break
			}
		}

		low, high = int(-1e9), sy
		for low < high {
			if askcnt++; askcnt > 300 {
				return
			}

			mid := (low + high) >> 1
			fmt.Println(sx, mid)
			scanf("%s\n", &res)
			if res == "CENTER" {
				goto next
			} else if res == "HIT" {
				high = mid
			} else { //MISS
				low = mid + 1
			}
		}
		y = low

		low, high = sy, int(1e9)
		for low < high {
			if askcnt++; askcnt > 300 {
				return
			}

			mid := (low + high) >> 1
			fmt.Println(sx, mid)
			scanf("%s\n", &res)
			if res == "CENTER" {
				goto next
			} else if res == "HIT" {
				if low = mid + 1; low == high {
					low--
					break
				}
			} else { //MISS
				high = mid
			}
		}
		y = (y + low) >> 1

		low, high = int(-1e9), sx
		for low < high {
			if askcnt++; askcnt > 300 {
				return
			}

			mid := (low + high) >> 1
			fmt.Println(mid, y)
			scanf("%s\n", &res)
			if res == "CENTER" {
				goto next
			} else if res == "HIT" {
				high = mid
			} else { //MISS
				low = mid + 1
			}
		}
		x = low

		low, high = sx, int(1e9)
		for low < high {
			if askcnt++; askcnt > 300 {
				return
			}

			mid := (low + high) >> 1
			fmt.Println(mid, y)
			scanf("%s\n", &res)
			if res == "CENTER" {
				goto next
			} else if res == "HIT" {
				if low = mid + 1; low == high {
					low--
					break
				}
			} else { //MISS
				high = mid
			}
		}
		x = (x + low) >> 1

		shooted[18][18] = true
		q = [][2]int{{x, y}}

		for qlen := len(q); ; qlen = len(q) {
			for i := 0; i < qlen; i++ {
				if askcnt++; askcnt > 300 {
					return
				}
				xy := q[i]
				fmt.Println(xy[0], xy[1])
				scanf("%s\n", &res)
				if res == "CENTER" {
					goto next
				}

				push(xy[0]+1, xy[1])
				push(xy[0]-1, xy[1])
				push(xy[0], xy[1]+1)
				push(xy[0], xy[1]-1)
				push(xy[0]+1, xy[1]+1)
				push(xy[0]-1, xy[1]-1)
				push(xy[0]-1, xy[1]+1)
				push(xy[0]+1, xy[1]-1)
			}
			q = q[qlen:]
		}
		return

	next:
		//printf("Case #%d: %s\n", t, string(ans))
	}
}
