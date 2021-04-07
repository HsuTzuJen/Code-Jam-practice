package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

//var f, _ = os.Open("test.txt")
//var reader *bufio.Reader = bufio.NewReader(f)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

var T int
var n, l int

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var x int

	scanf("%d\n", &T)
	for t := 1; t <= T; t++ {
		scanf("%d %d\n", &n, &l)

		nleft := n
		totalp := 0
		m := minheap{}

		for i := 0; i < l; i++ {
		    if i == l-1{
		        scanf("%d\n", &x)
		    }else{
		        scanf("%d ", &x)   
		    }
			nleft -= x
			totalp += getp(x)
			heap.Push(&m, x)
		}

		for i := l; i < n; i++ {
			heap.Push(&m, 0)
		}

		for nleft > 0 {
			x = heap.Pop(&m).(int)
			totalp += getp(x+1) - getp(x)
			heap.Push(&m, x+1)
			nleft--
		}

		printf("Case #%d: %d\n", t, totalp)
	}
}

func getp(x int) int { //use n<<1 as 1 uint
	return (200*x + n) / (n << 1)
}

type minheap []int

func (m minheap) Len() int              { return len(m) }
func (m minheap) Less(i, j int) bool    { return (200*m[i]+n)%(n<<1) > (200*m[j]+n)%(n<<1) }
func (m minheap) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *minheap) Push(x interface{})   { *m = append(*m, x.(int)) }
func (m *minheap) Pop() (x interface{}) { x, *m = (*m)[len(*m)-1], (*m)[:len(*m)-1]; return }