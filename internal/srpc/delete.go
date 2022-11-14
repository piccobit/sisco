package srpc

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

func (s *server) DeleteService(ctx context.Context, in *pb.DeleteServiceRequest) (*pb.DeleteServiceReply, error) {
	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), true)
	if !tokenIsValid || err != nil {
		return &pb.DeleteServiceReply{}, err
	}

	err = dbConn.DeleteService(ctx, in.GetService(), in.GetArea())

	return &pb.DeleteServiceReply{}, err
}
