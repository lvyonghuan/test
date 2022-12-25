package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"test/program/2/model"
	"test/program/2/service"
	"test/program/2/tool"
	"test/program/2/util"
)

func Create(c *gin.Context) {
	err := tool.Check_login_state(c)
	if err != nil {
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

func GoodsManage_in(c *gin.Context) {
	err := tool.Check_login_state(c)
	if err != nil {
		return
	}
	StoreHouseName := c.PostForm("storehouse name")
	GoodsName := c.PostForm("goods name")
	GoodsNum := c.PostForm("goods num")
	goodsnum, err := strconv.Atoi(GoodsNum)
	if err != nil {
		fmt.Println(err)
		return
	}
	if StoreHouseName == "" || GoodsName == "" || GoodsNum == "" {
		util.RespParamErr(c)
		return
	}
	err = service.StoreManage_in(StoreHouseName, model.Storehouse{
		GoodsName: GoodsName,
		GoodsNum:  goodsnum,
	})
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search message error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func GoodsManage_out(c *gin.Context) {
	err := tool.Check_login_state(c)
	if err != nil {
		return
	}
	StoreHouseName := c.PostForm("storehouse name")
	GoodsName := c.PostForm("goods name")
	err = service.StoreManage_out(StoreHouseName, GoodsName)
	if err != nil && err != sql.ErrNoRows { //这里没有处理找不到货物的情况，，，主要是先把大框弄好再说
		log.Printf("search message error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}
