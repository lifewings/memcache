package command

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

type Server struct {
	storage Storager
	logger  *zap.Logger
}

func New(storage Storager, logger *zap.Logger) *Server {
	return &Server{
		storage: storage,
		logger:  logger.Named("server"),
	}
}

func (s *Server) Set(ctx context.Context, key, data string) (string, error) {
	res, err := s.storage.Set(ctx, key, data)
	if err != nil {
		return "", fmt.Errorf("error set: %w", err)
	}
	return res, nil
}

func (s *Server) Get(ctx context.Context, key string) (string, bool, error) {
	data, err := s.storage.Get(ctx, key)
	if err != nil {
		if errors.Is(err, ErrorNotFound) {
			return "", false, nil
		}
		return "", false, fmt.Errorf("error get: %w", err)
	}

	return data, true, nil
}

func (s *Server) Delete(ctx context.Context, key string) (string, error) {
	res, err := s.storage.Delete(ctx, key)
	if err != nil {
		return "", fmt.Errorf("error delete: %w", err)
	}
	return res, nil
}
