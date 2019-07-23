package endpoints

import (
	"awesomeProject/definitions"
	"awesomeProject/endpoints/cookie"
	"awesomeProject/endpoints/count"
	"awesomeProject/endpoints/uppercase"
	"context"
)

func New() definitions.StringService {
	return &stringService{}
}

type stringService struct{}

func (s stringService) SetCookie(ctx context.Context, request definitions.CookieRequest) (definitions.CookieResponse, error) {
	return cookie.SetCookie(ctx, request)
}

func (s stringService) Uppercase(ctx context.Context, request definitions.UppercaseRequest) (definitions.UppercaseResponse, error) {
	return uppercase.Uppercase(ctx, request)
}

func (s stringService) Count(ctx context.Context, request definitions.CountRequest) (definitions.CountResponse, error) {
	return count.Count(ctx, request)
}
