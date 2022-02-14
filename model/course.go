package model

import (
	"log"
	"project/util/logs"
)

/*
 @Author: as
 @Date: Creat in 15:24 2022/2/12
 @Description: 课程
*/


type Course struct {
	ID        int64  `gorm:"auto_increment"`
	Name      string `gorm:"not null"`
	Cap       int    `gorm:"not null"`
	TeacherID int64  `gorm:""`
}

func InitCourse() {
	if DB.HasTable(&Course{}) {
		if err := DB.AutoMigrate(&Course{}).Error; err != nil {
			panic(err)
		}
	} else {
		if err := DB.CreateTable(&Course{}).Error; err != nil {
			panic(err)
		}
	}

}

// CreateCourse 增加单个课程
func (c *Course) CreateCourse() error {
	if err := DB.Model(&Course{}).Create(c).Error; err != nil {
		logs.PrintLogErr(logs.DB, "Create Courses error", err)
		return err
	}
	return nil
}

// GetNowCourse 获取当前课程信息
func (c *Course) GetNowCourse() (Course, error) {
	var cs Course
	if err := DB.Model(&Course{}).First(&cs).Where("name=?", c.Name).Error; err != nil {
		log.Println(logs.DB, "GetAll Courses error:", err)
		return cs, err
	}
	return cs, nil
}
