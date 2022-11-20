package srpc

import (
	"context"
	"sisco/internal/auth"
	"sisco/internal/rpc/pb"
)

func (s *server) RegisterArea(ctx context.Context, in *pb.RegisterAreaRequest) (*pb.RegisterAreaReply, error) {
	var err error

	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), auth.Admin)
	if !tokenIsValid || err != nil {
		return &pb.RegisterAreaReply{}, err
	}

	err = dbConn.CreateArea(ctx, in.GetArea(), in.GetDescription())
	if err != nil {
		return &pb.RegisterAreaReply{}, err
	}

	return &pb.RegisterAreaReply{}, err
}

func (s *server) RegisterService(ctx context.Context, in *pb.RegisterServiceRequest) (*pb.RegisterServiceReply, error) {
	var err error

	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), auth.Admin|auth.Service)
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
