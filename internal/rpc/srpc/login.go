package srpc

import (
	"context"
	"sisco/internal/cfg"
	"sisco/internal/ldapconn"
	"sisco/internal/rpc/pb"
)

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	var err error

	authToken, isAdminToken, err := dbConn.QuerySecretToken(ctx, in.GetUser(), in.GetPassword())

	_, err = ldapconn.New(&cfg.Config)
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		Token:        authToken,
		IsAdminToken: isAdminToken,
	}, err
}
