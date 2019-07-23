package transport

import (
	"awesomeProject/definitions"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

func DecodeCookieRequest(_ context.Context, r *http.Request) (interface{}, error) {
	request := definitions.CookieRequest{
		Options: make(map[string]string),
	}

	for k, v := range r.URL.Query() {
		request.Options[k] = v[0]
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	options := []string{"sessionToken=" + response.(definitions.CookieResponse).Session}
	for k, v := range response.(definitions.CookieResponse).Options {
		options = append(options, k+"="+v)
	}
	w.Header().Set("Set-Cookie", strings.Join(options, "; "))
	return json.NewEncoder(w).Encode(response)
}
