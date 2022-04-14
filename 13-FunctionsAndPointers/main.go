package main

import "fmt"

func show(n int) {
	fmt.Printf("the no is %d\n", n)
}

// func incr(n int) int {
// 	n++
// 	return n
// }

func incrPtr(n *int) {
	*n += 10
}

func multiply(i, j int) (res int) {
	res = i * j
	return
}

func main() {
	i := 5
	show(i)
	incrPtr(&i)
	show(i)
	res := multiply(12, 4)
	show(res)
}
