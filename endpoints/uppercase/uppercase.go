package uppercase

import (
	"awesomeProject/definitions"
	"context"
	"errors"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

func MakeUppercaseEndpoint(svc definitions.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(definitions.UppercaseRequest)
		return svc.Uppercase(ctx, req)
	}
}

func Uppercase(ctx context.Context, req definitions.UppercaseRequest) (definitions.UppercaseResponse, error) {
	if req.S == "" {
		return definitions.UppercaseResponse{}, ErrEmpty
	}
	return definitions.UppercaseResponse{V: strings.ToUpper(req.S)}, nil
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")
