package main

import (
	"fmt"
	"github.com/afritzler/gardener-docs-search"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Gardener Doc Search")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	log.Printf("starting to serve on http://%v:%v...\n", name, port)
	registerHandlers()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}


func registerHandlers() {
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/", returnOK)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	function.Search(w, r)
}

func returnOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}