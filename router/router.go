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
	g.POST("/course/create", api.CreateCourse)
	g.GET("/course/get", api.GetCourse)

	g.POST("/teacher/bind_course", api.BindCourse)
	g.POST("/teacher/unbind_course", api.UnbindCourse)
	g.GET("/teacher/get_course", api.GetTeacherCourse)
	g.POST("/course/schedule", api.GetScheduleCourse)

	// 抢课
	{
		g.POST("/student/book_course", api.BookCourse)
		g.GET("/student/course", api.GetCourse)
	}
}
