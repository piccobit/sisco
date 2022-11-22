package srpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sisco/internal/auth"
	"sisco/internal/rpc/pb"
)

func (s *server) ListService(ctx context.Context, in *pb.ListServiceRequest) (*pb.ListServiceReply, error) {
	token, err := dbConn.QueryAuthTokenInfo(ctx, in.GetBearer(), auth.Admin|auth.Service|auth.User)
	if !token.IsValid || err != nil {
		return &pb.ListServiceReply{}, status.Error(codes.PermissionDenied, err.Error())
	}

	se, err := dbConn.QueryService(ctx, in.GetName(), in.GetArea())
	if err != nil {
		return &pb.ListServiceReply{}, status.Error(codes.Aborted, err.Error())
	}

	var tags []string

	for _, tag := range se.Edges.Tags {
		tags = append(tags, tag.Name)
	}

	svc := pb.Service{
		Name:        in.GetName(),
		Area:        in.GetArea(),
		Description: se.Description,
		Protocol:    se.Protocol,
		Host:        se.Host,
		Port:        se.Port,
		Tags:        tags,
		Available:   se.Available,
		Owner:       se.Owner,
	}

	return &pb.ListServiceReply{
		Service: &svc,
	}, nil
}

func (s *server) ListServices(ctx context.Context, in *pb.ListServicesRequest) (*pb.ListServicesReply, error) {
	token, err := dbConn.QueryAuthTokenInfo(ctx, in.GetBearer(), auth.Admin|auth.Service|auth.User)
	if !token.IsValid || err != nil {
		return &pb.ListServicesReply{}, status.Error(codes.PermissionDenied, err.Error())
	}

	r, err := dbConn.QueryServices(ctx, in.GetArea(), in.GetTag())
	if err != nil {
		return &pb.ListServicesReply{}, status.Error(codes.Aborted, err.Error())
	}

	var data []*pb.Service

	for _, d := range r {
		var tags []string

		for _, tag := range d.Edges.Tags {
			tags = append(tags, tag.Name)
		}

		e := pb.Service{
			Name:        d.Name,
			Area:        d.Edges.Area.Name,
			Description: d.Description,
			Protocol:    d.Protocol,
			Host:        d.Host,
			Port:        d.Port,
			Tags:        tags,
			Available:   d.Available,
			Owner:       d.Owner,
		}
		data = append(data, &e)
	}

	return &pb.ListServicesReply{
		Services: data,
	}, nil
}

func (s *server) ListAreas(ctx context.Context, in *pb.ListAreasRequest) (*pb.ListAreasReply, error) {
	token, err := dbConn.QueryAuthTokenInfo(ctx, in.GetBearer(), auth.Admin|auth.Service|auth.User)
	if !token.IsValid || err != nil {
		return &pb.ListAreasReply{}, status.Error(codes.PermissionDenied, err.Error())
	}

	r, err := dbConn.QueryAreas(ctx)
	if err != nil {
		return &pb.ListAreasReply{}, status.Error(codes.Aborted, err.Error())
	}

	var data []*pb.Area

	for _, d := range r {
		e := pb.Area{
			Name:        d.Name,
			Description: d.Description,
		}
		data = append(data, &e)
	}

	return &pb.ListAreasReply{
		Areas: data,
	}, nil
}

func (s *server) ListTags(ctx context.Context, in *pb.ListTagsRequest) (*pb.ListTagsReply, error) {
	token, err := dbConn.QueryAuthTokenInfo(ctx, in.GetBearer(), auth.Admin|auth.Service|auth.User)
	if !token.IsValid || err != nil {
		return &pb.ListTagsReply{}, status.Error(codes.PermissionDenied, err.Error())
	}

	r, err := dbConn.QueryTags(ctx)
	if err != nil {
		return &pb.ListTagsReply{}, status.Error(codes.Aborted, err.Error())
	}

	var data []*pb.Tag

	for _, d := range r {
		e := pb.Tag{
			Name: d.Name,
		}
		data = append(data, &e)
	}

	return &pb.ListTagsReply{
		Tags: data,
	}, nil
}
