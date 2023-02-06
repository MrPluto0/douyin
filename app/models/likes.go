package models

// Table = Gorm Model
type Like struct {
	CommonModel
	User    User `gorm:"foreignKey:UserId"`
	UserId  int
	Video   Video `gorm:"foreignKey:VideoId"`
	VideoId int
}
