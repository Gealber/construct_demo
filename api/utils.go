package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Gealber/construct_demo/accounts"
	cjson "github.com/Gealber/construct_demo/serializer/json"
)

func jsonResponse(response interface{}, w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func decodeUserBodyReq(r *http.Request) *accounts.User {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Unable to read request body\n")
		return nil
	}

	user, err := cjson.Decode(data)
	if err != nil {
		return nil
	}
	return user
}
