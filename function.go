// Copyright © 2019 Andreas Fritzler <andreas.fritzler@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package function

import (
	"encoding/json"

	"github.com/blevesearch/bleve"

	"io/ioutil"
	"log"
	"net/http"

	"github.com/afritzler/gardener-docs-search/pkg/types"
)

// Search is the main entry point in our Cloud Function
func Search(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		returnOK(w, r)
	case "POST":
		log.Println("got a POST request")
		var replies []interface{}
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			log.Printf("failed to parse body: %+v\n", err)
			replies = append(replies, generateTextMessage(types.RequestErrorMessage, 0))
			break
		}

		var request types.Request
		err = json.Unmarshal(body, &request)
		if err != nil {
			log.Printf("failed to parse request: %+v\n", err)
			replies = append(replies, generateTextMessage(types.RequestErrorMessage, 0))
			break
		}

		data, err := getSearchIndex(request.IndexJsonUrl)
		if err != nil {
			log.Printf("failed to retrieve search index: %+v\n", err)
			replies = append(replies, generateTextMessage(types.RequestErrorMessage, 0))
			break
		}
		log.Printf("got the data: found %d entries", len(data))

		log.Println("building up new index")
		mapping := bleve.NewIndexMapping()
		index, err := bleve.NewMemOnly(mapping)
		if err != nil {
			panic(err)
		}
		index.Index("content", data)

		query := bleve.NewQueryStringQuery(request.Query)
		searchRequest := bleve.NewSearchRequest(query)
		searchResult, _ := index.Search(searchRequest)

		defer index.Close()

		log.Printf("found %d hits", len(searchResult.Hits))
		log.Printf("result %+v hits for query %s", searchResult, request.Query)

		//for _, d := range data {
		//	fmt.Printf("%+v\n", d.Title)
		//}

		//for _, s := range searchResult.Hits {
		//	fmt.Printf("%+v\n", s.Fields)
		//}

		output, err := json.Marshal(types.Replies{Replies: replies})
		if err != nil {
			log.Printf("failed to marshal replies: %+v\n", err)
			replies = append(replies, generateTextMessage(types.RequestErrorMessage, 0))
			break
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("StatusMethodNotAllowed"))
	}
}

// helper function to generate a 200 response
func returnOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func generateTextMessage(text string, delay int) types.TextMessage {
	return types.TextMessage{
		Type:    types.ButtonsType,
		Content: text,
		Delay:   delay,
	}
}

func getSearchIndex(url string) (types.DataArray, error) {
	resp, httpErr := http.Get(url)
	if httpErr != nil {
		return types.DataArray{}, httpErr
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return types.DataArray{}, readErr
	}

	data := types.DataArray{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		return types.DataArray{}, err
	}
	defer resp.Body.Close()

	return data, nil
}
