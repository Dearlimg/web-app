package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"web-app/models"
)

const sercet string = "123"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
)

func CheckUserExist(username string) (err error) {
	sqlStr := "select count(user_id) from user where username=?"
	var count int
	err = db.Get(&count, sqlStr, username)
	if err != nil {
		return
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user *models.User) (err error) {
	//
	user.Password, _ = encryptPassword(user.Password)
	//执行sql语句
	strSql := "insert into user(user_id,username,password) value(?,?,?)"
	_, err = db.Exec(strSql, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(password string) (string, error) {
	h := md5.New()
	h.Write([]byte(sercet))
	h.Sum([]byte(password))
	return hex.EncodeToString(h.Sum([]byte(password))), nil
}

func CheckUserPassword(User *models.ParamLogin) error {
	opassword := User.Password
	sqlStr := `select username,password from user where username=?`
	err := db.Get(User, sqlStr, User.Username)
	if err != nil {
		return ErrorUserNotExist
	}
	password, _ := encryptPassword(opassword)
	fmt.Println(opassword, User.Password, password)
	if password != User.Password {
		return ErrorInvalidPassword
	}
	return nil
}
