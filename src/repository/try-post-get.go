package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	table "stepp-backend/src/constant/database"
	response2 "stepp-backend/src/model/response"
)

func (r *repository) GetTestPost(ctx context.Context, postId int) (*response2.PostResponse, error) {
	response := response2.PostResponse{}

	err := r.gormDb.Table(table.POST+" post").
		Select("*, original.original_id AS vivi").
		Joins("JOIN "+table.ORIGINAL+" original ON post.original_id = original.original_id").
		Joins("JOIN "+table.LOCATION+" location ON original.location_id  = location.location_id").
		Joins("JOIN "+table.USER+" owner ON post.owner_id = owner.user_id").
		Joins("JOIN "+table.USER+" original_owner On original_owner.user_id = original.owner_id").
		Where("post_id = ?", postId).
		First(&response).Error

	//err := r.gormDb.Table(table.POST+" post").Select("*").
	//	Joins("JOIN "+table.ORIGINAL+" original ON post.original_id = original.original_id").
	//	Where("post_id = ?", postId).
	//	Preload("OriginalId").First(&response).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Err: Post Not Found !")
			return nil, err
		}
		return nil, err
	}

	return &response, nil
}
