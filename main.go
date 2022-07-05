package main

import "net/http"

func main() {
	http.HandleFunc("/route", hello)
	http.ListenAndServe(":9000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
