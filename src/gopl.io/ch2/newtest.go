package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	//一样的值，不同的地址
	x := newInt1()
	y := newInt2()

	fmt.Println(*x)
	fmt.Println(*y)
	fmt.Println(x == y)
}

func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dumy int
	return &dumy
}
