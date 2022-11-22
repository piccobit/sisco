package srpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sisco/internal/auth"
	"sisco/internal/rpc/pb"
	"strings"
	"time"
)

func (s *server) Heartbeat(ctx context.Context, in *pb.HeartbeatRequest) (*pb.HeartbeatReply, error) {
	token, err := dbConn.QueryAuthTokenInfo(ctx, in.GetBearer(), auth.Admin|auth.Service)
	if !token.IsValid || err != nil {
		return &pb.HeartbeatReply{}, status.Error(codes.PermissionDenied, err.Error())
	}

	se, err := dbConn.QueryService(ctx, in.GetService(), in.GetArea())
	if err != nil {
		return &pb.HeartbeatReply{}, status.Error(codes.PermissionDenied, err.Error())
	}

	if !strings.EqualFold(token.Requester, se.Owner) {
		return &pb.HeartbeatReply{}, status.Error(codes.PermissionDenied, fmt.Sprintf("requester '%s' is NOT owner of service '%s in area '%s", token.Requester, in.GetService(), in.GetArea()))
	}

	err = dbConn.UpdateServiceAvailableHeartbeat(
		ctx,
		in.GetService(),
		in.GetArea(),
		se.Owner,
		true,
		time.Now(),
	)
	if err != nil {
		return &pb.HeartbeatReply{}, status.Error(codes.Aborted, err.Error())
	}

	return &pb.HeartbeatReply{}, nil
}
