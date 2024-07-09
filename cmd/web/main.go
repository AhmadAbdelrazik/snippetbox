package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errLogger  *log.Logger
	infoLogger *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLogger: infoLog,
		errLogger: errorLog,
	}

	flag.Parse()

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("POST /snippet/create", app.snippetCreate)

	srv := &http.Server{
		Addr:    *addr,
		Handler: mux,
	}

	infoLog.Printf("Stating server on %s", *addr)

	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
