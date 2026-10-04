package main

import "fmt"

// Задание 1
func main() {
	var (
		a int
		b int8
		c int16
		d int32
		e int64
		f uint
		g uint16
		h uint32
		i uint64
		j float32
		k float64
		l bool
		m rune
		n byte
		o string
		p complex64
		q complex128
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)

	a = -10
	b = 10
	c = 100
	d = 1000
	e = 10000
	f = 20
	g = 200
	h = 2000
	i = 20000
	j = 2.3
	k = 23.3
	l = true
	m = 'Ы'
	n = 'A'
	o = "Hello World!"
	p = 3 + 4i
	q = 4i
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
}
