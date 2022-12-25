package service

import (
	"test/program/2/dao"
	"test/program/2/model"
)

func Create(name string) (err error) {
	err = dao.Crate(name)
	return err
}

func StoreManage_on(name string, s model.Storehouse) (err error) {
	err = dao.StoreManage_on(name, s)
	return err
}
