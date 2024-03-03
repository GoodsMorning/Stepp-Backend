package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	table "stepp-backend/src/constant/database"
	"stepp-backend/src/model/database"
	response2 "stepp-backend/src/model/response"
	"strings"
)

func (r *repository) GetPost(ctx context.Context, postId int) (*response2.PostResponse, error) {
	responseDB := database.PostResponseDB{}

	selectArr := []string{
		"*",
		"original.owner_id as original_owner_user_id",
		"original_owner.name as original_owner_name",
		"original_owner.profile_img_url as original_owner_profile_img_url",
		"original.name as original_name",
		"location.location_id as location_id ",
		"location.name as location_name ",
		"location.geo_point as location_geo_point ",
		"location.stepp_count as location_stepp_count ",
		"original.total_stepp as original_total_stepp",
		"original.stepp_count as original_stepp_count",
		"original.original_post_id as original_post_id",
		"owner.user_id as post_owner_user_id",
		"owner.name as post_owner_name",
		"owner.profile_img_url as post_owner_profile_img_url",
	}

	err := r.gormDb.Table(table.POST+" post").
		Select(strings.Join(selectArr, " , ")).
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
	{ //Original
		response.Original.OriginalId = db.OriginalId
		{ // Owner
			response.Original.Owner.UserId = db.OriginalOwnerUserId
			response.Original.Owner.Name = db.OriginalOwnerName
			response.Original.Owner.ProfileImgUrl = db.OriginalOwnerProfileImgUrl
		}
		response.Original.Name = db.OriginalName
		{ // Location
			response.Original.Location.LocationId = db.LocationId
			response.Original.Location.Name = db.LocationName
			response.Original.Location.GeoPoint = db.LocationGeoPoint
			response.Original.Location.SteppCount = db.LocationSteppCount
		}
		response.Original.TotalStepp = db.OriginalTotalStepp
		response.Original.SteppCount = db.OriginalSteppCount
		response.Original.OriginalPostId = db.OriginalPostId
	}
	{ // Owner
		response.Owner.UserId = db.PostOwnerUserId
		response.Owner.Name = db.PostOwnerName
		response.Owner.ProfileImgUrl = db.PostOwnerProfileImgUrl
	}
	response.CoverUrl = db.CoverUrl
	response.CreatedAt = db.CreatedAt
	response.Caption = db.Caption
	response.LikeCount = db.LikeCount
}
