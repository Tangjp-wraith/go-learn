package main

import "net/http"

func sayhello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/hello", sayhello)
	http.ListenAndServe(":8080", nil)
}
