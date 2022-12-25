package service

import (
	"test/program/2/dao"
	"test/program/2/model"
)

func SearchUserByUserName(Username string) (u model.User, err error) {
	u, err = dao.SearchUserByUserName(Username)
	return
}

func CreateUser(u model.User) (err error) {
	err = dao.CreateUser(u)
	return err
}

func InsertCookieToken(username string, token string) (err error) {
	err = dao.InsertCookieToken(username, token)
	return err
}

func SearchUsernameByCookie(cookie string) (username string, err error) {
	username, err = dao.SearchUsernameByCookie(cookie)
	return username, err
}
