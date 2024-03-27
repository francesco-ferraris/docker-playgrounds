package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	log.Println("Greeting", name)

	_, err := fmt.Fprintf(w, "Hello, %s!", name)
	if err != nil {
		w.WriteHeader(500)
	}
}
