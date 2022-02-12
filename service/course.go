package service

import (
	"project/dao"
	"project/model"
	"project/types"
	"strconv"
)

// CreateCourse 创建用户
func CreateCourse(request types.CreateCourseRequest) (string, types.ErrNo) {
	course := dao.GetCourseByName(request.Name)
	if course.ID != 0 {
		return strconv.FormatInt(course.ID, 10), types.UnknownError
	}

	entity := model.Course{
		Name: request.Name,
		Cap:  request.Cap,
	}
	id, err := dao.CreateCourse(entity)
	if err != nil {
		return "", types.UnknownError
	}
	return strconv.FormatInt(id, 10), types.OK
}

func GetCourse(request types.GetCourseRequest) (types.TCourse, types.ErrNo) {
	tCourse := types.TCourse{}
	_id, err := strconv.ParseInt(request.CourseID, 10, 64)
	if err != nil {
		return tCourse, types.UnknownError
	}

	course := dao.GetCourseById(_id)
	if course.ID == 0 {
		return tCourse, types.CourseNotExisted
	}

	res := types.TCourse{
		CourseID:  strconv.FormatInt(course.ID, 10),
		Name:      course.Name,
		TeacherID: strconv.FormatInt(course.TeacherID, 10),
	}
	return res, types.OK
}
