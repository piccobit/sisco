package srpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sisco/internal/auth"
	"sisco/internal/rpc/pb"
)

func (s *server) DeleteArea(ctx context.Context, in *pb.DeleteAreaRequest) (*pb.DeleteAreaReply, error) {
	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), auth.Admin)
	if !tokenIsValid || err != nil {
		return &pb.DeleteAreaReply{}, status.Error(codes.PermissionDenied, err.Error())
	}

	err = dbConn.DeleteArea(ctx, in.GetArea())
	if err != nil {
		return &pb.DeleteAreaReply{}, status.Error(codes.Aborted, err.Error())
	}

	return &pb.DeleteAreaReply{}, nil
}

func (s *server) DeleteService(ctx context.Context, in *pb.DeleteServiceRequest) (*pb.DeleteServiceReply, error) {
	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), auth.Admin|auth.Service)
	if !tokenIsValid || err != nil {
		return &pb.DeleteServiceReply{}, status.Error(codes.PermissionDenied, err.Error())
	}

	err = dbConn.DeleteService(ctx, in.GetService(), in.GetArea())
	if err != nil {
		return &pb.DeleteServiceReply{}, status.Error(codes.Aborted, err.Error())
	}

	return &pb.DeleteServiceReply{}, nil
}
