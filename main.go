package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.Handle("/assets/", http.FileServer(http.Dir(".")))
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Sever running at http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
}
