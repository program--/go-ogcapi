package main

import (
	"log"
	"net/http"
)

const staticLocation = "./../static"

func main() {
	const port = ":8080"
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", indexHandler)
	mux.HandleFunc("GET /api-definition", apiDefinitionHandler)
	mux.HandleFunc("GET /conformance", conformanceHandler)
	mux.HandleFunc("GET /collections", collectionsHandler)
	mux.HandleFunc("GET /collections/{collectionId}", collectionsIdHandler)
	mux.HandleFunc("GET /collections/{collectionId}/items", collectionIdItemHandler)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Printf("Serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("index")
	filePath := staticLocation + "/index.html"
	http.ServeFile(w, r, filePath)
}

func apiDefinitionHandler(w http.ResponseWriter, r *http.Request) {
	filePath := staticLocation + "/apiDefinition.html"
	http.ServeFile(w, r, filePath)
}

func conformanceHandler(w http.ResponseWriter, r *http.Request) {
	filePath := staticLocation + "/conformance.html"
	http.ServeFile(w, r, filePath)
}

func collectionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("collections"))

}

func collectionsIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("collectionId")
	log.Println("id handler")
	w.Write([]byte("collectionId: " + id))
}

func collectionIdItemHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("item")
	id := r.PathValue("collectionId")
	w.Write([]byte("items for collection: " + id))
}

func collectionIdItemFeature(w http.ResponseWriter, r *http.Request) {
	cId := r.PathValue("collectionId")
	fId := r.PathValue("featureId")
	w.Write([]byte("collection: " + cId + " feature: " + fId))
}
