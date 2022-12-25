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
		storehouse.POST("goodsin", GoodsManage_in)
		storehouse.POST("/goodsout", GoodsManage_out)
		storehouse.POST("/goodsstock", GoodsManage_stock)
	}
	r.Run()
}
