package handlers

import (
	"context"

	"go.uber.org/zap"

	"memcache/internal/command"
	memcache "memcache/proto"
)

type MemcacheHandler struct {
	memcacheServer *command.Server
	logger         *zap.Logger

	memcache.UnimplementedCacheServer
}

func NewMemcacheHandler(memcacheServer *command.Server, logger *zap.Logger) *MemcacheHandler {
	return &MemcacheHandler{
		logger:         logger.Named("memcache"),
		memcacheServer: memcacheServer,
	}
}

func (s *MemcacheHandler) Get(ctx context.Context, request *memcache.GetRequest) (*memcache.GetResponse, error) {
	data, hasValue, err := s.memcacheServer.Get(ctx, request.GetKey())
	if err != nil {
		s.logger.Error("failed to GetData", zap.Error(err))
		return nil, err
	}

	return &memcache.GetResponse{
		Value: data,
		Has:   hasValue,
	}, nil
}

func (s *MemcacheHandler) Set(ctx context.Context, request *memcache.SetRequest) (*memcache.EmptyResponse, error) {
	if _, err := s.memcacheServer.Set(ctx, request.GetKey(), request.GetValue()); err != nil {
		s.logger.Error("failed to SaveData", zap.Error(err))
		return nil, err
	}

	return &memcache.EmptyResponse{}, nil
}

func (s *MemcacheHandler) Delete(ctx context.Context, request *memcache.DeleteRequest) (*memcache.EmptyResponse, error) {
	if _, err := s.memcacheServer.Delete(ctx, request.GetKey()); err != nil {
		s.logger.Error("failed to delete", zap.Error(err))
		return nil, err
	}

	return &memcache.EmptyResponse{}, nil
}
