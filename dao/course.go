package dao

import (
	"project/model"
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
