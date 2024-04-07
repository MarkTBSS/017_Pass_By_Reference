package main

import "fmt"

func plus2(n *int) {
	*n += 2
}

func main() {
	a := 8
	fmt.Println(&a)
	plus2(&a)
	fmt.Println(a)
}
