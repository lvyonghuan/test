package tool

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"test/program/2/service"
	"test/program/2/util"
)

func Check_login_state(c *gin.Context) (err error) {
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
	return
}
