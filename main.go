package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Sheriff-Hoti/go_std_api/dto"
	"github.com/Sheriff-Hoti/go_std_api/middleware"
)

func main() {
	fmt.Println("hello")

	router := http.NewServeMux()

	router.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {

		var req dto.Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		first, err1 := strconv.ParseFloat(req.First, 64)
		second, err2 := strconv.ParseFloat(req.Second, 64)

		if err1 != nil {

			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}

		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusBadRequest)
			return

		}

		w.Write([]byte("The response" + strconv.FormatFloat(first+second, 'f', 2, 64)))
	})

	router.HandleFunc("POST /multiply", func(w http.ResponseWriter, r *http.Request) {
		var req dto.Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		first, err1 := strconv.ParseFloat(req.First, 64)
		second, err2 := strconv.ParseFloat(req.Second, 64)

		if err1 != nil {

			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}

		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusBadRequest)
			return

		}

		w.Write([]byte("The response" + strconv.FormatFloat(first*second, 'f', 2, 64)))
	})

	router.HandleFunc("POST /divide", func(w http.ResponseWriter, r *http.Request) {
		var req dto.Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		first, err1 := strconv.ParseFloat(req.First, 64)
		second, err2 := strconv.ParseFloat(req.Second, 64)

		if err1 != nil {

			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}

		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusBadRequest)
			return

		}

		w.Write([]byte("The response" + strconv.FormatFloat(first/second, 'f', 2, 64)))
	})

	router.HandleFunc("POST /subtract", func(w http.ResponseWriter, r *http.Request) {
		var req dto.Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		first, err1 := strconv.ParseFloat(req.First, 64)
		second, err2 := strconv.ParseFloat(req.Second, 64)

		if err1 != nil {

			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}

		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusBadRequest)
			return

		}

		if second == 0 {
			http.Error(w, "The second value must not be empty", http.StatusBadRequest)
			return
		}

		w.Write([]byte("The response" + strconv.FormatFloat(first/second, 'f', 2, 64)))
	})

	stack := middleware.CreateStack(middleware.Logging)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	log.Println("Started server on port 8080")

	err := server.ListenAndServe()

	log.Fatalf("there was an error: %v", err.Error())
}
