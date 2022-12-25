package cmd

import (
	"test/program/2/api"
	"test/program/2/dao"
)

func mian() {
	dao.InitDB()
	api.InitRouter()
}
