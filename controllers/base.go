package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (this *BaseController) IsGet(c *gin.Context) bool {
	return c.Request.Method == "GET"
}

func (this *BaseController) IsPost(c *gin.Context) bool {
	return c.Request.Method == "POST"
}

func (this *BaseController) AjaxMsg(c *gin.Context, code int64, msg string) {
	out := make(map[string]interface{})

	if code == 0 {
		out["result"] = "ok"
	} else {
		out["result"] = "error"
	}
	out["code"] = code
	out["desc"] = msg

	this.JsonResult(c, out)
}

func (this *BaseController) JsonResult(c *gin.Context, out interface{}) {
	c.JSON(http.StatusOK, out)
}
