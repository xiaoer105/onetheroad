package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

type MainController struct {
	BaseController
}

var DB *sql.DB

const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "ontheroad"
)

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@/%s", USERNAME, PASSWORD, DATABASE)
	DB, err = sql.Open("mysql", dsn)
	log.Printf("################实例化sql")
	if err != nil {
		log.Fatalf("Open mysql failed,err:%v\n", err)
	}
	return
}

func (this *MainController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "view/main/index.html", map[string]interface{}{})
}

type MainLoginReq struct {
	Username string `json:"username"`
	Password string `json:"Password"`
}

type MainLoginResp struct {
	Result   string `json:"result"`
	Id       int64  `db:"id"`
	IdStr    string `db:"id_str"`
	NickName string `db:"nickname"`
	Password string `db:"password"`
}

func (this *MainController) Login(c *gin.Context) {
	if this.IsGet(c) {
		c.HTML(http.StatusOK, "view/main/login.html", map[string]interface{}{})
		return
	}

	var item MainLoginReq

	item.Username = c.PostForm("Username")
	item.Password = c.PostForm("Password")

	var v MainLoginResp

	err := DB.QueryRow("SELECT id,nickname FROM usr WHERE username = ? AND password = ?", item.Username, item.Password).Scan(&v.Id, &v.NickName)
	v.IdStr = strconv.Itoa(int(v.Id))

	if err != nil {
		log.Println("Login[err]:", err)
		this.AjaxMsg(c, 2001, fmt.Sprintf("%v", err))
		return
	}

	this.JsonResult(c, map[string]interface{}{
		"result": "ok",
	})

}

func CreateUser(usr *MainLoginResp) (tokenss string, err error) {
	// 自定义claim
	//claim := jwt.StandardClaims{
	//	Id: usr.IdStr,
	//}
	return
}

func (this *MainController) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "view/main/register.html", map[string]interface{}{})
}
