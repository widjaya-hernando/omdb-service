package server

import (

	// holdersv1 "tutorial/gen/go/proto/holders"

	"flag"
	"log"
	"net"

	pb "github.com/widjaya-hernando/omdb-service/gen"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedOMDBServiceServer
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}

func (s *Server) Run() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gSer := grpc.NewServer()
	pb.RegisterOMDBServiceServer(gSer, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := gSer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
