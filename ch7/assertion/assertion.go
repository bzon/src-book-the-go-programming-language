package main

import "fmt"

type Awesome interface {
	Print() string
}

type Word string
type Number string
type MyStruct struct {
	Message string
}

func (w Word) Print() string     { return fmt.Sprint(w) }
func (n Number) Print() string   { return fmt.Sprint(n) }
func (m MyStruct) Print() string { return fmt.Sprint("a struct") }

func main() {
	var w Word = "hi"
	printInterfaceType(w)
	printInterfaceType(MyStruct{"helloooo struct"})
}

func printInterfaceType(a Awesome) {
	switch a.(type) {
	case Word:
		fmt.Println("type is Word")
	case Number:
		fmt.Println("type is Number")
	case MyStruct:
		fmt.Println("type is MyStruct")
		m := a.(MyStruct)
		fmt.Println(m.Message)
	}
}
