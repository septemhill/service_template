package http

import (
	"context"
	"net/http"
)

type RequestFunc func(context.Context, *http.Request) context.Context

type ServerResponseFunc func(context.Context, http.ResponseWriter) context.Context
