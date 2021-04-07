package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//var f, _ = os.Open("test.txt")
//var reader *bufio.Reader = bufio.NewReader(f)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string) interface{} {
	var s string
	for {
		fmt.Fscanf(reader, "%s", &s)
		if s != "" {
			switch f {
			case "d":
				d, _ := strconv.Atoi(s)
				return d
			case "s":
				return s
			}
		}
	}
}

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	T := scanf("d").(int)

	for t := 1; t <= T; t++ {
		n := scanf("d").(int)
		if n == -1 {return}
		
		sold := make([]bool, n)
		showedCnt := make([]int, n)

		for i := 0; i < n; i++ {
			chosen, minCnt := -1, 200

			d := scanf("d").(int)
			if d == -1 {return}

			for i := 0; i < d; i++ {
				id := scanf("d").(int)
				if id == -1 {return}
				
				if showedCnt[id]++; !sold[id] && minCnt > showedCnt[id] {
					minCnt = showedCnt[id]
					chosen = id
				}
			}

			if chosen != -1 {
				sold[chosen] = true
			}

			fmt.Println(chosen)
		}
	}
}