package models

type TwitterFollower struct {
	TweetUserId string `xorm:"not null pk VARCHAR(12)"`
	Name        string `xorm:"not null VARCHAR(255)"`
	Id          string `xorm:"not null VARCHAR(12)"`
}
