package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"test/program/2/service"
	"test/program/2/util"
)

func Create(c *gin.Context) {
	cookie, err := c.Cookie("LoginState")
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search message error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	_, err = service.SearchUsernameByCookie(cookie)
	if err != nil {
		if err.Error() == "用户未登录或登陆状态失效" {
			util.NormErr(c, 20013, "用户未登录或登陆状态失效")
			return
		}
		log.Printf("search message error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	StoreHouseName := c.PostForm("storehouse name")
	if StoreHouseName == "" {
		util.RespParamErr(c)
		return
	}
	err = service.Create(StoreHouseName)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func GoodsManage(c *gin.Context) {
	cookie, err := c.Cookie("LoginState")
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search message error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	_, err = service.SearchUsernameByCookie(cookie)
	if err != nil {
		if err.Error() == "用户未登录或登陆状态失效" {
			util.NormErr(c, 20013, "用户未登录或登陆状态失效")
			return
		}
		log.Printf("search message error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	StoreHouseName := c.PostForm("storehouse name")
	GoodsName := c.PostForm("goods name")
	GoodsNum := c.PostForm("goods num")
	if StoreHouseName == "" || GoodsName == "" || GoodsNum == "" {
		util.RespParamErr(c)
		return
	}
}
