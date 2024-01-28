package service

import (
	"context"
	"stepp-backend/src/client"
	"stepp-backend/src/repository"
)

type service struct {
	repository repository.Repository
	client client.Client
}

func NewService(ctx context.Context) Service {
	return &service{
		repository: repository.NewRepository(ctx),
		client: client.NewClient(ctx),
	}
}