package main

import (
	"fmt"
	"net/http"
	"log"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type Database map[string]dollars

func (db Database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}

}

func main()  {
	db := Database{"shoes": 50, "socks": 5}

	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
