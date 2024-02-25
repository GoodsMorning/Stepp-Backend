package repository

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"log"
	"stepp-backend/initializer"
)

type repository struct {
	db     *sql.DB
	gormDb *gorm.DB
}

func NewRepository(ctx context.Context) Repository {
	db, gormDb, err := initializer.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	//defer db.Close()

	return &repository{
		db:     db,
		gormDb: gormDb,
	}
}
