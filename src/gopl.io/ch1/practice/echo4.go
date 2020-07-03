package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep = "", " "
	for temp, arg := range os.Args {
		fmt.Println(temp, ":", arg)
		s += arg + sep
	}
	fmt.Println(s)
}
