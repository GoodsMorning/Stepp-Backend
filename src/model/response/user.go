package response

type UserFragmentPost struct {
	UserId        int64  `json:"user_id"`
	Name          string `json:"name"`
	ProfileImgUrl string `json:"profile_img_url"`
}
