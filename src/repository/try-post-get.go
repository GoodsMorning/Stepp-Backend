package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	table "stepp-backend/src/constant/database"
	"stepp-backend/src/model/database"
	response2 "stepp-backend/src/model/response"
)

func (r *repository) GetTestPost(ctx context.Context, postId int) (*response2.PostResponse, error) {
	responseDB := database.PostResponseDB{}
	err := r.gormDb.Table(table.POST+" post").
		Select("*, post.like_count AS like_count_test").
		Joins("JOIN "+table.ORIGINAL+" original ON post.original_id = original.original_id").
		Joins("JOIN "+table.LOCATION+" location ON original.location_id  = location.location_id").
		Joins("JOIN "+table.USER+" owner ON post.owner_id = owner.user_id").
		Joins("JOIN "+table.USER+" original_owner On original_owner.user_id = original.owner_id").
		Where("post_id = ?", postId).
		First(&responseDB).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Err: Post Not Found !")
			return nil, err
		}
		return nil, err
	}
	response := response2.PostResponse{}
	mapPostResponseFromDB(&responseDB, &response)
	return &response, nil
}

func mapPostResponseFromDB(db *database.PostResponseDB, response *response2.PostResponse) {
	response.PostId = db.PostId
	response.LikeCount = db.LikeCount
}
