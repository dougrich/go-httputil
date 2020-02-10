package httputil

import (
	"net/http"
)

type ValidationOption func(http.ResponseWriter, *http.Request) bool

type ValidationSet []ValidationOption
