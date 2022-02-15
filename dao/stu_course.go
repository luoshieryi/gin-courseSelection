package dao

import (
	"project/model"
)

func GetCourseIDsByStudentID(stuID string) ([]string, error) {
	courseIDs := make([]string, 0)
	res := make([]model.StuCourse, 0)
	err := model.DB.Find(&res, "stu_id = ?", stuID).Error
	if err != nil {
		return courseIDs, err
	}
	for _, data := range res {
		courseIDs = append(courseIDs, data.CourseID)
	}
	return courseIDs, err
}
