package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MainController struct {
	BaseController
}

func (this *MainController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "view/main/index.html", map[string]interface{}{})
}

func (this *MainController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "view/main/login.html", map[string]interface{}{})
}

func (this *MainController) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "view/main/register.html", map[string]interface{}{})
}
