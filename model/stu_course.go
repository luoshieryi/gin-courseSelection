package model

import (
	"github.com/jinzhu/gorm"
	"project/util/logs"
	"time"
)

/*
 @Author: as
 @Date: Creat in 17:29 2022/2/13
 @Description: user.go
*/

// StuCourse 用户与课程为多对多的关系
type StuCourse struct {
	gorm.Model
	CourseID string `json:"course_id" gorm:"not null"`
	StuID    string `json:"stu_id" gorm:"not null"`
}

// InitStuCourse 初始化
func InitStuCourse() {
	if DB.HasTable(&StuCourse{}) {
		if err := DB.AutoMigrate(&StuCourse{}).Error; err != nil {
			panic(err)
		}
	} else {
		if err := DB.CreateTable(&StuCourse{}).Error; err != nil {
			panic(err)
		}
	}
}

// MustAdd 一定写入
func (s *StuCourse)MustAdd(){
	c:=0
	for{
		if err:=s.Add();err==nil{
			break
		}
		c++
		// 大于10次，一定是出什么问题了
		// 暂时先这么解决
		if c>10 {
			time.Sleep(time.Second*3)
			c=0
		}
	}
}

func (s *StuCourse)Add()error{
	if err := DB.Model(&StuCourse{}).Create(s).Error; err != nil {
		logs.PrintLogErr(logs.DB,"add user course error:",err)
		return err
	}
	return nil
}

// FindStuAllCourses 查找某个人所有以获取的课程Id
func(s *StuCourse)FindStuAllCourses()([]StuCourse,error){
	var stus []StuCourse
	if err := DB.Model(&StuCourse{}).Find(&stus).Where("stu_id=?", s.StuID).Error; err != nil {
		logs.PrintLogErr(logs.DB,"get user all courses error:",err)
		return nil,err
	}
	return stus,nil
}