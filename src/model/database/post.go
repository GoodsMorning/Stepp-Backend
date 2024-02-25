package database

type PostDb struct {
	PostId     int64  `json:"post_id"`
	OriginalId int64  `json:"original_id"`
	OwnerId    int64  `json:"owner_id"`
	CoverUrl   string `json:"cover_url"`
	CreatedAt  string `json:"created_at"`
	Caption    string `json:"caption"`
	LikeCount  int64  `json:"like_count"`
}

type OriginalDb struct {
	OriginalId     int64  `json:"original_id"`
	OwnerId        int64  `json:"owner_id"`
	CreatedAt      string `json:"created_at"`
	Name           string `json:"name"`
	LocationId     int64  `json:"location_id"`
	TotalStepp     int64  `json:"total_stepp"`
	SteppCount     int64  `json:"stepp_count"`
	OriginalPostId int64  `json:"original_post_id"`
}
