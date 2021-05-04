package main

import (
	"fmt"
	"log"
	"net/http"
)

var Port string

func main() {
	port := Port
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "hello from b-nova")
		if err != nil {
			log.Fatalf("error printing message to response writer %s", err)
			return
		}
	})

	log.Printf("Now listening on port %s.", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

