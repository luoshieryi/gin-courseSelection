package router

import (
	"github.com/gin-gonic/gin"
	"project/api"
)

func Init(r *gin.Engine) {
	g := r.Group("/api/v1")

	// 成员管理
	g.POST("/member/create", api.CreateMember)
	g.GET("/member", api.GetMember)
	g.GET("/member/list", api.GetMemberList)
	g.POST("/member/update", api.UpdateMember)
	g.POST("/member/delete", api.DeleteMember)

	// 登录
	g.POST("/auth/login", api.Login)
	g.POST("/auth/logout", api.Logout)
	g.GET("/auth/whoami", api.Whoami)

	// 排课
	g.POST("/course/create")
	g.GET("/course/get")

	g.POST("/teacher/bind_course")
	g.POST("/teacher/unbind_course")
	g.GET("/teacher/get_course")
	g.POST("/course/schedule")

	// 抢课
	g.POST("/student/book_course")
	g.GET("/student/course")
}
