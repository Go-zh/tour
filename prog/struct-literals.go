package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // 类型为 Vertex
	q = &Vertex{1, 2} // 类型为 *Vertex
	r = Vertex{X: 1}  // Y:0 被省略
	s = Vertex{}      // X:0 和 Y:0
)

func main() {
	fmt.Println(p, q, r, s)
}
