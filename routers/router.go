package routers

import (
	"app/org/web/OnTheRoad/controllers"
	"log"

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

	apiRouter := router.Group("/")
	apiRouter.Use(VerifyLoginToken())
	{
		apiRouter.Std("/", (&controllers.MainController{}).Index)
		apiRouter.Std("/index", (&controllers.MainController{}).Index)
		apiRouter.Std("/login", (&controllers.MainController{}).Login)
		apiRouter.Std("/register", (&controllers.MainController{}).Register)
	}

	return router
}

//
// 验证用户登陆token
func VerifyLoginToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pathUrl = c.Request.URL.Path
		log.Println("###############url", pathUrl)
		tokenId, err := c.Cookie("TokenId")
		log.Println("###############tokenId", tokenId)

		if tokenId, err = c.Cookie("TokenId"); tokenId == "" || err != nil {
			if pathUrl != "/login" { // 不是登陆页需要显示登陆页，进行登陆
				c.HTML(http.StatusOK, "view/main/login.html", map[string]interface{}{})
				c.Abort()
			}
		}

		c.Next()
	}
}
