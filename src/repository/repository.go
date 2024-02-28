package repository

import (
	"context"
	"stepp-backend/src/model/database"
	response2 "stepp-backend/src/model/response"
)

type Repository interface {
	GetPost(ctx context.Context, postId int) (*database.PostDb, error)
	GetOriginal(ctx context.Context, originalId int) (*database.OriginalDb, error)
	GetUser(ctx context.Context, userId int) (*database.UserDb, error)
	GetLocation(ctx context.Context, locationId int) (*database.LocationDb, error)

	GetTestPost(ctx context.Context, postId int) (*response2.PostResponse, error)
}
