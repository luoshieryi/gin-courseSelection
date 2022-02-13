package dao

import (
	"project/model"
	"project/types"
	"strconv"
)

func CreateCourse(course model.Course) (int64, error) {
	err := model.DB.Create(&course).Error
	if err != nil {
		return 0, err
	}
	return course.ID, nil
}

// GetCourseByName 通过课程名查询课程
func GetCourseByName(name string) model.Course {
	course := model.Course{}

	model.DB.Find(&course, "name = ?", name)

	return course
}

// GetCourseById 通过ID查询课程
func GetCourseById(id int64) model.Course {
	course := model.Course{}
	model.DB.Find(&course, "id = ?", id)
	return course
}

// UpdateCourse 更新课程信息
func UpdateCourse(course model.Course) error {
	err := model.DB.Model(&course).Update(&course).Error

	return err
}

func DeleteTeacherByID(id int64) error {
	err := model.DB.Model(&model.Course{}).Where("id = ?", id).Update(map[string]interface{}{"TeacherID": 0}).Error

	return err
}

// GetCourseByTeacherId 得到老师的ID对应的课程
func GetCourseByTeacherId(TeacherID int64) ([]*types.TCourse, error) {
	res := make([]model.Course, 0)
	err := model.DB.Find(&res, "teacher_id = ?", TeacherID).Error
	courses := make([]*types.TCourse, 0)
	for _, data := range res {
		course := types.TCourse{CourseID: strconv.FormatInt(data.ID, 10), Name: data.Name, TeacherID: strconv.FormatInt(data.TeacherID, 10)}
		courses = append(courses, &course)
	}
	return courses, err
}

/*
 @Author: as
 @Date: Creat in 15:52 2022/2/12
 @Description: user.go
*/

// GetCourseCap 获取课程当前的容量
func GetCourseCap(name string) (int, error) {
	// 从数据库拿
	c := model.Course{}
	c.Name = name
	now, err := c.GetNowCourse()
	if err != nil {
		return 0, err
	}
	return now.Cap, nil
}

// BookCourse 成功抢到课
func BookCourse(courseID, studentID string) error {
	cID, _ := strconv.ParseInt(courseID, 10, 64)
	sID, _ := strconv.ParseInt(studentID, 10, 64)
	sc := model.StudentCourse{
		StudentID: sID,
		CourseID:  cID,
	}
	err := model.DB.Create(&sc).Error
	return err
}

func StuHaveCourse(courseID, studentID string) bool {
	cID, _ := strconv.ParseInt(courseID, 10, 64)
	sID, _ := strconv.ParseInt(studentID, 10, 64)
	sc := model.StudentCourse{
		StudentID: sID,
		CourseID:  cID,
	}
	model.DB.Find(&sc, "courseID = ? AND studentID = ?", cID, sID)
	if sc.ID == 0 {
		return false
	}
	return true
}
