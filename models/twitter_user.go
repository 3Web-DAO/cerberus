package models

type TwitterUser struct {
	UserId      string `xorm:"not null pk VARCHAR(12)"`
	Id          string `xorm:"not null VARCHAR(12)"`
	Name        string `xorm:"not null VARCHAR(255)"`
	LastTweetId string `xorm:"not null default '0' VARCHAR(32)"`
}
