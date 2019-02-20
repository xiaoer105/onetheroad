package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"app/org/web/OneTheRoad/controllers"
)

func AppTLSRouter() http.Handler{
	return AppRouter()
}

func AppRouter() http.Handler {
	// 创建默认多路复用器
	router := gin.Default()

	// 加载资源文件
	router.LoadHTMLGlob("views/**/**")

	router.Static("/static", "static")
	// 注册函数
	router.Std("/", (&controllers.MainController{}).Index)
	router.Std("/index", (&controllers.MainController{}).Index)
	router.Std("/login", (&controllers.MainController{}).Login)

	return router
}
