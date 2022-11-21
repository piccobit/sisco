package crpc

import (
	"context"
	"sisco/internal/auth"
	"sisco/internal/rpc/pb"
	"time"
)

func (c *Client) Login(user string, password string) (string, auth.Permissions, error) {
	l := pb.NewLoginClient(c.grpcClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	r, err := l.Login(ctx, &pb.LoginRequest{
		User:     user,
		Password: password,
	})
	if err != nil {
		return "", auth.Unknown, err
	}

	return r.GetToken(), auth.Permissions(r.GetPermissions()), nil
}
