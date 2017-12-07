package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world! This is Ap2 container.")
}

func hogeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hogehoge.")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hoge", hogeHandler)
	http.ListenAndServe(":3000", nil)
}
