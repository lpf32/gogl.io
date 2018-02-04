package main

import (
	"net/http"
	"fmt"
)

func main()  {
	res, err := http.Head("https://www.baidu.com")
	if err != nil {
		fmt.Println("http head error", err)
	}

	fmt.Println(res)
}
