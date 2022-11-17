package srpc

import (
	"context"
	"sisco/internal/rpc/pb"
)

func (s *server) ListService(ctx context.Context, in *pb.ListServiceRequest) (*pb.ListServiceReply, error) {
	var err error

	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), true)
	if !tokenIsValid || err != nil {
		return &pb.ListServiceReply{}, err
	}

	se, err := dbConn.QueryServiceInArea(ctx, in.GetName(), in.GetArea())
	if err != nil {
		return &pb.ListServiceReply{}, err
	}

	var tags []string

	for _, tag := range se.Edges.Tags {
		tags = append(tags, tag.Name)
	}

	return &pb.ListServiceReply{
		Name:        in.GetName(),
		Area:        in.GetArea(),
		Description: se.Description,
		Protocol:    se.Protocol,
		Host:        se.Host,
		Port:        se.Port,
		Tags:        tags,
	}, nil
}

func (s *server) ListServices(ctx context.Context, in *pb.ListServicesRequest) (*pb.ListServicesReply, error) {
	var err error

	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), true)
	if !tokenIsValid || err != nil {
		return &pb.ListServicesReply{}, err
	}

	r, err := dbConn.QueryServices(ctx)
	if err != nil {
		return &pb.ListServicesReply{}, err
	}

	var data []*pb.Service

	for _, d := range r {
		e := pb.Service{
			Name:        d.Name,
			Description: d.Description,
		}
		data = append(data, &e)
	}

	return &pb.ListServicesReply{
		Services: data,
	}, nil
}

func (s *server) ListAreas(ctx context.Context, in *pb.ListAreasRequest) (*pb.ListAreasReply, error) {
	var err error

	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), true)
	if !tokenIsValid || err != nil {
		return &pb.ListAreasReply{}, err
	}

	r, err := dbConn.QueryAreas(ctx)
	if err != nil {
		return &pb.ListAreasReply{}, err
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
	var err error

	tokenIsValid, err := dbConn.CheckToken(ctx, in.GetBearer(), true)
	if !tokenIsValid || err != nil {
		return &pb.ListTagsReply{}, err
	}

	r, err := dbConn.QueryTags(ctx)
	if err != nil {
		return &pb.ListTagsReply{}, err
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
