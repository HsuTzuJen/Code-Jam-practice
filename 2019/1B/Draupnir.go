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
	var w int
	var cnt0, cnt int

	scanf("%d %d\n", &T, &w)
	for t := 1; t <= T; t++ {
		fmt.Println(56)
		scanf("%d\n", &cnt0)
		k1cnt := cnt0/(1<<56)
		cnt0 -= k1cnt*(1<<56)
		k2cnt := cnt0/(1<<28)
		cnt0 -= k2cnt*(1<<28)
		
		fmt.Println(224)
		scanf("%d\n", &cnt)
		k4cnt := cnt/(1<<56)
		cnt -= k4cnt*(1<<56)
		k5cnt := cnt/(1<<44)
		cnt -= k5cnt*(1<<44)
		k6cnt := cnt/(1<<37)
		cnt0 -= k4cnt*(1<<14)+k5cnt*(1<<11)+k6cnt*(1<<9)
		k3cnt := cnt0/(1<<18)
		
		fmt.Println(k1cnt, k2cnt, k3cnt, k4cnt, k5cnt, k6cnt)		
		res := 0
		scanf("%d\n", &res)
		if res == -1 {return}
		//printf("Case #%d: \n",t)
	}
}