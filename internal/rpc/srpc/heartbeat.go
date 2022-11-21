package srpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sisco/internal/auth"
	"sisco/internal/rpc/pb"
	"time"
)

func (s *server) Heartbeat(ctx context.Context, in *pb.HeartbeatRequest) (*pb.HeartbeatReply, error) {
	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), auth.Admin|auth.Service)
	if !tokenIsValid || err != nil {
		return &pb.HeartbeatReply{}, status.Error(codes.PermissionDenied, err.Error())
	}

	err = dbConn.UpdateServiceAvailableHeartbeat(
		ctx,
		in.GetService(),
		in.GetArea(),
		true,
		time.Now(),
	)
	if err != nil {
		return &pb.HeartbeatReply{}, status.Error(codes.Aborted, err.Error())
	}

	return &pb.HeartbeatReply{}, nil
}
