package storage

import (
	"context"
	"sync"

	"memcache/internal/command"
)

const (
	DeleteSuccess = "delete"
)

type Storage struct {
	mu   sync.RWMutex
	data map[string]string
}

func New() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}

func (s *Storage) Get(_ context.Context, key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if val, found := s.data[key]; found {
		return val, nil
	}

	return "", command.ErrorNotFound
}

func (s *Storage) Set(_ context.Context, key, value string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
	return "", nil
}

func (s *Storage) Delete(_ context.Context, key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, found := s.data[key]; found {
		delete(s.data, key)
		return DeleteSuccess, nil
	}

	return "", command.ErrorNotFound
}
