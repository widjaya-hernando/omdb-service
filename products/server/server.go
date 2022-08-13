package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	// holdersv1 "tutorial/gen/go/proto/holders"

	pb "github.com/widjaya-hernando/prod-trans-user-service/gen"
	"github.com/widjaya-hernando/prod-trans-user-service/lib/database_transaction"
	"github.com/widjaya-hernando/prod-trans-user-service/products/repository"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

const (
	defaultPort = "60002"
)

type Server struct {
	db                 *gorm.DB
	repository         repository.RepoMethod
	transactionManager database_transaction.Client
}

func NewServer(ctx context.Context, db *gorm.DB, repository repository.RepoMethod, transactionManager database_transaction.Client) (*Server, error) {
	return &Server{
		db:                 db,
		repository:         repository,
		transactionManager: transactionManager,
	}, nil
}

func (s *Server) Run() {
	go func() {
		mux := runtime.NewServeMux()

		pb.RegisterProductsAPIHandlerServer(context.Background(), mux, s)

		log.Fatalln(http.ListenAndServe("0.0.0.0:8004", mux))
	}()

	port := "8001"
	if port == "" {
		port = defaultPort
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Print("net.Listen failed")
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterProductsAPIServer(grpcServer, s) // use authogenerated code to register the server
	reflection.Register(grpcServer)

	log.Printf("Starting Products server on port %s", port)
	go func() {
		grpcServer.Serve(listener)
	}()
}
