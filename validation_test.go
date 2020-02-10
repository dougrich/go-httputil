package httputil

import (
	"testing"
	"net/http"
	htu "github.com/dougrich/httptestutil"
)



func CreateHandler(options ValidationSet) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		if !options.IsValidRequest(w, r) {
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func TestValidateMethod(t *testing.T) {
	htu.TestSet{
		htu.Test(
			"RejectGET405",
			htu.RequestMethod(http.MethodGet),
			htu.ResponseStatus(http.StatusMethodNotAllowed),
		),
		htu.Test(
			"AcceptPOST200",
			htu.RequestMethod(http.MethodPost),
			htu.ResponseStatus(http.StatusOK),
		),
	}.Run(t, CreateHandler(ValidationSet{
		ValidateMethod(http.MethodPost),
	}))
}

func TestValidateContentType(t *testing.T) {
	htu.TestSet{
		htu.Test(
			"RejectGarbageType",
			htu.RequestHeader(HeaderContentType, "application/garbage"),
			htu.ResponseStatus(http.StatusUnsupportedMediaType),
		),
		htu.Test(
			"AcceptFormType",
			htu.RequestHeader(HeaderContentType, MimeTypeUrlEncodedForm),
			htu.ResponseStatus(http.StatusOK),
		),
	}.Run(t, CreateHandler(ValidationSet{
		ValidateContentType(MimeTypeUrlEncodedForm),
	}))
}