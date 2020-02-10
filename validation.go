package httputil

import (
	"net/http"
)

/*
This actually does the checking; loops through every validation option and calls it
*/
func (v ValidationSet) IsValidRequest(
	w http.ResponseWriter,
	r *http.Request,
) (
	bool,
) {
	for _, o := range v {
		if !o(w, r) {
			return false
		}
	}
	return true
}

// Checks to make sure that the method used is one of the ones passed in; otherwise returns a 405
func ValidateMethod(allowed ...string) ValidationOption {
	return func (w http.ResponseWriter, r *http.Request) bool {
		for _, method := range allowed {
			if r.Method == method {
				return true
			}
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
}

func ValidateContentType(allowed ...string) ValidationOption {
	return func (w http.ResponseWriter, r *http.Request) bool {
		types := r.Header[HeaderContentType]
		for _, tactual := range types {
			for _, texpected := range allowed {
				if tactual == texpected {
					return true
				}
			}
		}
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return false
	}
}