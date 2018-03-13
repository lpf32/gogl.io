package main

import (
	"net"
	"log"
	"io"
	"os"
	"fmt"
	"strings"
)

type clockzone struct {
	id int
	location string
	con net.Conn
}

func main()  {
	for i, arg := range os.Args[1:] {
		args := strings.Split(arg, "=")

		conn, err := net.Dial("tcp", args[1])
		if err != nil {
			log.Fatal(err)
		}
		clock := clockzone{i, args[0], conn}
		go mustCopy(os.Stdout, clock)
	}
}

func mustCopy(dst io.Writer, src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
