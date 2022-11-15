package crpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	"sisco/pb"
)

type ServiceInArea struct {
	Service     string   `json:"service"`
	Area        string   `json:"area"`
	Description string   `json:"description"`
	Host        string   `json:"host"`
	Protocol    string   `json:"protocol"`
	Port        string   `json:"port"`
	Tags        []string ` json:"tags"`
}

func (c *Client) ListServiceInArea(bearer string, serviceName string, areaName string) (*ServiceInArea, error) {
	lsiac := pb.NewListServiceInAreaClient(c.grpcClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	r, err := lsiac.ListServiceInArea(ctx, &pb.ListServiceInAreaRequest{
		Bearer:  bearer,
		Service: serviceName,
		Area:    areaName,
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("adding area failed: %s", err))
	}

	data := ServiceInArea{
		Service:     r.GetService(),
		Area:        r.GetArea(),
		Description: r.GetDescription(),
		Host:        r.GetHost(),
		Protocol:    r.GetProtocol(),
		Port:        r.GetPort(),
		Tags:        r.GetTags(),
	}

	return &data, err
}

func (c *Client) ListServices(bearer string) (*[]string, error) {
	lsiac := pb.NewListServicesClient(c.grpcClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	r, err := lsiac.ListServices(ctx, &pb.ListServicesRequest{
		Bearer: bearer,
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("adding area failed: %s", err))
	}

	data := r.GetServices()

	return &data, err
}
