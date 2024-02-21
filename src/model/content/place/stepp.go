package place

type OriginalSteppPlace struct {
	OriginalSteppID int64 // uuid

	//	content
	Name string

	// foreign key
	OriginalID     int64
	OriginalPostID int64
	LocationID     int64

	// attribute
	Sequence int
}

type SteppPlace struct {
	SteppID int64

	// content
	Caption    string
	ContentUrl string
	IsVdo      bool

	// foreign key
	PostID          int64
	OriginalSteppID int64
	LocationID      int64

	// attribute
	Sequence int
}
