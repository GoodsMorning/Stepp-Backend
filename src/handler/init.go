package handler

import (
	"context"
	"stepp-backend/src/service"
)


type handler struct {
	service	service.Service
}

func NewHandler(ctx context.Context) Handler {
	return &handler{
		service: service.NewService(ctx),
	}
}