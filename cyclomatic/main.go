package main

import "fmt"

func main() {
	a, b, c, d := true, true, false, false
	if a {
		fmt.Println("a")
	}
	if b {
		fmt.Println("b")
	}
	if c {
		fmt.Println("c")
	}
	if d {
		fmt.Println("d")
	}
}
