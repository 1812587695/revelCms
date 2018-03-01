package admin

import (
	"github.com/revel/revel"
	"revelCms/app/models"
)

func init()  {
	revel.OnAppStart(Init)

	// 检查用户登录
	revel.InterceptMethod((*Base).checkUser, revel.BEFORE)

}

func Init()  {
	engineRead = models.GetEngineRead()
	engineWrite = models.GetEngineWrite()
}