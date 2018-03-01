package admin

import (
	"github.com/revel/revel"
)

type Login struct {
	*revel.Controller
}

func (c Login) Login() revel.Result {
	return c.Render()
}

func (c Login) Post(username, password string) revel.Result{

	c.Validation.Required(username).Message("用户名不能为空")
	c.Validation.Required(password).Message("密码不能为空")

	if c.Validation.HasErrors() {
		//c.Flash.Error("用户名或者密码不能为空")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Login.Login)
	}
	return c.Redirect(Admin.Index)
}