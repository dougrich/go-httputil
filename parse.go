package httputil

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func ParseJson(
	w http.ResponseWriter,
	r *http.Request,
	o interface{},
) bool {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, ErrorBodyRead, http.StatusInternalServerError)
		return false
	}

	registration := Request{}
	if err := json.Unmarshal(b, &registration); err != nil {
		http.Error(w, ErrorBadJson, http.StatusBadRequest)
		return false
	}

	return true
}