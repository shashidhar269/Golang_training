package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hi", helloHandler)
	err := http.ListenAndServe("localhost:7070", nil)
	log.Fatal(err)
}

func helloHandler(a http.ResponseWriter, b *http.Request) {
	a.Write([]byte("hey"))
}
