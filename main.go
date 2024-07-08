package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From Abdelrazik's Home\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Started listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}