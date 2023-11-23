package http

import (
	"context"
	"net/http"
)

type DecodeRequestFunc[Request any] func(context.Context, *http.Request) (*Request, error)

type EncodeResponseFunc func(context.Context, interface{}) error
