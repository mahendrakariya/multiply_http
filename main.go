package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Numbers struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Response struct {
	Product int `json:"product"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("%v", r)
		fmt.Println(r.Body)
		decoder := json.NewDecoder(r.Body)
		var n Numbers
		err := decoder.Decode(&n)
		if err != nil {
			panic(err)
		}

		resp := &Response{
			Product: n.A * n.B,
		}
		// implement here
		json.NewEncoder(w).Encode(resp)
	}).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}