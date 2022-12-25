package dao

import (
	"errors"
	"fmt"
	"test/program/2/model"
)

func SearchUserByUserName(name string) (u model.User, err error) {
	row := DB.QueryRow("select * from user where username=?", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.UserName, &u.PassWord, &u.Token)
	return
}

func CreateUser(u model.User) (err error) {
	_, err = DB.Exec("insert into user(username,password,token)", u.UserName, u.PassWord, 0)
	return err
}

func InsertCookieToken(username string, token string) (err error) {
	_, err = DB.Exec("update user set token=? where username=?", token, username)
	return err
}

func SearchUsernameByCookie(cookie string) (username string, err error) {
	var u model.User
	row := DB.QueryRow("select * from user where cookie = ?", cookie)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.UserName, &u.PassWord, &u.Token)
	fmt.Println(u.UserName)
	if u.UserName == "" {
		return "", errors.New("用户未登录或登陆状态失效")
	}
	return u.UserName, err
}
