package main

import (
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("heavy process starts")
	time.Sleep(5 * time.Second)
	log.Println("done")

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("hello\n"))
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8000", nil)
}
