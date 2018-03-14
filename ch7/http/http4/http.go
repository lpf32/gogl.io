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

func (db Database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db Database) price(w http.ResponseWriter, req *http.Request)  {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

func (f HandlerFunc) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func main()  {
	db := Database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
