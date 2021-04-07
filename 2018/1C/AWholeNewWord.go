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
	var n, l int

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d %d\n", &n, &l)
		showed := make(map[string]struct{}, 2000)
		canuse := make([]map[rune]struct{}, l, l)
		for i := range canuse {
			canuse[i] = map[rune]struct{}{}
		}
		for i := 0; i < n; i++ {
			var str string
			scanf("%s\n", &str)
			showed[str] = struct{}{}
			for j, ch := range str {
				canuse[j][ch] = struct{}{}
			}
		}

		dp := make([]map[string]struct{}, l)
		for i := range dp {
			dp[i] = map[string]struct{}{}
		}

		var dfs func(i int, buf []rune) []rune
		dfs = func(i int, buf []rune) []rune {
			if i == l {
				if _, ok := showed[string(buf)]; ok {
					return nil
				}
				return buf
			}

			if _, ok := dp[i][string(buf)]; ok {
				return nil
			}

			for ch := range canuse[i] {
				if tmp := dfs(i+1, append(buf, ch)); tmp != nil {
					return tmp
				}
			}

			dp[i][string(buf)] = struct{}{}

			return nil
		}

		var ans string
		if res := dfs(0, make([]rune, 0, l)); res == nil {
			ans = "-"
		} else {
			ans = string(res)
		}

		printf("Case #%d: %s\n", t, ans)
	}
}