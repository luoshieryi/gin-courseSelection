package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/service"
	"project/types"
	"project/util/resp"
)

func GetScheduleCourse(c *gin.Context) {
	request := types.ScheduleCourseRequest{}
	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, resp.ScheduleCourseRes(types.ParamInvalid, map[string]string{}))
		return
	}
	errNo, Data := service.GetScheduleCourse(request)

	if errNo != 0 {
		c.JSON(http.StatusBadRequest, resp.ScheduleCourseRes(errNo, Data))
		return
	}
	c.JSON(http.StatusOK, resp.ScheduleCourseRes(errNo, Data))
}
