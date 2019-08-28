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
		fmt.Printf("%v", r.Header)
		decoder := json.NewDecoder(r.Body)
		var n Numbers
		err := decoder.Decode(&n)
		if err != nil {
			panic(err)
		}

		resp := &Response{
			Product: n.A * n.B,
		}

		w.Header().Set("Content-Type", "application/json")
		otHeaders := []string{"x-request-id", "x-b3-traceid", "x-b3-spanid", "x-b3-parentspanid", "x-b3-sampled", "x-b3-flags", "x-ot-span-context", "User-Agent"}
		for _, header := range otHeaders {
			value, present := r.Header[header]
			if present {
				fmt.Printf("Setting %s value as %s", header, value[0])
				w.Header().Set(header, value[0])
			}
		}

		json.NewEncoder(w).Encode(resp)
	}).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
