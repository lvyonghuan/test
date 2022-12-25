package main

import (
	"test/program/2/api"
	"test/program/2/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
