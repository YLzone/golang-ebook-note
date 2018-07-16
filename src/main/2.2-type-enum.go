package main

import (
	"fmt"
)

const (
	x = iota
	y = iota
	z = iota
)

const v = iota

const (
	h, i, j = iota, iota, iota
)

const (
	a       = iota
	b       = "B"
	c       = 123
	d, e, f = iota, iota, iota
	g       = iota
)

func main() {
	fmt.Println(x, y, z)
	fmt.Println(v)
	fmt.Println(h, i, j)
	fmt.Println(a, b, c, d, e, f, g)
}
