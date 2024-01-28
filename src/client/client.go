package client

import "context"

type Client interface {

}

func NewClient(ctx context.Context) Client {
	return &client{}
}