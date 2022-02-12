package service

import (
	"project/dao"
	"project/model"
	"project/types"
	"strconv"
)

// CreateCourse 创建课程
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

func BindCourse(request types.BindCourseRequest) types.ErrNo {
	Id, err := strconv.ParseInt(request.CourseID, 10, 64)
	if err != nil {
		return types.UnknownError
	}

	teacherId, err := strconv.ParseInt(request.TeacherID, 10, 64)
	if err != nil {
		return types.UnknownError
	}

	res := dao.GetCourseById(Id)
	if res.ID == 0 {
		return types.CourseNotExisted
	}
	if res.TeacherID != 0 {
		return types.CourseHasBound
	}

	res.TeacherID = teacherId
	err = dao.UpdateCourse(res)
	if err != nil {
		return types.UnknownError
	}
	return types.OK
}

//其实teacherid用不上
func UnbindCourse(request types.UnbindCourseRequest) types.ErrNo {
	Id, err := strconv.ParseInt(request.CourseID, 10, 64)
	if err != nil {
		return types.UnknownError
	}

	teacherId, err := strconv.ParseInt(request.TeacherID, 10, 64)
	if err != nil {
		return types.UnknownError
	}

	res := dao.GetCourseById(Id)
	if res.ID == 0 {
		return types.CourseNotExisted
	}
	if res.TeacherID == 0 {
		return types.CourseNotBind
	}
	if res.TeacherID != teacherId {
		return types.UnknownError
	}
	res.TeacherID = 0
	println(res.ID)
	err = dao.DeleteTeacherByID(res.ID)
	if err != nil {
		return types.UnknownError
	}
	return types.OK
}

func GetTeacherCourse(request types.GetTeacherCourseRequest) ([]*types.TCourse, types.ErrNo) {
	tcource := make([]*types.TCourse, 0)
	TeacherID, err := strconv.ParseInt(request.TeacherID, 10, 64)
	if err != nil {
		return tcource, types.UnknownError
	}
	tcource, err = dao.GetCourseByTeacherId(TeacherID)
	if err != nil {
		return tcource, types.UnknownError
	}
	return tcource, types.OK
}
