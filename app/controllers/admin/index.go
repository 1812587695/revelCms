package admin

import (
	"github.com/revel/revel"
)

type Admin struct {
	Base
}

func (c Admin) Index() revel.Result {
	return c.Render()
}