package http

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	httptransport "sample/internal/kit/transport/http"

	"net/http"
	"pb"
)

func JSONRequestDecoder[T any](_ context.Context, r *http.Request) (*T, error) {
	var req T

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil && errors.Is(err, io.EOF) {
		return nil, err
	}

	return &req, nil
}

func JSONResponseEncoder(ctx context.Context, response interface{}) error {
	return nil
}

func (transport *HTTPTransport) HealthCheck() http.Handler {
	return httptransport.NewServer(
		transport.service.HealthCheck,
		JSONRequestDecoder[pb.HealthCheckRequest],
		JSONResponseEncoder,
	)
}
