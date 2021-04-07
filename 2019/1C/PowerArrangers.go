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

var poss = make([]int, 119)

func init() {
	for i := range poss {
		poss[i] = 1 + i*5
	}
}

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var T int
	var f int
	var s string

	scanf("%d %d\n", &T, &f)
	for t := 1; t <= T; t++ {
		usingposs := make([]int, 119)
		copy(usingposs, poss)

		found := [5]bool{}
		for _, target := range []int{23, 5, 1, 0} {
			eachpos := make([][]int, 5)
			for _, i := range usingposs {
				fmt.Println(i)
				scanf("%s\n", &s)
				eachpos[s[0]-'A'] = append(eachpos[s[0]-'A'], i+1)
			}
			for ch, poss := range eachpos {
				if found[ch] {
					continue
				}
				if len(poss) == target {
					printf("%s", string(byte(ch)+'A'))
					found[ch] = true
					usingposs = poss
				}
			}
		}
		for ch := range found {
			if !found[ch] {
				printf("%s\n", string(byte(ch)+'A'))
				break
			}
		}
		writer.Flush()
		res := ""
		scanf("%s\n", &res)
		if res == "N" {
			return
		}
	}
}