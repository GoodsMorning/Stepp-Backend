package database

type PostResponseDB struct {
	PostId     int64 `json:"post_id"`
	OriginalId int64 `json:"original_id"`
	//Original   OriginalFragment `json:"original"`
	OwnerId int64 `json:"owner_id"`
	//Owner      UserFragmentPost `json:"owner"`
	CoverUrl  string `json:"cover_url"`
	CreatedAt string `json:"created_at"`
	Caption   string `json:"caption"`
	LikeCount int64  `json:"like_count" gorm:"column:like_count_test"`
}
