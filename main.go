package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	log.Print("Server v1 started on " + port)
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(port, nil))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
