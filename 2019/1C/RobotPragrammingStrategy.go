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
	var n int
	var s string

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d\n", &n)
		asm := make(map[byte]struct{}, n)
		as := make([]string, 0, n)
		for i := 0; i < n; i++ {
			scanf("%s\n", &s)
			asm[byte(i)] = struct{}{}
			as = append(as, s)
		}

		var ans []byte

		var dfs func(i int, prog []byte) bool
		dfs = func(i int, prog []byte) bool {
			if len(asm) == 0 {
				ans = prog
				return true
			}

			if len(prog) == 500 {
				return false
			}

			rcnt, pcnt, scnt := 0, 0, 0
			for j := range asm {
				a := as[j]
				switch a[i%len(a)] {
				case 'R':
					rcnt++
				case 'P':
					pcnt++
				default: //'S'
					scnt++
				}
			}
			if rcnt > 0 && pcnt > 0 && scnt > 0 {
				return false
			}

			for _, choice := range []byte{'R', 'S', 'P'} {
				var del byte

				switch choice {
				case 'R':
					if pcnt > 0 || scnt == 0 {
						continue
					}
					del = 'S'
				case 'P':
					if scnt > 0 || rcnt == 0 {
						continue
					}
					del = 'R'
				default: //'S'
					if rcnt > 0 || pcnt == 0 {
						continue
					}
					del = 'P'
				}
				tmp := []byte{}
				for j := range asm {
					a := as[j]
					if a[i%(len(a))] == del {
						delete(asm, j)
						tmp = append(tmp, j)
					}
				}
				if dfs(i+1, append(prog, choice)) {
					return true
				}
				for _, j := range tmp {
					asm[j] = struct{}{}
				}
			}
			return false
		}

		if dfs(0, make([]byte, 0, 500)) {
			printf("Case #%d: %s\n", t, string(ans))
		} else {
			printf("Case #%d: IMPOSSIBLE\n", t)
		}
	}
}

