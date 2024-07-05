package main

import (
	"log"
	"net/http"
)

func main() {
	const port = ":8080"
	log.Printf("Serving on port: %s\n", port)

	// consider switching from default mux when we want to add middleware
	log.Fatal(http.ListenAndServe(port, nil))
}
