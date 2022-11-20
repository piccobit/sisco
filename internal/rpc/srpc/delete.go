package srpc

import (
	"context"
	"sisco/internal/auth"
	"sisco/internal/rpc/pb"
)

func (s *server) DeleteArea(ctx context.Context, in *pb.DeleteAreaRequest) (*pb.DeleteAreaReply, error) {
	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), auth.Admin)
	if !tokenIsValid || err != nil {
		return &pb.DeleteAreaReply{}, err
	}

	err = dbConn.DeleteArea(ctx, in.GetArea())

	return &pb.DeleteAreaReply{}, err
}

func (s *server) DeleteService(ctx context.Context, in *pb.DeleteServiceRequest) (*pb.DeleteServiceReply, error) {
	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), auth.Admin|auth.Service)
	if !tokenIsValid || err != nil {
		return &pb.DeleteServiceReply{}, err
	}

	err = dbConn.DeleteService(ctx, in.GetService(), in.GetArea())

	return &pb.DeleteServiceReply{}, err
}
