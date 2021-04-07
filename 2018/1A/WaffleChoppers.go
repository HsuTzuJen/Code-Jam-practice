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
	var r, c, h, v int
	var s string

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d %d %d %d\n", &r, &c, &h, &v)

		hasChip := make([][]bool, r)
		for i := range hasChip {
			hasChip[i] = make([]bool, c)
		}

		chipsCnt := 0
		rcnt := make([]int, 0, r)
		for _, row := range hasChip {
			scanf("%s\n", &s)
			cnt := 0
			for j := range row {
				if s[j] == '@' {
					row[j] = true
					cnt++
				}
			}
			rcnt = append(rcnt, cnt)
			chipsCnt += cnt
		}

		total := (h + 1) * (v + 1)
		eachHNeed := chipsCnt / (h + 1)
		hcntAcc := 0
		hCuts := make([]int, 0, h+1)
		eachNeed := chipsCnt / total
		eachGet := make([]int, h+1)
		satisfyCnt := 0

		if chipsCnt%total != 0 {
			goto bad
		} else if chipsCnt == 0 {
			goto good
		}

		for i, cnt := range rcnt {
			if hcntAcc += cnt; hcntAcc == eachHNeed {
				hCuts = append(hCuts, i+1)
				hcntAcc = 0
			} else if hcntAcc > eachHNeed {
				goto bad
			}
		}

		for j := 0; j < c; j++ {
			prevhcut := 0
			for k, hcut := range hCuts {
				for i := prevhcut; i < hcut; i++ {
					if hasChip[i][j] {
						if eachGet[k]++; eachGet[k] == eachNeed {
							satisfyCnt++
						} else if eachGet[k] > eachNeed {
							goto bad
						}
					}
					prevhcut = hcut
				}
			}
			if satisfyCnt == len(eachGet) {
				for k := range eachGet {
					eachGet[k] = 0
				}
				satisfyCnt = 0
			}
		}

	good:
		printf("Case #%d: POSSIBLE\n", t)
		continue
	bad:
		printf("Case #%d: IMPOSSIBLE\n", t)
	}
}