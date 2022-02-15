package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/service"
	"project/types"
	"project/util/resp"
)

// CreateMember 添加成员
func CreateMember(c *gin.Context) {

	auth, err := c.Cookie("camp-session")
	if err != nil {
		c.JSON(http.StatusForbidden, resp.CreateMemberRes(types.LoginRequired, ""))
		return
	}
	userType := service.GetUserTypeByCookie(auth)
	if userType != 1 {
		c.JSON(http.StatusForbidden, resp.CreateMemberRes(types.PermDenied, ""))
		return
	}

	request := types.CreateMemberRequest{}

	err = c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, resp.CreateMemberRes(types.ParamInvalid, ""))
		return
	}

	id, errNo := service.CreateMember(request)
	if errNo != 0 {
		c.JSON(http.StatusBadRequest, resp.CreateMemberRes(errNo, id))
		return
	}

	c.JSON(http.StatusOK, resp.CreateMemberRes(errNo, id))
}

// GetMember 获取单个成员
func GetMember(c *gin.Context) {
	request := types.GetMemberRequest{}

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, resp.GetMemberRes(types.ParamInvalid, types.TMember{}))
		return
	}

	tMember, errNo := service.GetMember(request)
	if errNo != 0 {
		c.JSON(http.StatusBadRequest, resp.GetMemberRes(errNo, tMember))
		return
	}

	c.JSON(http.StatusOK, resp.GetMemberRes(errNo, tMember))

}

func GetMemberList(c *gin.Context) {
	request := types.GetMemberListRequest{}

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, resp.GetMemberListRes(types.ParamInvalid, []types.TMember{}))
		return
	}

	tMembers, errNo := service.GetMemberList(request)
	if errNo != 0 {
		c.JSON(http.StatusBadRequest, resp.GetMemberListRes(errNo, tMembers))
		return
	}

	c.JSON(http.StatusOK, resp.GetMemberListRes(errNo, tMembers))
}

func UpdateMember(c *gin.Context) {
	request := types.UpdateMemberRequest{}

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, resp.UpdateMemberRes(types.ParamInvalid))
		return
	}

	errNo := service.UpdateMember(request)

	if errNo != 0 {
		c.JSON(http.StatusBadRequest, resp.UpdateMemberRes(errNo))
		return
	}

	c.JSON(http.StatusOK, resp.UpdateMemberRes(errNo))
}

func DeleteMember(c *gin.Context) {
	request := types.DeleteMemberRequest{}

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, resp.UpdateMemberRes(types.ParamInvalid))
		return
	}

	errNo := service.DeleteMember(request)

	if errNo != 0 {
		c.JSON(http.StatusBadRequest, resp.DeleteMemberRes(errNo))
		return
	}
	errNo = service.DeleteCookies(request.UserID)
	if errNo != 0 {
		c.JSON(http.StatusBadRequest, resp.DeleteMemberRes(errNo))
		return
	}

	c.JSON(http.StatusOK, resp.DeleteMemberRes(errNo))

}
