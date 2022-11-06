package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"sisco/internal/db"
	"sisco/pb"
)

type server struct {
	pb.UnimplementedLoginServer
}

var (
	dbConn *db.Client
)

// Run implements the gRPC server.
func Run(s *grpc.Server, listenAddr string) {
	var err error

	dbConn, err = db.New()
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer dbConn.Close()

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterLoginServer(s, &server{})

	log.Printf("gRPC server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("gRPC server failed to start: %v", err)
	}
}
