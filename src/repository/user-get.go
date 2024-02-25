package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	table "stepp-backend/src/constant/database"
	"stepp-backend/src/model/database"
)

func (r *repository) GetUser(ctx context.Context, userId int) (*database.UserDb, error) {
	response := database.UserDb{}

	err := r.gormDb.Table(table.USER).Where("user_id = ?", userId).First(&response).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Err: Post Not Found !")
			return nil, err
		}
		return nil, err
	}

	return &response, nil
}
