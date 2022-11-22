package srpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sisco/internal/cfg"
	"sisco/internal/ldapconn"
	"sisco/internal/rpc/pb"
)

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	authToken, permissions, err := dbConn.QueryAuthToken(ctx, in.GetUser(), in.GetPassword())
	if err != nil {
		return &pb.LoginReply{}, status.Error(codes.PermissionDenied, err.Error())
	}

	_, err = ldapconn.New(&cfg.Config)
	if err != nil {
		return &pb.LoginReply{}, status.Error(codes.Aborted, err.Error())
	}

	return &pb.LoginReply{
		Token:       authToken,
		Permissions: uint64(permissions),
	}, err
}
