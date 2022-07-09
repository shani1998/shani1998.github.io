package main

import (
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcLogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func NewGRPCServer() *grpc.Server {
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())

	// Add middleware to grpc svc
	unaryInterceptorChain := []grpc.UnaryServerInterceptor{
		// logging middle interceptor for grpc logging
		grpcLogrus.UnaryServerInterceptor(logrusEntry),
		// prometheus' metrics interceptor for getting grpc metrics
		grpcPrometheus.UnaryServerInterceptor,
		// validation interceptor  for validating incoming requests
		grpcValidator.UnaryServerInterceptor(),
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(unaryInterceptorChain...)),
	)

	return grpcServer
}
