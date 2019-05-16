package main

import (
	"fmt"
	"os"

	"github.com/blevesearch/bleve"
)

func main() {
	var index bleve.Index

	data := struct {
		Name string
	}{
		Name: "text",
	}

	// open a new index
	mapping := bleve.NewIndexMapping()
	if _, err := os.Stat("example.bleve"); os.IsNotExist(err) {
		index, err := bleve.New("example.bleve", mapping)
		if err != nil {
			fmt.Println(err)
			return
		}
		// index some data
		index.Index("id", data)
	} else {
		index, _ = bleve.Open("example.bleve")
	}

	// search for some text
	query := bleve.NewMatchQuery("text")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
	for _, s := range searchResults.Hits {
		fmt.Printf("%s\n", s)
	}
}