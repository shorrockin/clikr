package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shorrockin/clikr/log"
)

func init() {
	Map("GET", "/series/{series}/season/{season}/episode/{episode}", JsonEncoder(RetrieveEpisodeInfo))
	Map("PUT", "/objects/{object_id}/click", JsonEncoder(RegisterObjectClick))
}

func Variables(request *http.Request) map[string]string {
	out := mux.Vars(request)

	user := request.Header.Get("X-Clikr-User")
	if user != "" {
		out["X-Clikr-User"] = user
	}

	return out
}

func JsonEncoder(next func(map[string]string) interface{}) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		result := next(Variables(request))
		bytes, error := json.Marshal(result)

		if error != nil {
			log.Warn("encountered error marshelling object for json: %v", error)
		}

		fmt.Fprint(writer, string(bytes))
	}
}
