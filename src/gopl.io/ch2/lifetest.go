package main

var global *int

func f1() {
	var x int
	x = 1
	global = &x
}

func g() {
	y := new(int)
	*y = 1
}

func main() {
	f1()
	g()
}
