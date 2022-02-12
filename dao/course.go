package dao

import (
	"project/model"
)

/*
 @Author: as
 @Date: Creat in 15:52 2022/2/12
 @Description: user.go
*/

// GetCourseCap 获取课程当前的容量
func GetCourseCap(name string) (int, error) {
	// 从数据库拿
	c := model.Course{}
	c.Name=name
	now,err:=c.GetNowCourse()
	if err!=nil{
		return 0, err
	}
	return now.Cap, nil
}

// BookCourse 成功抢到课
//TODO: 用户写入抢到的课
func BookCourse(courseId,stuId string)error{

	return nil
}

//TODO: 用户是否拥有该课程
func StuHaveCourse(courseId,stuId string)bool{
	return false
}
