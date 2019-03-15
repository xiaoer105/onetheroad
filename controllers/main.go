package controllers

import (
	"app/org/web/OnTheRoad/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

type MainController struct {
	BaseController
}

func (this *MainController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "view/main/index.html", map[string]interface{}{})
}

type MainLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"Password"`
}

type MainLoginResp struct {
	Result   string `json:"result"`
	Id       int64  `db:"id"`
	IdStr    string `db:"id_str"`
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

func (this *MainController) Login(c *gin.Context) {
	if this.IsGet(c) {
		c.HTML(http.StatusOK, "view/main/login.html", map[string]interface{}{})
		return
	}

	var item MainLoginReq

	item.Email = c.PostForm("Email")
	item.Password = c.PostForm("Password")

	var v MainLoginResp

	err := models.DB.QueryRow("SELECT id,nickname,email FROM usr WHERE email = ? AND password = ?", item.Email, item.Password).Scan(&v.Id, &v.Name, &v.Email)
	v.IdStr = strconv.Itoa(int(v.Id))

	if err != nil {
		log.Println("Login[err]:", err)
		this.AjaxMsg(c, 2001, fmt.Sprintf("%v", err))
		return
	}

	j := &models.JWT{
		[]byte("xiaoer"),
	}

	claims := &models.CustomClaims{
		v.Id,
		v.Name,
		v.Email,
		jwt.StandardClaims{
			ExpiresAt: 15000, //time.Now().Add(24 * time.Hour).Unix()
			Issuer:    "xiaoer",
		},
	}
	token, err := j.CreateToken(*claims)

	c.SetCookie("TokenId", token, 7200, "/", "localhost", false, true)

	this.JsonResult(c, map[string]interface{}{
		"result": "ok",
	})

}

func (this *MainController) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "view/main/register.html", map[string]interface{}{})
}
