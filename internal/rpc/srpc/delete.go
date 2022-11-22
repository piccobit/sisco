package srpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sisco/internal/auth"
	"sisco/internal/rpc/pb"
	"strings"
)

func (s *server) DeleteArea(ctx context.Context, in *pb.DeleteAreaRequest) (*pb.DeleteAreaReply, error) {
	token, err := dbConn.QueryAuthTokenInfo(ctx, in.GetBearer(), auth.Admin)
	if !token.IsValid || err != nil {
		return &pb.DeleteAreaReply{}, status.Error(codes.PermissionDenied, err.Error())
	}

	err = dbConn.DeleteArea(ctx, in.GetArea())
	if err != nil {
		return &pb.DeleteAreaReply{}, status.Error(codes.Aborted, err.Error())
	}

	return &pb.DeleteAreaReply{}, nil
}

func (s *server) DeleteService(ctx context.Context, in *pb.DeleteServiceRequest) (*pb.DeleteServiceReply, error) {
	token, err := dbConn.QueryAuthTokenInfo(ctx, in.GetBearer(), auth.Admin|auth.Service)
	if !token.IsValid || err != nil {
		return &pb.DeleteServiceReply{}, status.Error(codes.PermissionDenied, err.Error())
	}

	se, err := dbConn.QueryService(ctx, in.GetService(), in.GetArea())
	if err != nil {
		return &pb.DeleteServiceReply{}, status.Error(codes.PermissionDenied, err.Error())
	}

	if !strings.EqualFold(token.Requester, se.Owner) {
		return &pb.DeleteServiceReply{}, status.Error(codes.PermissionDenied, fmt.Sprintf("requester '%s' is NOT owner of service '%s in area '%s", token.Requester, in.GetService(), in.GetArea()))
	}

	err = dbConn.DeleteService(ctx, in.GetService(), in.GetArea(), se.Owner)
	if err != nil {
		return &pb.DeleteServiceReply{}, status.Error(codes.Aborted, err.Error())
	}

	return &pb.DeleteServiceReply{}, nil
}
