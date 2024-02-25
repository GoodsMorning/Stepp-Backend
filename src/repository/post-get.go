package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	table "stepp-backend/src/constant/database"
	"stepp-backend/src/model/database"
)

func (r *repository) GetPost(ctx context.Context, postId int) (*database.PostDb, error) {

	response := database.PostDb{}

	err := r.gormDb.Table(table.POST).Where("post_id = ?", postId).First(&response).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Err: Post Not Found !")
			return nil, err
		}
		return nil, err
	}

	return &response, nil
}
