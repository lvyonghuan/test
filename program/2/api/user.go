package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"test/program/2/model"
	"test/program/2/service"
	"test/program/2/tool"
	"test/program/2/util"
)

func Register(c *gin.Context) {
	Username := c.PostForm("username")
	PassWord := c.PostForm("Password")
	if Username == "" || PassWord == "" {
		util.RespParamErr(c)
		return
	}
	u, err := service.SearchUserByUserName(Username)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	if u.UserName != "" {
		util.NormErr(c, 300, "用户名已经注册")
		return
	}
	err = service.CreateUser(model.User{
		UserName: Username,
		PassWord: PassWord,
	})
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func Login(c *gin.Context) {
	Username := c.PostForm("username")
	Password := c.PostForm("password")
	u, err := service.SearchUserByUserName(Username)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	if u.UserName == "" {
		util.NormErr(c, 300, "用户不存在")
		return
	}
	if u.PassWord != Password {
		util.NormErr(c, 300, "密码错误")
		return
	}
	token, err := tool.GenerateToken()
	if err != nil {
		log.Printf("search user error:%v", err)
		return
	}
	err = service.InsertCookieToken(Username, token)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	c.SetCookie("LoginState", token, 3600, "/", "localhost", false, true)
	util.RespOK(c)
}
