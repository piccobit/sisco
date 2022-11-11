package server

import (
	"context"

	"sisco/pb"
)

func (s *server) RegisterArea(ctx context.Context, in *pb.RegisterAreaRequest) (*pb.RegisterAreaReply, error) {
	var err error

	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), true)
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
