package models

import (
	"time"
)

type Cms_user struct {
	Id           int64
	Name         string
	Email        string
	Password     string
	Type         int // 1管理员，2普通用户
	Status       int
	Created      time.Time `xorm:"created"`
	Updated      time.Time `xorm:"updated"`
}
