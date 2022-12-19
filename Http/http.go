package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hei", helloHandler)
	err := http.ListenAndServe("localhost:2020", nil)
	log.Fatal(err)
}

func helloHandler(a http.ResponseWriter, b *http.Request) {
	a.Write([]byte("hey"))
}
