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

var dppath = [10][10][10][10][][2]int{}
var path44 = [][2]int{}

func init() {
	visited := [10][10]bool{}
	var path [][2]int
	var r, c, target int

	var dfs func(x, y, cnt int) bool
	dfs = func(x, y, cnt int) bool {
		if cnt++; cnt == target {
			path[cnt-1][0], path[cnt-1][1] = x+1, y+1
			return true
		}

		visited[x][y] = true
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				if visited[i][j] || x == i || y == j || x-y == i-j || x+y == i+j {
					continue
				}

				if dfs(i, j, cnt) {
					path[cnt-1][0], path[cnt-1][1] = x+1, y+1
					return true
				}
			}
		}
		visited[x][y] = false

		return false
	}

	r = 2
	for c = 5; c < 10; c++ {
		target = r * c
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				path = make([][2]int, target)
				dfs(i, j, 0)
				dppath[r][c][i][j] = path
				visited = [10][10]bool{}
			}
		}
	}
	c = 2
	for r = 5; r < 10; r++ {
		target = r * c
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				path = make([][2]int, target)
				dfs(i, j, 0)
				dppath[r][c][i][j] = path
				visited = [10][10]bool{}
			}
		}
	}
	r = 3
	for c = 4; c < 8; c++ {
		target = r * c
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				path = make([][2]int, target)
				dfs(i, j, 0)
				dppath[r][c][i][j] = path
				visited = [10][10]bool{}
			}
		}
	}
	c = 3
	for r = 4; r < 8; r++ {
		target = r * c
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				path = make([][2]int, target)
				dfs(i, j, 0)
				dppath[r][c][i][j] = path
				visited = [10][10]bool{}
			}
		}
	}
	r, c, target = 4, 4, 16
	path = make([][2]int, target)
	dfs(0, 0, 0)
	path44 = path
}

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var T, r, c int

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d %d\n", &r, &c)
		if r+c < 7 {
			printf("Case #%d: IMPOSSIBLE\n", t)
			continue
		}
		var path [][2]int
		if r == 4 && c == 4 {
			path = path44
		} else {
			path = make([][2]int, 0, r*c)
			si, sj := 0, 0

			if r < c {
				for si < r {
					rlen, cbase := 3, 4
					if si+2 == r || si+4 == r {
						rlen, cbase = 2, 5
					}

					for sj < c {
						var clen int
						if c-sj-cbase < cbase {
							clen = c - sj
						} else {
							clen = cbase
						}
						for i := 0; i < rlen; i++ {
							for j := 0; j < clen; j++ {
								subpath := dppath[rlen][clen][i][j]
								if len(path) > 0 {
									x, y := path[len(path)-1][0], path[len(path)-1][1]
									ti, tj := subpath[0][0]+si, subpath[0][1]+sj
									if x == ti || y == tj || x-y == ti-tj || x+y == ti+tj {
										continue
									}
								}
								for _, xy := range subpath {
									xy[0] += si
									xy[1] += sj
									path = append(path, xy)
								}
								goto next
							}
						}
					next:
						sj += clen
					}
					sj = 0
					si += rlen
				}
			} else {
				for sj < c {
					clen, rbase := 3, 4
					if sj+2 == c || sj+4 == c {
						clen, rbase = 2, 5
					}

					for si < r {
						var rlen int
						if r-si-rbase < rbase {
							rlen = r - si
						} else {
							rlen = rbase
						}
						for i := 0; i < rlen; i++ {
							for j := 0; j < clen; j++ {
								subpath := dppath[rlen][clen][i][j]
								if len(path) > 0 {
									x, y := path[len(path)-1][0], path[len(path)-1][1]
									ti, tj := subpath[0][0]+si, subpath[0][1]+sj
									if x == ti || y == tj || x-y == ti-tj || x+y == ti+tj {
										continue
									}
								}
								for _, xy := range subpath {
									xy[0] += si
									xy[1] += sj
									path = append(path, xy)
								}
								goto next2
							}
						}
					next2:
						si += rlen
					}
					si = 0
					sj += clen
				}
			}
		}
		printf("Case #%d: POSSIBLE\n", t)
		for _, xy := range path {
			printf("%d %d\n", xy[0], xy[1])
		}
	}
}
