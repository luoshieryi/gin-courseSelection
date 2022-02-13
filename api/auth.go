package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/service"
	"project/types"
	"project/util/resp"
)

//Login 登陆
func Login(c *gin.Context) {
	request := types.LoginRequest{}

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, resp.LoginRes(types.WrongPassword, ""))
		return
	}

	auth, id, errNo := service.Login(request)

	if errNo != 0 {
		c.JSON(http.StatusBadRequest, resp.LoginRes(errNo, id))
		return
	}

	c.SetCookie("camp-session", auth, 30*24*60*60, "/", "localhost", false, true)
	c.JSON(http.StatusOK, resp.LoginRes(errNo, id))
}

func Logout(c *gin.Context) {
	auth, err := c.Cookie("camp-session")
	if err != nil {
		c.JSON(http.StatusForbidden, resp.LogoutRes(types.LoginRequired))
		return
	}

	errNo := service.Logout(auth)

	if errNo != 0 {
		c.JSON(http.StatusBadRequest, resp.LogoutRes(errNo))
		return
	}

	c.SetCookie("camp-session", auth, -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, resp.LogoutRes(errNo))
}

func Whoami(c *gin.Context) {
	auth, err := c.Cookie("camp-session")
	if err != nil {
		c.JSON(http.StatusForbidden, resp.LogoutRes(types.LoginRequired))
		return
	}

	tMember, errNo := service.Whoami(auth)
	if errNo != 0 {
		c.JSON(http.StatusBadRequest, resp.WhoamiRes(errNo, tMember))
		return
	}

	c.JSON(http.StatusOK, resp.WhoamiRes(errNo, tMember))
}
