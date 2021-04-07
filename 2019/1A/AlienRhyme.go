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

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var T, n int
	w := ""
	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d\n", &n)

		root := &trie{}
		for i := 0; i < n; i++ {
			scanf("%s\n", &w)
			node := root
			for i := len(w) - 1; i >= 0; i-- {
				ch := w[i] - 'A'
				if node.ch[ch] == nil {
					node.ch[ch] = &trie{}
				}
				node = node.ch[ch]
			}
			node.isword = true
		}

		ans := 0

		var dfs func(node *trie) int
		dfs = func(node *trie) int {
			cnt := 0
			if node.isword {
				cnt++
			}
			for _, ch := range node.ch {
				if ch != nil {
					cnt += dfs(ch)
				}
			}
			if cnt > 1 {
				ans += 2
				cnt -= 2
			}
			return cnt
		}

		for _, node := range root.ch {
			if node != nil {
				dfs(node)
			}
		}

		printf("Case #%d: %d\n", t, ans)
	}
}

type trie struct {
	ch     [26]*trie
	isword bool
}

