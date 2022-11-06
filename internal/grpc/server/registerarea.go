package server

import (
	"context"

	"sisco/pb"
)

func (s *server) RegisterArea(ctx context.Context, in *pb.RegisterAreaRequest) (*pb.RegisterAreaReply, error) {
	var err error

	success := true

	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), true)
	if !tokenIsValid || err != nil {
		success = false
	}

	return &pb.RegisterAreaReply{
		Success: success,
	}, err
}
