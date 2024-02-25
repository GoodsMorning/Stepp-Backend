package service

import (
	"context"
	basepayload "stepp-backend/src/model/base"
)

type Service interface {
	GetHello(ctx context.Context) (*basepayload.BasePayload, error)
	GetPost(ctx context.Context, postId int) (*basepayload.BasePayload, error)
}
