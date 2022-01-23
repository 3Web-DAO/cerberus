package models

type Email struct {
	UserId       string `xorm:"not null pk VARCHAR(9)"`
	EmailAddress string `xorm:"not null VARCHAR(64)"`
}
