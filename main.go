package main

import (
	"fmt"
	"net"

	db "github.com/poc_grpc/db_connect"

	"github.com/poc_grpc/middleware"
	proto "github.com/poc_grpc/pb"
	"github.com/poc_grpc/service"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

func main() {

	_ = db.InitDB()

	addr := fmt.Sprintf(":%d", 50051)

	log.Info().Str("app", "Server gRPC").Msgf("Starting Grpc server addr %s", addr)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal().Str("app", "Server gRPC").Err(err).Msg("failed to listen")
	}

	grpc_prometheus.EnableHandlingTimeHistogram()

	notebookService := service.NotebookService{}

	sOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			middleware.Interceptor(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_prometheus.StreamServerInterceptor,
		)),
	}

	grpcServer := grpc.NewServer(sOpts...)

	proto.RegisterNotebookServiceServer(grpcServer, notebookService)

	log.Fatal().Err(grpcServer.Serve(lis)).Msg("failed to start grpc")

}
