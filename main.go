package main

// 导入gin包
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project/config"
	"project/router"
)

// 入口函数
func main() {
	// 初始化一个http服务对象
	r := gin.Default()
	// 使用中间件
	r.Use(gin.Recovery())
	// 注册路由
	router.Init(r)
	// 启动程序
	err := r.Run(fmt.Sprintf("%s:%d", config.Server.Address, config.Server.Port))
	if err != nil {
		return
	} // 监听并在 Address:Port (当前设置为localhost:8081) 上启动服务
}
