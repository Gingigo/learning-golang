package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println(p)
	*p = 2
	fmt.Println(x)

	var z, y int
	fmt.Println(&z == &z, &z == &y, &z == nil)
	fmt.Println(&z, &y)

	var q = f()
	fmt.Println(q)

	v := 1
	fmt.Println(&v)
	k := incr(&v)
	fmt.Println(v)
	fmt.Println(&k)

}
func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
