package crpc

import (
	"context"
	"time"

	"sisco/pb"
)

func (c *Client) Login(user string, password string) (string, bool, error) {
	l := pb.NewLoginClient(c.grpcClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	r, err := l.Login(ctx, &pb.LoginRequest{
		User:     user,
		Password: password,
	})
	if err != nil {
		return "", false, err
	}

	return r.GetToken(), r.GetIsAdminToken(), err
}
