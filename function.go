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
	"io/ioutil"
	"net/http"

	"github.com/afritzler/gardener-docs-search/pkg/types"
)

func Search(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		returnOK(w, r)
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		var request types.Request
		err = json.Unmarshal(body, &request)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		//responseType := request.ResponseType

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("StatusMethodNotAllowed"))
	}
}

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
