package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/service"
	"project/types"
	"project/util/resp"
)

func CreateCourse(c *gin.Context) {

	request := types.CreateCourseRequest{}

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, resp.CreateCourseRes(types.ParamInvalid, ""))
		return
	}

	id, errNo := service.CreateCourse(request)
	if errNo != 0 {
		c.JSON(http.StatusBadRequest, resp.CreateMemberRes(errNo, id))
		return
	}

	c.JSON(http.StatusOK, resp.CreateMemberRes(errNo, id))
}

func GetCourse(c *gin.Context) {
	request := types.GetCourseRequest{}

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, resp.GetCourseRes(types.ParamInvalid, types.TCourse{}))
		return
	}

	tCourse, errNo := service.GetCourse(request)
	if errNo != 0 {
		c.JSON(http.StatusBadRequest, resp.GetCourseRes(errNo, tCourse))
		return
	}

	c.JSON(http.StatusOK, resp.GetCourseRes(errNo, tCourse))
}

func BindCourse(c *gin.Context) {
	request := types.BindCourseRequest{}
	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, resp.BindCourseRes(types.ParamInvalid))
		return
	}

	errNo := service.BindCourse(request)
	if errNo != 0 {
		c.JSON(http.StatusBadRequest, resp.BindCourseRes(errNo))
		return
	}
	c.JSON(http.StatusOK, resp.BindCourseRes(errNo))
}

func UnbindCourse(c *gin.Context) {
	request := types.UnbindCourseRequest{}
	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, resp.UnbindCourseRes(types.ParamInvalid))
		return
	}

	errNo := service.UnbindCourse(request)
	if errNo != 0 {
		c.JSON(http.StatusBadRequest, resp.UnbindCourseRes(errNo))
		return
	}
	c.JSON(http.StatusOK, resp.UnbindCourseRes(errNo))
}

func GetTeacherCourse(c *gin.Context) {
	request := types.GetTeacherCourseRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, resp.GetTeacherCourseRes(types.ParamInvalid, []*types.TCourse{}))
		return
	}
	TCourses, errNo := service.GetTeacherCourse(request)
	if errNo != 0 {
		c.JSON(http.StatusBadRequest, resp.GetTeacherCourseRes(errNo, TCourses))
		return
	}
}
