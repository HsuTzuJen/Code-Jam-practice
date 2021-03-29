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

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d\n", &n)
		ps := make([]string, n, n)
		for i := range ps {
			scanf("%s\n", &ps[i])
		}

		prefix, suffix := []byte{}, []byte{}
		for { //compare prefix
			pre := byte('*')
			for i := range ps {
				if len(ps[i]) == 0 {continue}

				if ps[i][0] != '*' {
					if pre == '*' {
						pre = ps[i][0]
					} else if pre != ps[i][0] { //prefix not match
						goto bad
					}
					ps[i] = ps[i][1:]
				}
			}

			if pre == '*' {break} //no more prefix needs matching

			prefix = append(prefix, pre)
		}

		for { //compare suffix
			pre := byte('*')
			for i := range ps {
				l := len(ps[i])
				if l == 0 {continue}

				l--
				if ps[i][l] != '*' {
					if pre == '*' {
						pre = ps[i][l]
					} else if pre != ps[i][l] { //suffix not match 
						goto bad
					}
					ps[i] = ps[i][:l]
				}
			}
			if pre == '*' {break}//no more suffix needs matching

			suffix = append(suffix, pre)
		}

		for _, p := range ps { //since prefix and suffix is matched, we can have any characters between them
			for i := range p {
				if p[i] != '*' {
					prefix = append(prefix, p[i])
				}
			}
		}

		for i := len(suffix) - 1; i >= 0; i-- {prefix = append(prefix, suffix[i])}

		printf("Case #%d: %s\n", t, string(prefix))
		continue
	bad:
		printf("Case #%d: *\n", t)
	}
}
