package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {

	var s string
	var i int
	var arr [3]int

	fmt.Println(s)
	fmt.Println(i)
	fmt.Println(arr)

	//anim := gif.GIF{LoopCount: nframes}
	freq := rand.Float64() * 3.0
	t := 0.0

	//fmt.Println(anim)
	fmt.Println(freq)
	fmt.Println(t)

	in, err := os.Open("123")
	//..
	//in,err := os.Open("1234")  编译错误
	in, err1 := os.Open("1234")
	fmt.Println(in)
	fmt.Println(err)
	fmt.Println(err1)

}
