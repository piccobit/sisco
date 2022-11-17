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
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ServiceExtended struct {
	Name        string   `json:"name"`
	Area        string   `json:"area"`
	Description string   `json:"description"`
	Host        string   `json:"host"`
	Protocol    string   `json:"protocol"`
	Port        string   `json:"port"`
	Tags        []string ` json:"tags"`
}

type Tag struct {
	Name string `json:"tag"`
}

func (c *Client) ListServiceInArea(bearer string, serviceName string, areaName string) (*ServiceExtended, error) {
	l := pb.NewListServiceInAreaClient(c.grpcClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	r, err := l.ListServiceInArea(ctx, &pb.ListServiceInAreaRequest{
		Bearer: bearer,
		Name:   serviceName,
		Area:   areaName,
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("listing service in area failed: %v", err))
	}

	data := ServiceExtended{
		Name:        r.GetName(),
		Area:        r.GetArea(),
		Description: r.GetDescription(),
		Host:        r.GetHost(),
		Protocol:    r.GetProtocol(),
		Port:        r.GetPort(),
		Tags:        r.GetTags(),
	}

	return &data, err
}

func (c *Client) ListServices(bearer string, areaName string) ([]*Service, error) {
	l := pb.NewListServicesClient(c.grpcClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	r, err := l.ListServices(ctx, &pb.ListServicesRequest{
		Bearer: bearer,
		Area:   areaName,
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("listing services failed: %v", err))
	}

	var data []*Service

	pbServices := r.GetServices()

	for _, pbs := range pbServices {
		d := Service{
			Name:        pbs.GetName(),
			Description: pbs.GetDescription(),
		}
		data = append(data, &d)
	}

	return data, err
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
