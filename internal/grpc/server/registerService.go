package server

import (
	"context"

	"sisco/pb"
)

func (s *server) RegisterService(ctx context.Context, in *pb.RegisterServiceRequest) (*pb.RegisterServiceReply, error) {
	var err error

	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), true)
	if !tokenIsValid || err != nil {
		return &pb.RegisterServiceReply{}, err
	}

	err = dbConn.CreateService(ctx,
		in.GetService(),
		in.GetArea(),
		in.GetDescription(),
		in.GetProtocol(),
		in.GetHost(),
		in.GetPort(),
		in.GetTags(),
	)
	if err != nil {
		return &pb.RegisterServiceReply{}, err
	}

	return &pb.RegisterServiceReply{}, err
}
