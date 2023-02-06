package models

// Table = Gorm Model
type Video struct {
	CommonModel
	User          User `gorm:"foreignKey:UserId"`
	UserId        int
	Title         string
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int
	CommentCount  int
	IsFavorite    bool
}
