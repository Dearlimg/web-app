package logic

import (
	"web-app/dao/mysql"
	"web-app/models"
	"web-app/pkg/jwt"
	"web-app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	//判断用户存不存在

	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		//数据库查询出错
		return
	}
	//生成uuid
	userID := snowflake.GenID()
	u := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//密码加密
	//保存到数据库
	return mysql.InsertUser(u)
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.CheckUserPassword(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Username)

	//return mysql.CheckUserPassword(p)
}
