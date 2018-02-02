package main

import (
	"fmt"
	"strings"
)

var res string

func match(a []string, b []string) int {
	bLen := len(b)

	for i:=0; i < bLen; i++ {
		if a[i] != b[i] {
			return i
		}
	}
	return 0
}

func main()  {
	a := "/a/b/c.txt"
	b := "/a/x/c/d.txt"


	aA := strings.Split(a, "/")
	bA := strings.Split(b, "/")

	index := match(aA[:len(aA)-1], bA[:len(bA)-1])

	for i := index; i < len(bA)-1; i++ {
		res += "../"
	}

	for i := index; i < len(aA)-1; i++ {
		res += aA[i] + "/"
	}

	res += aA[len(aA)-1]

	fmt.Println(res)

}
