package models

import (
	"time"
)

type User struct {
	UserId     string    `xorm:"not null pk VARCHAR(9)"`
	Password   string    `xorm:"not null VARCHAR(20)"`
	UserName   string    `xorm:"not null VARCHAR(64)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP(6)' TIMESTAMP(6)"`
}
