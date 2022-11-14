package srpc

import (
	"context"

	"sisco/pb"
)

func (s *server) ListServiceInArea(ctx context.Context, in *pb.ListServiceInAreaRequest) (*pb.ListServiceInAreaReply, error) {
	var err error

	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), true)
	if !tokenIsValid || err != nil {
		return &pb.ListServiceInAreaReply{}, err
	}

	se, err := dbConn.QueryServiceInArea(ctx, in.GetService(), in.GetArea())
	if err != nil {
		return &pb.ListServiceInAreaReply{
			Service:     in.GetService(),
			Area:        in.GetArea(),
			Description: se.Description,
			Protocol:    se.Protocol,
			Host:        se.Host,
			Port:        se.Port,
			Tags:        nil,
		}, err
	}

	return &pb.ListServiceInAreaReply{}, err
}
