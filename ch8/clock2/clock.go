package main

import (
	"net"
	"log"
	"io"
	"time"
	"os"
	"flag"
)

var port = flag.String("port", "8000", "port number")

func main()  {
	flag.Parse()

	zone := os.Getenv("TZ")
	if zone == "" {
		zone = "Asia/Shanghai"
	}
	loc, err := time.LoadLocation(zone)
	listener, err := net.Listen("tcp", "localhost:" + *port)

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, loc)
	}
}

func handleConn(conn net.Conn, loc *time.Location)  {
	defer conn.Close()
	for {

		_, err := io.WriteString(conn, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
