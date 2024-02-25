package repository

import (
	"context"
	"stepp-backend/src/model/database"
)

type Repository interface {
	GetPost(ctx context.Context, postId int) (*database.PostDb, error)
	GetOriginal(ctx context.Context, originalId int) (*database.OriginalDb, error)
	GetUser(ctx context.Context, userId int) (*database.UserDb, error)
	GetLocation(ctx context.Context, locationId int) (*database.LocationDb, error)
}
