package http

import (
	"context"
	"fmt"
	"net/http"
	"pb"

	"github.com/gorilla/mux"
)

type Service interface {
	HealthCheck(context.Context, *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error)
}

type HTTPTransport struct {
	service Service
	server  *http.Server
}

func (transport *HTTPTransport) getRoutes(ctx context.Context) http.Handler {
	mux := mux.NewRouter()

	mux.Methods(http.MethodGet).Path("/healthz").Handler(transport.HealthCheck())

	return mux
}

func (transport *HTTPTransport) Start(ctx context.Context) error {
	transport.server.Handler = transport.getRoutes(ctx)

	go func() {
		_ = transport.server.ListenAndServe()
	}()

	return nil
}

func (transport *HTTPTransport) Stop(ctx context.Context) {
	_ = transport.server.Shutdown(ctx)
}

func NewHTTPTransport(service Service, port int) *HTTPTransport {
	return &HTTPTransport{
		service: service,
		server: &http.Server{
			Addr: fmt.Sprintf(":%d", port),
		},
	}
}
