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
