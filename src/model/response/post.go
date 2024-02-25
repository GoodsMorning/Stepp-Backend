package response

import "stepp-backend/src/model/database"

type PostResponse struct {
	PostId    int64            `json:"post_id"`
	Original  OriginalFragment `json:"original"`
	Owner     UserFragmentPost `json:"owner"`
	CoverUrl  string           `json:"cover_url"`
	CreatedAt string           `json:"created_at"`
	Caption   string           `json:"caption"`
	LikeCount int64            `json:"like_count"`
}

type OriginalFragment struct {
	OriginalId     int64               `json:"original_id"`
	Owner          UserFragmentPost    `json:"owner"`
	Name           string              `json:"name"`
	Location       database.LocationDb `json:"location"`
	TotalStepp     int64               `json:"total_stepp"`
	SteppCount     int64               `json:"stepp_count"`
	OriginalPostId int64               `json:"original_post_id"`
}
