package main

import (
	"log"
	"net/http"
)

const staticLocation = "./../static"

func main() {
	const port = ":8080"
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(staticLocation))

	mux.Handle("/", fs)
	mux.HandleFunc("/apiDefinition", apiDefinitionHandler)
	mux.HandleFunc("/conformance", conformanceHandler)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Printf("Serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}

func apiDefinitionHandler(w http.ResponseWriter, r *http.Request) {
	filePath := staticLocation + "/apiDefinition.html"
	http.ServeFile(w, r, filePath)
}

func conformanceHandler(w http.ResponseWriter, r *http.Request) {
	filePath := staticLocation + "/conformance.html"
	http.ServeFile(w, r, filePath)
}
