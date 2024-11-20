package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello")
	router := http.NewServeMux()
	router.HandleFunc("/item/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		w.Write([]byte("recieved request for item: " + id))
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Started server on port 8080")
	err := server.ListenAndServe()

	log.Fatalf("there was an error: %v", err.Error())
}
