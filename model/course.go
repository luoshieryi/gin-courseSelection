package model

type Course struct {
	ID        int64  `gorm:"auto_increment"`
	Name      string `gorm:"not null"`
	Cap       int    `gorm:"not null"`
	TeacherID int64  `gorm:""`
}
