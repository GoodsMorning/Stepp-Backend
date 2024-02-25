package place

type OriginalPlace struct {
	OriginalID int64

	//	content
	Name string

	// foreign key
	OriginalPostID int64
	OwnerID        int64
	LocationID     int64

	// attribute
	// CreateAt	timestamp
	TotalStepp int

	// stat
	LikeCount  int
	SteppCount int
}

type PostPlace struct {
	PostID int64

	// content
	Caption     string
	CoverImgUrl string

	// foreign key
	OriginalID int64
	OwnerID    int64

	// attribute
	// CreateAt	timestamp

	// stat
	LikeCount int
}
