package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	use := r.Group("/user")
	{
		use.POST("register", Register)
		use.POST("login", Login)
	}
	storehouse := r.Group("/storehouse")
	{
		storehouse.POST("create", Create)
		storehouse.POST("goodson", GoodsManage_in)
	}
	r.Run()
}
