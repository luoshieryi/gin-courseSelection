package model

type StudentCourse struct {
	ID        int64 `gorm:"auto_increment"`
	StudentID int64 `gorm:"not null"`
	CourseID  int64 `gorm:""`
}
