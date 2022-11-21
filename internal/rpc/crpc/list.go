package crpc

import (
	"context"
	"errors"
	"fmt"
	"sisco/internal/rpc/pb"
	"time"
)

type Area struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Service struct {
	Name        string   `json:"name"`
	Area        string   `json:"area"`
	Description string   `json:"description"`
	Host        string   `json:"host"`
	Protocol    string   `json:"protocol"`
	Port        string   `json:"port"`
	Tags        []string ` json:"tags"`
	Available   bool     `json:"available"`
}

type Tag struct {
	Name string `json:"tag"`
}

func (c *Client) ListService(bearer string, serviceName string, areaName string) (*Service, error) {
	l := pb.NewListServiceClient(c.grpcClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	r, err := l.ListService(ctx, &pb.ListServiceRequest{
		Bearer: bearer,
		Name:   serviceName,
		Area:   areaName,
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("listing service in area failed: %v", err))
	}

	data := Service{
		Name:        r.GetService().GetName(),
		Area:        r.GetService().GetArea(),
		Description: r.GetService().GetDescription(),
		Host:        r.GetService().GetHost(),
		Protocol:    r.GetService().GetProtocol(),
		Port:        r.GetService().GetPort(),
		Tags:        r.GetService().GetTags(),
		Available:   r.GetService().GetAvailable(),
	}

	return &data, nil
}

func (c *Client) ListServices(bearer string, areaName string, tagName string) ([]*Service, error) {
	l := pb.NewListServicesClient(c.grpcClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	r, err := l.ListServices(ctx, &pb.ListServicesRequest{
		Bearer: bearer,
		Area:   areaName,
		Tag:    tagName,
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("listing services failed: %v", err))
	}

	var data []*Service

	pbServices := r.GetServices()

	for _, pbs := range pbServices {
		d := Service{
			Name:        pbs.GetName(),
			Area:        pbs.GetArea(),
			Description: pbs.GetDescription(),
			Protocol:    pbs.GetProtocol(),
			Host:        pbs.GetHost(),
			Port:        pbs.GetPort(),
			Tags:        pbs.GetTags(),
			Available:   pbs.GetAvailable(),
		}
		data = append(data, &d)
	}

	return data, nil
}

func (c *Client) ListAreas(bearer string) ([]*Area, error) {
	l := pb.NewListAreasClient(c.grpcClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	r, err := l.ListAreas(ctx, &pb.ListAreasRequest{
		Bearer: bearer,
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("listing areas failed: %v", err))
	}

	var data []*Area

	for _, pba := range r.GetAreas() {
		d := Area{
			Name:        pba.GetName(),
			Description: pba.GetDescription(),
		}
		data = append(data, &d)
	}

	return data, nil
}

func (c *Client) ListTags(bearer string) ([]*Tag, error) {
	l := pb.NewListTagsClient(c.grpcClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	r, err := l.ListTags(ctx, &pb.ListTagsRequest{
		Bearer: bearer,
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("listing tags failed: %v", err))
	}

	var data []*Tag

	for _, pba := range r.GetTags() {
		d := Tag{
			Name: pba.GetName(),
		}
		data = append(data, &d)
	}

	return data, nil
}
