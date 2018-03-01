package admin

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/go-xorm/xorm"
	"revelCms/app/models"
)

var (
	engineRead *xorm.Engine
	engineWrite *xorm.Engine
)

type Base struct {
	*revel.Controller
}

// 用户登录检查
func (c *Base) checkUser() revel.Result {
	fmt.Println("--------------------------------2222222222222")

	var user []models.Cms_user

	fmt.Println(user)

	engineRead.Find(&user)

	fmt.Println(user)
	return nil
}