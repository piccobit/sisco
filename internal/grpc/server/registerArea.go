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
