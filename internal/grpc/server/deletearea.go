package server

import (
	"context"

	"sisco/pb"
)

func (s *server) DeleteArea(ctx context.Context, in *pb.DeleteAreaRequest) (*pb.DeleteAreaReply, error) {
	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), true)
	if !tokenIsValid || err != nil {
		return &pb.DeleteAreaReply{}, err
	}

	err = dbConn.DeleteArea(ctx, in.GetArea())

	return &pb.DeleteAreaReply{}, err
}
