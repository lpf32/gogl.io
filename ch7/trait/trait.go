package trait

import "fmt"

type Print interface {
	Print()
}

type A struct {
}

func (a A) Print() {
	fmt.Println("A")
}

type B struct {
}

func (b B) Print() {
	fmt.Println("B")
}

func Hello(o Print) {
	o.Print()
}
