package service

import (
	"test/program/2/dao"
	"test/program/2/model"
)

func Create(name string) (err error) {
	err = dao.Crate(name)
	return err
}

func StoreManage_in(name string, s model.Storehouse) (err error) {
	err = dao.StoreManage_in(name, s)
	return err
}

func StoreManage_out(storeName string, name string) (err error) {
	err = dao.StoreManage_out(storeName, name)
	return err
}

func StoreManage_stock(storeName string, goodsName string, changeNum int) (err error) {
	err = dao.StoreManage_stock(storeName, goodsName, changeNum)
	return err
}
