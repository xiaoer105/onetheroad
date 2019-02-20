package routers

import (
	"app/org/web/OneTheRoad/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AppTLSRouter() http.Handler {
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
	router.Std("/register", (&controllers.MainController{}).Register)

	return router
}
