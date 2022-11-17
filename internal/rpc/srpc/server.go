package srpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"sisco/internal/db"
	"sisco/internal/rpc/pb"
)

type server struct {
	pb.UnimplementedLoginServer
	pb.UnimplementedRegisterAreaServer
	pb.UnimplementedRegisterServiceServer
	pb.UnimplementedDeleteAreaServer
	pb.UnimplementedDeleteServiceServer
	pb.UnimplementedListServiceInAreaServer
	pb.UnimplementedListServicesServer
	pb.UnimplementedListAreasServer
	pb.UnimplementedListTagsServer
}

type Server struct {
	grpcServer  *grpc.Server
	listenAddr  string
	useTLS      bool
	tlsCertFile string
	tlsKeyFile  string
}

var (
	dbConn *db.Client
)

func New(opts ...func(*Server)) (*Server, error) {
	newServer := Server{}

	for _, s := range opts {
		s(&newServer)
	}

	if newServer.useTLS {
		creds, err := credentials.NewServerTLSFromFile(newServer.tlsCertFile, newServer.tlsKeyFile)
		if err != nil {
			log.Fatalf("failed to setup TLS: %v", err)
		}

		newServer.grpcServer = grpc.NewServer(grpc.Creds(creds))
	} else {
		newServer.grpcServer = grpc.NewServer()
	}

	return &newServer, nil
}

// Run implements the gRPC server.
func (s *Server) Run() {
	var err error

	dbConn, err = db.New()
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer dbConn.Close()

	lis, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterLoginServer(s.grpcServer, &server{})
	pb.RegisterRegisterAreaServer(s.grpcServer, &server{})
	pb.RegisterRegisterServiceServer(s.grpcServer, &server{})
	pb.RegisterDeleteAreaServer(s.grpcServer, &server{})
	pb.RegisterDeleteServiceServer(s.grpcServer, &server{})
	pb.RegisterListServiceInAreaServer(s.grpcServer, &server{})
	pb.RegisterListServicesServer(s.grpcServer, &server{})
	pb.RegisterListAreasServer(s.grpcServer, &server{})
	pb.RegisterListTagsServer(s.grpcServer, &server{})

	log.Printf("gRPC server listening at %v", lis.Addr())

	if err := s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server failed to start: %v", err)
	}
}

func (s *Server) GracefulStop() {
	s.grpcServer.GracefulStop()
}

func ListenAddr(listenAddr string) func(*Server) {
	return func(server *Server) {
		server.listenAddr = listenAddr
	}
}

func UseTLS(yesOrNo bool) func(*Server) {
	return func(server *Server) {
		server.useTLS = yesOrNo
	}
}

func TLSCertFile(tlsCertFile string) func(*Server) {
	return func(server *Server) {
		server.tlsCertFile = tlsCertFile
	}
}

func TLSKeyFile(tlsKeyFile string) func(*Server) {
	return func(server *Server) {
		server.tlsKeyFile = tlsKeyFile
	}
}
