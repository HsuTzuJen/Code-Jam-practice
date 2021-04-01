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
	var m, u int
	var str string

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d\n", &u)

		chsCnt := [256]int{}
		showed := [256]bool{}
		for i := 0; i < 10000; i++ {
			scanf("%d %s\n", &m, &str)
			chsCnt[str[0]]++
			for _, ch := range str {
				showed[ch] = true
			}
		}
		strs := make([]string, 0, 10)
		cnts := make([]int, 0, 10)
		for ch, cnt := range chsCnt {
			if cnt > 0 {
				showed[ch] = false
				strs = append(strs, string([]byte{byte(ch)}))
				cnts = append(cnts, cnt)
			}
		}

		sort.Slice(cnts, func(i, j int) bool {
			if cnts[i] > cnts[j] {
				strs[i], strs[j] = strs[j], strs[i]
				return true
			}
			return false
		})

		printf("Case #%d: ", t)

		for i, ok := range showed {
			if ok {
				printf("%s", string(i))
				break
			}
		}

		for _, s := range strs {
			printf("%s", s)
		}
		printf("\n")

	}
}
