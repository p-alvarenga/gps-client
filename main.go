package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintln(w, "Hello! Go HTTP server is working.")
	})

	server := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: mux,
	}

	log.Println("HTTP server listening on 127.0.0.1:8000")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
