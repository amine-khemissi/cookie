package cookie

import (
	"awesomeProject/definitions"
	"context"
	"os"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/twinj/uuid"
)

func MakeCookieEndpoint(svc definitions.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(definitions.CookieRequest)
		return svc.SetCookie(ctx, req)

	}
}

func SetCookie(ctx context.Context, request definitions.CookieRequest) (definitions.CookieResponse, error) {

	sessionID := uuid.NewV4().String()
	options := []string{sessionID}
	for k, v := range request.Options {
		options = append(options, k+"="+v)
	}

	f, err := os.Create("/tmp/sessionToken-" + sessionID)
	if err != nil {
		return definitions.CookieResponse{}, err
	}
	defer f.Close()
	f.WriteString(strings.Join(options, "; "))
	return definitions.CookieResponse{
		Session: sessionID,
		Options: request.Options}, nil
}
