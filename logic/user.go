package logic

import (
	"web-app/dao/mysql"
	"web-app/models"
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
_:
	mysql.InsertUser(u)
	return

}
