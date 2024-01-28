package repository

import "context"

type repository struct {
	// graphql
	// db
}

func NewRepository(ctx context.Context) Repository {
	return &repository{}
}