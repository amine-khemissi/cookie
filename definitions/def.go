package definitions

import "context"

type StringService interface {
	Uppercase(ctx context.Context, request UppercaseRequest) (UppercaseResponse, error)
	Count(ctx context.Context, request CountRequest) (CountResponse, error)
	SetCookie(ctx context.Context, request CookieRequest) (CookieResponse, error)
}

type UppercaseRequest struct {
	S string `json:"s"`
}

type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type CountRequest struct {
	S string `json:"s"`
}

type CountResponse struct {
	V int `json:"v"`
}

type CookieRequest struct {
	Options map[string]string
}

type CookieResponse struct {
	Session string            `json:"session"`
	Options map[string]string `json:"options"`
}
