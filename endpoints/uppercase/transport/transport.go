package transport

import (
	"awesomeProject/definitions"
	"context"
	"encoding/json"
	"net/http"
)

func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request definitions.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
