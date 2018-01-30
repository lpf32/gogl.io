package main

import (
	"go-recipes/gopl.io/ch6/intset"
	//"fmt"
	//"net/url"
	"fmt"
	//"time"
)

type Rocket struct {}

func (r *Rocket) Launch()  {
	fmt.Println("hello")
}

func main()  {
	s := intset.IntSet{}

	s.Add(1535)
	s.Add(9)

	fmt.Println(s.String())

	y := intset.IntSet{}
	y.Add(1)
	y.Add(2)
	y.Add(9)
	fmt.Println(y.String())

	s.UnionWith(&y)
	fmt.Println(s.String())

	fmt.Println(s.Has(9), s.Has(1))

	fmt.Println(&s)
	fmt.Println(s.String())
	fmt.Println(s)
}
