package database

type PostResponseDB struct {
	PostId int64

	OriginalId                 int64
	OriginalOwnerUserId        int64  `gorm:"column:original_owner_user_id"`
	OriginalOwnerName          string `gorm:"column:original_owner_name"`
	OriginalOwnerProfileImgUrl string `gorm:"column:original_owner_profile_img_url"`

	OriginalName string `gorm:"column:original_name"`

	LocationId         int64  `gorm:"column:location_id"`
	LocationName       string `gorm:"column:location_name"`
	LocationGeoPoint   string `gorm:"column:location_geo_point"`
	LocationSteppCount int64  `gorm:"column:location_stepp_count"`

	OriginalTotalStepp int64 `gorm:"column:original_total_stepp"`
	OriginalSteppCount int64 `gorm:"column:original_stepp_count"`
	OriginalPostId     int64 `gorm:"column:original_post_id"`

	PostOwnerUserId        int64  `gorm:"column:post_owner_user_id"`
	PostOwnerName          string `gorm:"column:post_owner_name"`
	PostOwnerProfileImgUrl string `gorm:"column:post_owner_profile_img_url"`

	CoverUrl  string
	CreatedAt string
	Caption   string
	LikeCount int64
}
