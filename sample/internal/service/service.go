package service

import (
	"context"
	"pb"
)

type Service struct{}

func (srv *Service) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{}, nil
}

func NewService() *Service {
	return &Service{}
}
