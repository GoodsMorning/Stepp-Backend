package repository

import (
	"context"
	"database/sql"
	"log"
	"stepp-backend/initializer"
)

type repository struct {
	db *sql.DB
	// graphql
}

func NewRepository(ctx context.Context) Repository {
	db, err := initializer.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	return &repository{
		db: db,
	}
}
