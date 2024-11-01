package logic

import (
	"web-app/dao/mysql"
	"web-app/pkg/snowflake"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	//判断用户存不存在
	mysql.QueryUserByUsername()

	//生成uuid
	snowflake.GenID()
	//密码加密
	//保存到数据库
	mysql.InsertUser()

}
