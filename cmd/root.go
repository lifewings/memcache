package cmd

import (
	"log"
	"net"

	"github.com/vrischmann/envconfig"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"memcache/internal/command"
	handlers "memcache/internal/handler"
	"memcache/internal/storage"
	memcache "memcache/proto"
)

type Server struct {
	logger *zap.Logger
}

var conf struct {
	Address string
}

func Execute() error {
	if err := envconfig.InitWithPrefix(&conf, "memcache"); err != nil {
		log.Panic(err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	defer logger.Sync()

	srv := command.New(storage.New(), logger)

	listen, err := net.Listen("tcp", conf.Address)
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}
	defer listen.Close()

	server := grpc.NewServer()
	memcache.RegisterCacheServer(server, handlers.NewMemcacheHandler(srv, logger))
	memcache.RegisterHealthCheckServer(server, handlers.NewHealthHandler())
	if err := server.Serve(listen); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}

	return nil
}
