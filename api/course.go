package api

import (
	"github.com/gin-gonic/gin"
	"project/service"
	"project/types"
)

/*
 @Author: as
 @Date: Creat in 15:40 2022/2/12
 @Description: 课程相关
*/

// GetCourse 获取课程
func GetCourse(c *gin.Context) {

}

// BookCourse 抢课
func BookCourse(c *gin.Context) {

	q:=types.BookCourseRequest{}
	if err := c.ShouldBind(&q);err!=nil{
		c.JSON(200,types.BookCourseResponse{Code: types.ParamInvalid})
		return
	}


	// 抢课
	switch service.BookCourse(q.CourseID,q.StudentID) {
	case service.CourseOver:
		c.JSON(200,types.BookCourseResponse{Code: types.CourseNotAvailable})
	case service.StuHaveCourse:
		c.JSON(200,types.BookCourseResponse{Code: types.StudentHasCourse})
	case nil:
		c.JSON(200,types.BookCourseResponse{Code: types.OK})
	default:
		c.JSON(200,types.BookCourseResponse{Code: types.UnknownError})
	}

}
