package database

type UserDb struct {
	UserId        int64  `json:"user_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	ProfileImgUrl string `json:"profile_img_url"`
	SteppCount    int64  `json:"stepp_count"`
	FollowerCount int64  `json:"follower_count"`
	OriginalCount int64  `json:"original_count"`
	Bio           string `json:"bio"`
	CreatedAt     string `json:"created_at"`
}
