package count

import (
	"awesomeProject/definitions"
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeCountEndpoint(svc definitions.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(definitions.CountRequest)
		return svc.Count(ctx, req)

	}
}

func Count(ctx context.Context, req definitions.CountRequest) (definitions.CountResponse, error) {
	return definitions.CountResponse{V: len(req.S)}, nil
}
