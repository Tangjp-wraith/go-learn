package main

import (
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get connection : %s", r.Proto)
	w.Write([]byte("hello, this is a http 2 message!"))
}

func main() {
	srv := &http.Server{
		Addr:    ":8088",
		Handler: http.HandlerFunc(handle),
	}
	log.Printf("Servering on https://0.0.0.0:8088")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}
