package handlers

import (
	"context"

	memcache "memcache/proto"
)

type HealthHandler struct {
	memcache.UnimplementedHealthCheckServer
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (s *HealthHandler) Check(_ context.Context, _ *memcache.EmptyRequest) (*memcache.HealthResponse, error) {
	return &memcache.HealthResponse{Status: memcache.HealthResponse_SUCCESS}, nil
}
