package main

import (
	"fmt"
	"net"
	"net/http"

	db "github.com/poc_grpc/db_connect"
	"github.com/poc_grpc/middleware"
	"github.com/poc_grpc/migrations"
	"github.com/poc_grpc/observability"
	proto "github.com/poc_grpc/pb"
	"github.com/poc_grpc/service"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	dbconnection := db.InitDB()
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	closer := observability.InitJaeger("Server grpc tracking")
	defer closer.Close()

	addr := fmt.Sprintf(":%d", 50051)
	addrHttp := fmt.Sprintf(":%d", 5000)

	log.Info().Str("app", "Server gRPC").Msgf("Starting Grpc server addr %s", addr)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal().Str("app", "Server gRPC").Err(err).Msg("failed to listen")
	}

	grpc_prometheus.EnableHandlingTimeHistogram()

	notebookService := service.NotebookService{}
	loginService := service.LoginService{}

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
	proto.RegisterLoginServer(grpcServer, loginService)

	http.Handle("/metrics", promhttp.Handler())

	go log.Fatal().Err(grpcServer.Serve(lis)).Msg("failed to start grpc")
	log.Fatal().Err(http.ListenAndServe(addrHttp, nil)).Msg("failed to start http server")

}
