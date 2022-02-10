package model

type Member struct {
	ID       int64  `gorm:"auto_increment"`
	Nickname string `gorm:"not null"`
	Username string `gorm:"not null;UNIQUE"`
	Password string `gorm:"not null"`
	UserType int    `gorm:"not null"`
	Deleted  bool   `gorm:"not null"`
}
