package main

import (
	"ex05/queue"
	"fmt"
)

func main() {
	q := queue.Queue{}
	q.Push(3)
	q.Push("yongckim")
	fmt.Printf("%#v", q)
}
