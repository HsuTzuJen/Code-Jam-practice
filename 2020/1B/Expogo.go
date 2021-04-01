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
	var x, y int

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d %d\n", &x, &y)
		if x&1 == y&1 {
			printf("Case #%d: IMPOSSIBLE\n", t)
			continue
		}

		dx, dy := [2]byte{'E', 'W'}, [2]byte{'N', 'S'}
		ans := []byte{}
		if x < 0 {
			x = -x
			dx[0], dx[1] = dx[1], dx[0]
		}

		if y < 0 {
			y = -y
			dy[0], dy[1] = dy[1], dy[0]
		}

		for {
			if x == 0 {
				if y == 1 {
					ans = append(ans, dy[0])
					break
				} else if y == -1 {
					ans = append(ans, dy[1])
					break
				}
			}

			if y == 0 {
				if x == 1 {
					ans = append(ans, dx[0])
					break
				} else if x == -1 {
					ans = append(ans, dx[1])
					break
				}
			}

			if x&1 == 1 {
				if (y>>1)&1 == 1 {
					if ((x-1)>>1)&1 == 0 {
						x--
						ans = append(ans, dx[0])
					} else {
						x++
						ans = append(ans, dx[1])
					}
				} else { //(y>>1)&1 == 0
					if ((x+1)>>1)&1 == 1 {
						x++
						ans = append(ans, dx[1])
					} else {
						x--
						ans = append(ans, dx[0])
					}
				}
			} else { //y&1 == 1
				if (x>>1)&1 == 1 {
					if ((y-1)>>1)&1 == 0 {
						y--
						ans = append(ans, dy[0])
					} else {
						y++
						ans = append(ans, dy[1])
					}
				} else { //(x>>1)&1 == 0
					if ((y+1)>>1)&1 == 1 {
						y++
						ans = append(ans, dy[1])
					} else {
						y--
						ans = append(ans, dy[0])
					}
				}
			}

			x >>= 1
			y >>= 1
		}

		printf("Case #%d: %s\n", t, string(ans))
	}
}
