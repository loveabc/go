package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32
type database map[string]dollars

func main() {
	fmt.Println("server is up...")
	db := database{"shoes": 50, "socks": 100}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.Handle("/add", http.HandlerFunc(db.add))

	log.Fatal(http.ListenAndServe("127.0.0.1:8200", mux))
}

func(db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func(db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		fmt.Fprintf(w, "There is no %v", item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func(db database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 64)
	if err != nil {
		fmt.Fprintf(w,"price error")
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w,"add")
}
