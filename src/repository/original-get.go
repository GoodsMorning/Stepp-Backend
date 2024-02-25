package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	table "stepp-backend/src/constant/database"
	"stepp-backend/src/model/database"
)

func (r *repository) GetOriginal(ctx context.Context, originalId int) (*database.OriginalDb, error) {

	response := database.OriginalDb{}
	err := r.gormDb.Table(table.ORIGINAL).Where("original_id = ?", originalId).First(&response).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Err: Post Not Found !")
			return nil, err
		}
		return nil, err
	}

	return &response, nil
}
