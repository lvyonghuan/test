package service

import "test/program/2/dao"

func Create(name string) (err error) {
	err = dao.Crate(name)
	return err
}
