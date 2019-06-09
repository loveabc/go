package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32
type database map[string]dollars

func(d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func(db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			fmt.Fprintf(w, "There is no %v", item)
			return
		}
		fmt.Fprintf(w, "%s: %s\n", item, price)
	case "/add":
		item := req.URL.Query().Get("item")
		price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 64)
		if err != nil {
			fmt.Fprintf(w,"price error")
			return
		}
		db[item] = dollars(price)
		fmt.Fprintf(w,"add")
	default:
		fmt.Fprintf(w,"default")
	}
}

func main() {
	fmt.Println("server is up...")
	db := database{"shoes": 50, "socks": 100}
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", db))
}
