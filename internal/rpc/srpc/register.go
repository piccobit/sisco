package srpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sisco/internal/auth"
	"sisco/internal/rpc/pb"
)

func (s *server) RegisterArea(ctx context.Context, in *pb.RegisterAreaRequest) (*pb.RegisterAreaReply, error) {
	var err error

	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), auth.Admin)
	if !tokenIsValid || err != nil {
		return &pb.RegisterAreaReply{}, status.Error(codes.PermissionDenied, err.Error())
	}

	err = dbConn.CreateArea(ctx, in.GetArea(), in.GetDescription())
	if err != nil {
		return &pb.RegisterAreaReply{}, status.Error(codes.Aborted, err.Error())
	}

	return &pb.RegisterAreaReply{}, nil
}

func (s *server) RegisterService(ctx context.Context, in *pb.RegisterServiceRequest) (*pb.RegisterServiceReply, error) {
	var err error

	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), auth.Admin|auth.Service)
	if !tokenIsValid || err != nil {
		return &pb.RegisterServiceReply{}, status.Error(codes.PermissionDenied, err.Error())
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
		return &pb.RegisterServiceReply{}, status.Error(codes.Aborted, err.Error())
	}

	return &pb.RegisterServiceReply{}, nil
}
