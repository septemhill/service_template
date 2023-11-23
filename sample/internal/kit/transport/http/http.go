package http

import (
	"net/http"
	"sample/internal/kit/endpoint"
)

type ServerOption[Request, Response any] func(*Server[Request, Response])

func ServerBefore[Request, Response any](before ...RequestFunc) ServerOption[Request, Response] {
	return func(s *Server[Request, Response]) {
		s.before = append(s.before, before...)
	}
}

func ServerAfter[Request, Response any](after ...ServerResponseFunc) ServerOption[Request, Response] {
	return func(s *Server[Request, Response]) {
		s.after = append(s.after, after...)
	}
}

type Server[Request, Response any] struct {
	decoder  DecodeRequestFunc[Request]
	encoder  EncodeResponseFunc
	before   []RequestFunc
	endpoint endpoint.Endpoint[Request, Response]
	after    []ServerResponseFunc
}

func NewServer[Request, Response any](
	e endpoint.Endpoint[Request, Response],
	dec DecodeRequestFunc[Request],
	enc EncodeResponseFunc,
	opts ...ServerOption[Request, Response],
) *Server[Request, Response] {
	return &Server[Request, Response]{
		endpoint: e,
		decoder:  dec,
		encoder:  enc,
	}
}

func (s *Server[Request, Response]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	request, err := s.decoder(ctx, r)
	if err != nil {
		// TODO: error handling
	}

	for _, fn := range s.before {
		ctx = fn(ctx, r)
	}

	response, err := s.endpoint(ctx, request)
	if err != nil {
		// TODO: error handling
	}

	for _, fn := range s.after {
		ctx = fn(ctx, w)
	}

	if err := s.encoder(ctx, response); err != nil {
		// TODO: error handling
	}
}
